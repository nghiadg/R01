package repository

import (
	"context"
	"r01/internal/entity"
)

type ICarRepository interface {
	Create(ctx context.Context, car entity.Car) error
	HasExisted(ctx context.Context, numberPlate string) (bool, error)
	List(ctx context.Context, pageSize, page, lastID int) ([]entity.Car, error)
}
