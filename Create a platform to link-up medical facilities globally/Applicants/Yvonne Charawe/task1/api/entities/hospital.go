package entities

//Hospital is a model containing the properties for hospitals 
type Hospital struct {
	ID uint64 `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Country string `json:"country"`
	Region string `json:"region"`
	OperatingHours string `json:"operating hours"`
	PhoneNumber string `json:"phone number"`
	Services string `json:"services"`
	Website string `json:"website"`
}

//Title to show list of hospitals

type HospitalList struct {
	Title string `json:"title"`
	Hospitals []Hospital `json:"hospitals"`
}