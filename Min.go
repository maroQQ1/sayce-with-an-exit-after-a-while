package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "start working!")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "")
	})

	server := &http.Server{
		Addr: ":8080",
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	doWork(ctx)
	
	<-ctx.Done()
	log.Println("Shutting down server...")
	server.Shutdown(context.Background())
}

func doWork(ctx context.Context) {
	newCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	for {
		select {
		case <-newCtx.Done():
			log.Printf("Deadline: %v", newCtx.Err())
			return
		default:
			log.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}
}
