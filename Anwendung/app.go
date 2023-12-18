package Anwendung

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	router http.Handler
	DB     *mongo.Database
	
}

func New() (*App, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	fmt.Println("Connected to MongoDB")

	databaseName := "Docker"
	DB := client.Database(databaseName)

	productCollection := DB.Collection("lager")

	productHandler := &Product{Collection: productCollection}

	app := &App{
		router: loadRoutes(productHandler),
		DB:     DB,
	}
	return app, nil
}




func (p *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: p.router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to listen to server: %w", err)
	}
	return nil
}