package entity

type Car struct {
	ID          int    `json:"id"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Type        string `json:"type"`
	Engine      string `json:"engine"`
	Chassis     string `json:"chassis"`
	GrossWeight int    `json:"grossWeight"`
	Seats       int    `json:"seats"`
	Color       string `json:"color"`
	TotalMass   int    `json:"totalMass"`
	TowerMass   int    `json:"towerMass"`
	NumberPlate string `json:"numberPlate"`
	// car driver
	DriverID int `json:"driverID"`
}
