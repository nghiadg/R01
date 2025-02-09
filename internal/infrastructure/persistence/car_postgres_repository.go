package persistence

import (
	"context"
	"r01/internal/entity"
	"r01/internal/usecase/manage_car"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// const
const (
	DuplicateKeyError     = "23505"
	ConstraintNumberPlate = "cars_no_plate_unique"
	ConstraintModel       = "cars_model_unique"
	ConstraintEngine      = "cars_engine_unique"
	ConstraintChassis     = "cars_chassis_unique"
)

type CarPostgresRepository struct {
	db *pgxpool.Pool
}

func NewCarPostgresRepository(db *pgxpool.Pool) *CarPostgresRepository {
	return &CarPostgresRepository{db: db}
}

func (cpr CarPostgresRepository) Create(ctx context.Context, car entity.Car) error {
	// query to database
	query := `INSERT INTO cars (brand, model, type, engine, chassis, gross_weight, seats, color, total_mass, tower_mass, number_plate, driver_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := cpr.db.Exec(ctx, query, car.Brand, car.Model, car.Type, car.Engine,
		car.Chassis, car.GrossWeight, car.Seats, car.Color, car.TotalMass, car.TowerMass, car.NumberPlate, car.DriverID)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == DuplicateKeyError {
			switch pgErr.ConstraintName {
			case ConstraintModel:
				return manage_car.ErrModelExisted
			case ConstraintNumberPlate:
				return manage_car.ErrNumberPlateExisted
			case ConstraintEngine:
				return manage_car.ErrEngineExisted
			case ConstraintChassis:
				return manage_car.ErrChassisExisted
			}
		}
		return err
	}
	return nil
}

func (cpr CarPostgresRepository) HasExisted(ctx context.Context, numberPlate string) (bool, error) {
	// query to database
	err := cpr.db.QueryRow(ctx, `SELECT id FROM cars WHERE number_plate = $1`, numberPlate).Scan(new(int))
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (cpr CarPostgresRepository) List(ctx context.Context, pageSize, page, cursorID int) ([]entity.Car, error) {
	// query to database
	var rows pgx.Rows
	var err error

	if page == 1 {
		rows, err = cpr.db.Query(ctx, `SELECT id, brand, model, type, engine, chassis, gross_weight, seats, color, total_mass, tower_mass, number_plate, driver_id 
		FROM cars ORDER BY id ASC LIMIT $1`, pageSize)
	} else {
		rows, err = cpr.db.Query(ctx, `SELECT id, brand, model, type, engine, chassis, gross_weight, seats, color, total_mass, tower_mass, number_plate, driver_id 
		FROM cars WHERE id > $1 ORDER BY id ASC LIMIT $2`, cursorID, pageSize)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []entity.Car
	for rows.Next() {
		var car entity.Car
		err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Type, &car.Engine, &car.Chassis,
			&car.GrossWeight, &car.Seats, &car.Color, &car.TotalMass, &car.TowerMass, &car.NumberPlate, &car.DriverID)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (cpr CarPostgresRepository) GetByID(ctx context.Context, id int) (entity.Car, error) {
	// query to database
	var car entity.Car
	err := cpr.db.QueryRow(ctx, `SELECT id, brand, model, type, engine, chassis, gross_weight, seats, color, total_mass, tower_mass, number_plate, driver_id 
	FROM cars WHERE id = $1`, id).Scan(&car.ID, &car.Brand, &car.Model, &car.Type, &car.Engine, &car.Chassis,
		&car.GrossWeight, &car.Seats, &car.Color, &car.TotalMass, &car.TowerMass, &car.NumberPlate, &car.DriverID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.Car{}, manage_car.ErrNotFound
		}
		return entity.Car{}, err
	}
	return car, nil
}
