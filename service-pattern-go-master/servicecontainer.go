package main

import (
	"service-pattern-go-master/controllers"
	"service-pattern-go-master/infrastructures"
	"service-pattern-go-master/repositories"
	"service-pattern-go-master/services"
	"sync"

	"database/sql"
)

type IServiceContainer interface {
	InjectPlayerController() controllers.PlayerController
}

type kernel struct{}

func (k *kernel) InjectPlayerController() controllers.PlayerController {

	sqlConn, _ := sql.Open("sqlite3", "/var/tmp/tennis.db")
	sqliteHandler := &infrastructures.SQLiteHandler{}
	sqliteHandler.Conn = sqlConn

	playerRepository := &repositories.PlayerRepository{sqliteHandler}
	playerService := &services.PlayerService{&repositories.PlayerRepositoryWithCircuitBreaker{playerRepository}}
	playerController := controllers.PlayerController{playerService}

	return playerController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
