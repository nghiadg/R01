package manage_car

import (
	"context"
	"r01/internal/repository"
)

type DeleteCarUsecase struct {
	carRepo repository.ICarRepository
}

func NewDeleteCarUsecase(carRepo repository.ICarRepository) *DeleteCarUsecase {
	return &DeleteCarUsecase{carRepo: carRepo}
}

func (dc DeleteCarUsecase) Execute(ctx context.Context, id int) error {
	return dc.carRepo.Delete(ctx, id)
}
