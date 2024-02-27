package country

type CreateCountryRequest struct {
	Name      string `json:"name" binding:"required"`
	Continent string `json:"continent"`
}
