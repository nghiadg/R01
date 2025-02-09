package handler

import (
	"encoding/json"
	"net/http"
	"r01/internal/delivery/http/http_utils"
	"r01/internal/usecase/manage_car"
)

type ManageCarHandler struct {
	CreateCarUsecase manage_car.CreateCarUsecase
}

func NewManageCarHandler(createCarUsecase manage_car.CreateCarUsecase) *ManageCarHandler {
	return &ManageCarHandler{CreateCarUsecase: createCarUsecase}
}

func (mch *ManageCarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	var params manage_car.CreateCarParams
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	err = mch.CreateCarUsecase.Execute(r.Context(), params)

	if err != nil {
		switch err {
		case manage_car.ErrExistedCar:
			http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
				Status:  http.StatusBadRequest,
				Message: "Car already existed",
			})
		case manage_car.ErrModelExisted:
			http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
				Status:  http.StatusBadRequest,
				Message: "Model already existed",
				Error: []http_utils.Error{
					{
						Code:    "model_existed",
						Message: "Model already existed",
						Field:   "model",
					},
				},
			})
		case manage_car.ErrNumberPlateExisted:
			http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
				Status:  http.StatusBadRequest,
				Message: "Number plate already existed",
				Error: []http_utils.Error{
					{
						Code:    "number_plate_existed",
						Message: "Number plate already existed",
						Field:   "numberPlate",
					},
				},
			})
		case manage_car.ErrEngineExisted:
			http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
				Status:  http.StatusBadRequest,
				Message: "Engine already existed",
				Error: []http_utils.Error{
					{
						Code:    "engine_existed",
						Message: "Engine already existed",
						Field:   "engine",
					},
				},
			})
		case manage_car.ErrChassisExisted:
			http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
				Status:  http.StatusBadRequest,
				Message: "Chassis already existed",
				Error: []http_utils.Error{
					{
						Code:    "chassis_existed",
						Message: "Chassis already existed",
						Field:   "chassis",
					},
				},
			})
		default:
			http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal server error",
			})
		}

		return

	}

	http_utils.WriteJSONResponse(w, http_utils.HttpResponse{
		Status:  http.StatusOK,
		Message: "Car created",
	})
}
