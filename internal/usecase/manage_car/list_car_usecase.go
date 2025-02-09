package manage_car

import (
	"context"
	"r01/internal/entity"
	"r01/internal/repository"
)

type ListCarUsecase struct {
	carRepo repository.ICarRepository
}

func NewListCarUsecase(carRepo repository.ICarRepository) *ListCarUsecase {
	return &ListCarUsecase{carRepo: carRepo}
}

type ListCarParams struct {
	PageSize int `json:"pageSize"`
	Page     int `json:"page"`
	CursorID int `json:"cursorID"`
}

func (lc ListCarUsecase) Execute(ctx context.Context, listParams ListCarParams) ([]entity.Car, error) {
	// Set default values for pageSize and page if not provided
	pageSize := listParams.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	page := listParams.Page
	if page == 0 {
		page = 1
	}

	// Get all cars
	cars, err := lc.carRepo.List(ctx, pageSize, page, listParams.CursorID)
	if err != nil {
		return nil, err
	}

	return cars, nil
}
