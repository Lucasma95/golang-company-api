package entities

import (
	"time"

	contract "github.com/Lucasma95/golang-company-api/src/api/http/contracts/company"
	"github.com/lithammer/shortuuid/v4"
	"gorm.io/gorm"
)

type Company struct {
	ID          string `gorm:"type:varchar;primaryKey"`
	Name        string `gorm:"type:varchar;not null"`
	Description string
	CountryName string 
	Country     Country `gorm:"foreignKey:CountryName"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (q *Company) BeforeCreate(tx *gorm.DB) (err error) {
	q.ID = shortuuid.New()
	return
}

func NewCompany(request *contract.CreateCompanyRequest) Company {

	return Company{
		Name:        request.Name,
		Description: request.Description,
		CountryName: request.CountryName,
	}
}

type CompanyDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryName string `json:"country_name"`
	Continent   string `json:"continent"`
}

func NewCompanyDTO(company *Company) CompanyDTO {
	return CompanyDTO{
		ID:          company.ID,
		Name:        company.Name,
		Description: company.Description,
		CountryName: company.CountryName,
		Continent:   company.Country.Continent,
	}
}
