package entity

type CarDetails struct {
	ID             int    `json:"ID"`
	Brand          string `json:"car"`
	Model          string `json:"model"`
	Year           int    `json:"model_year"`
	OwnerFirstName string `json:"owner_firstname"`
	OwnerLastName  string `json:"owner_lastname"`
	OwnerEmail     string `json:"owner_email"`
	OwnerJobTitle  string `json:"owner_job_title"`
}
