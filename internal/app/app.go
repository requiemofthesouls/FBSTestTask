package app

import (
	r "FBSTestTask/internal/repository"
	"FBSTestTask/internal/server"
	"FBSTestTask/internal/service"
	"FBSTestTask/internal/transport/grpc"
	"FBSTestTask/internal/transport/http"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"main": "redis:6379",
		},
	})

	rc := cache.New(&cache.Options{
		Redis: ring,
	})

	repo := r.NewFibonacciRepository(rc)
	svc := service.NewFibonacciService(*repo)

	// Build HTTP Server
	handlers := http.NewHandler(*svc)
	srv := server.NewServer(handlers.Init())
	go func() {
		if err := srv.RunHTTP(); err != nil {
			log.Println(err)
		}
	}()

	log.Println("server running on port 8080")

	//Build GRPC Server
	grpcServer := grpc.NewServer(*svc)
	go func() {
		if err := srv.RunGRPC(*grpcServer); err != nil {
			log.Println(err)
		}
	}()
	log.Println("running grpc server on 5300")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

}
