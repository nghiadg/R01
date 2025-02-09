package repository

import (
	"context"
	"r01/internal/entity"
)

type ICarRepository interface {
	Create(ctx context.Context, car entity.Car) error
	HasExisted(ctx context.Context, numberPlate string) (bool, error)
}
