package manage_car

// define errors for manage car usecase

import "errors"

var (
	ErrExistedCar         = errors.New("car already existed")
	ErrNumberPlateExisted = errors.New("number plate has existed")
	ErrModelExisted       = errors.New("model has existed")
	ErrEngineExisted      = errors.New("engine has existed")
	ErrChassisExisted     = errors.New("chassis has existed")

	// details car error
	ErrNotFound = errors.New("car not found")
)
