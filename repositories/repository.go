package repositories

import (
	"github.com/cbodonnell/api-template/models"
)

type Repository interface {
	Close()
	QueryUserByName(name string) (models.User, error)
}
