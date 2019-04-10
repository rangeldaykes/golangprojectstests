package interfaces

import "service-pattern-go-master/models"

type IPlayerRepository interface {
	GetPlayerByName(name string) (models.PlayerModel, error)
}
