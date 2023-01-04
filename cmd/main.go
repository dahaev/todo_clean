package main

import (
	"os"

	"github.com/dahaev/todo.git"
	"github.com/dahaev/todo.git/pkg/handler"
	"github.com/dahaev/todo.git/pkg/repository"
	"github.com/dahaev/todo.git/pkg/services"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("config down - %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("env error - %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   viper.GetString("db.name"),
		SSLmode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Error connect to DB - %s", err.Error())
	}
	logrus.Println("Starting server...")
	repos := repository.NewRepository(db)
	services := services.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(todo.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Server cannot run %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
