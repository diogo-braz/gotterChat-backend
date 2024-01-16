package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joaovds/chat/configs"
	"github.com/joaovds/chat/infra/webserver/routes"
)

func main() {
	configs.LoadEnv()

	router := chi.NewRouter()

	routes.SetupRoutes(router)

  go func() {
    for {
      printMemoryStats()
      time.Sleep(time.Second)
    }
  }()

	fmt.Println("Server running on port", configs.ENV.Port)
	log.Fatal(http.ListenAndServe(":"+configs.ENV.Port, router))
}

func printMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc: %v MB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc: %v MB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys: %v MB\n", bToMb(m.Sys))
	fmt.Printf("NumGC: %v\n", m.NumGC)
  fmt.Printf("Goroutines: %v\n", runtime.NumGoroutine())
  fmt.Printf("NumCPU: %v\n", runtime.NumCPU())
	fmt.Println("------")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
