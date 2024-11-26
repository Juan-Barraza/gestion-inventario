package main

import (
	"context"
	"inventory/internal/config"
	"inventory/routes"
	"inventory/pkg"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	db, err := config.SetConfig()
	if err != nil {
		log.Fatalf("error al configurar la BD: %v", err)
		return
	}

	app, err := pkg.InitFiber()
	if err != nil {
		log.Fatal("error al configurar fiber")
	}
	routes.SetRoutes(app,db)

	contx, sleep := signal.NotifyContext(context.Background(), os.Interrupt)
	defer sleep()

	go func ()  {
		if err := app.Listen(":8000"); err != nil {
			log.Fatalf("no se pudo iniciar el servidor: %v", err)
		}
	}()

	<- contx.Done()
	log.Println("interruption signal")

	time.Sleep(3 * time.Second)
}