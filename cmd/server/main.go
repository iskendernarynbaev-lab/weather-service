package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-co-op/gocron/v2"
)

const serverPort = ":3000"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			log.Println(err)
		}
	})

	_, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := http.ListenAndServe(serverPort, r)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("server listening on port %s", serverPort)

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		scheduler, err := gocron.NewScheduler()

		if err != nil {
			log.Fatal(err)
		}

		_, err = runJobs(scheduler)

		if err != nil {
			log.Fatal(err)
		}

		scheduler.Start()
	}()

	// ловим ctrl+c
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	cancel()
	wg.Wait()
}
