package manage_car

import (
	"context"
	"r01/internal/entity"
	"r01/internal/repository"
)

type DetailsCarUsecase struct {
	carRepo repository.ICarRepository
}

func NewDetailsCarUsecase(carRepo repository.ICarRepository) *DetailsCarUsecase {
	return &DetailsCarUsecase{carRepo: carRepo}
}

func (dc DetailsCarUsecase) Execute(ctx context.Context, id int) (entity.Car, error) {
	car, err := dc.carRepo.GetByID(ctx, id)
	if err != nil {
		return entity.Car{}, err
	}

	return car, nil
}
