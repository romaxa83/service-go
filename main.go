package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/romaxa83/service-go/internal/rabbit"
	"github.com/romaxa83/service-go/internal/server"
	"github.com/romaxa83/service-go/internal/service"
	"github.com/romaxa83/service-go/pkg/mongo"
	"log"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}
}

func main() {
	log.Println("start")
	_, cancel := start()
	defer shutdown(cancel)
	service.WaitShutdown()
}

func start() (ctx context.Context, cancel context.CancelFunc) {
	// This is the main context for the service. When it is canceled it means the service is going down.
	// All the tasks must be canceled
	ctx, cancel = context.WithCancel(context.Background())
	if err := mongo.Start(); err != nil {
		log.Printf("couldnt start database error [%s]\n", err)
	}

	rabbit.Start(ctx)
	server.Start()

	return
}

func shutdown(cancel context.CancelFunc) {
	cancel()
	ctx, cancelTimeout := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelTimeout()

	doneHTTP := server.Shutdown(ctx)
	doneRabbit := rabbit.Shutdown(ctx)
	err := service.WaitUntilIsDoneOrCanceled(ctx, doneHTTP, doneRabbit)
	if err != nil {
		log.Printf("service stopped by timeout %s\n", err)
	}
	if err := mongo.Close(context.Background()); err != nil {
		log.Printf("couldnt close mongo connection error [%s]\n", err)
	}
	time.Sleep(time.Millisecond * 200)
	log.Println("bye bye")
}
