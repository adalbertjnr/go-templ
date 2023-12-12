package main

type UserData struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	Address   struct {
		ID             int     `json:"id"`
		Street         string  `json:"street"`
		StreetName     string  `json:"streetName"`
		BuildingNumber string  `json:"buildingNumber"`
		City           string  `json:"city"`
		Zipcode        string  `json:"zipcode"`
		Country        string  `json:"country"`
		CountyCode     string  `json:"county_code"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
	} `json:"address"`
	Website string `json:"website"`
	Image   string `json:"image"`
}

type Users struct {
	Status string     `json:"status"`
	Code   int        `json:"code"`
	Total  int        `json:"total"`
	Data   []UserData `json:"data"`
}
