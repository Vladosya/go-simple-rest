package main

import (
	todo "github.com/Vladosya/go-test-rest"
	"github.com/Vladosya/go-test-rest/pkg/handler"
	"github.com/Vladosya/go-test-rest/pkg/repository"
	"github.com/Vladosya/go-test-rest/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	initConfig() // проверяем есть ли в .env нужные параметры

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DBName:   "go_test",
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)                                     // сначала создаётся репозиторий (работа с базой данных)
	services := service.NewService(repos)                                     // потом сервис который зависит от репозитория (бизнес логика)
	handlers := handler.NewHandler(services)                                  // берем обработчики, где роуты, который зависит от сервиса (роуты и обработчики)
	srv := new(todo.Server)                                                   // создаем сервер
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil { // делаем проверку если сервер не запустился то вызываем ошибку
		log.Fatalf("error accured while running http server in main.go: %s", err.Error())
	}
}

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	for _, k := range []string{"PORT"} {
		if _, ok := os.LookupEnv(k); !ok {
			log.Fatal("set environment variable -> ", k)
		}
	}
}
