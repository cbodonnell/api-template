package services

import (
	"github.com/cbodonnell/api-template/models"
)

type Service interface {
	GetUserByName(name string) (models.User, error)
}
