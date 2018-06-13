package models

import "github.com/go-ozzo/ozzo-validation"

// Country represents an artist record.
type Country struct {
	Id   		int    `json:"id" db:"id"`
	Iso 		string `json:"iso" db:"iso"`
	Iso3 		string `json:"iso3" db:"iso3"`
	IsoNumber	string `json:"iso_number" db:"iso_number"`
	Name 		string `json:"name" db:"name"`
	Capital 	string `json:"capital" db:"capital"`
	Continent 	string `json:"continent" db:"continent"`
	Currency 	string `json:"currency" db:"currency"`
	CurrencyName string `json:"currency_name" db:"currency_name"`
	Phone 		string `json:"phone" db:"phone"`
}

// Validate validates the Country fields.
func (m Country) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	)
}
