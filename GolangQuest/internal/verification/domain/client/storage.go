package user

import (
	"github.com/eNViDAT0001/Backend/internal/verification/domain/jwt/storage/io"
)

type Storage interface {
	FindClientByKey(clientKey string) (io.Client, error)
	FindClientById(clientId string) (io.Client, error)
}
