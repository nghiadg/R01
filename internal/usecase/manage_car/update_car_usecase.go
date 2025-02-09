package manage_car

import (
	"context"
	"log"
	"r01/internal/entity"
	"r01/internal/repository"

	"github.com/go-playground/validator/v10"
)

type UpdateCarUsecase struct {
	carRepo repository.ICarRepository
}

func NewUpdateCarUsecase(carRepo repository.ICarRepository) *UpdateCarUsecase {
	return &UpdateCarUsecase{
		carRepo: carRepo,
	}
}

type UpdateCarParams struct {
	ID          int    `json:"id" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Model       string `json:"model" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Engine      string `json:"engine" validate:"required"`
	Chassis     string `json:"chassis"`
	GrossWeight int    `json:"grossWeight"`
	Seats       int    `json:"seats" validate:"required"`
	Color       string `json:"color" validate:"required"`
	TotalMass   int    `json:"totalMass"`
	TowerMass   int    `json:"towerMass"`
	NumberPlate string `json:"numberPlate" validate:"required"`
	// car driver
	DriverID int `json:"driverID" validate:"required"`
}

func (cc UpdateCarParams) Validate() error {
	validate := validator.New()
	return validate.Struct(cc)
}

func (uc *UpdateCarUsecase) Execute(ctx context.Context, carParams UpdateCarParams) error {
	if err := carParams.Validate(); err != nil {
		log.Println(err)
		return err
	}

	car := entity.Car{
		ID:          carParams.ID,
		Brand:       carParams.Brand,
		Model:       carParams.Model,
		Type:        carParams.Type,
		Engine:      carParams.Engine,
		Chassis:     carParams.Chassis,
		GrossWeight: carParams.GrossWeight,
		Seats:       carParams.Seats,
		Color:       carParams.Color,
		TotalMass:   carParams.TotalMass,
		TowerMass:   carParams.TowerMass,
		NumberPlate: carParams.NumberPlate,
		DriverID:    carParams.DriverID,
	}

	return uc.carRepo.Update(ctx, car)

}
