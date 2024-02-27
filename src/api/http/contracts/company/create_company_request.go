package company

type CreateCompanyRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	CountryName string `json:"country_name"`
}
