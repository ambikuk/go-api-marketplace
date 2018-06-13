package models

import "github.com/go-ozzo/ozzo-validation"

// Province represents an artist record.
type Province struct {
	Id   		int    `json:"id" db:"id"`
	Name 		string `json:"name" db:"name"`
	CountryId 	int		`json:"country_id" db:"country_id"`
}

// Validate validates the Province fields.
func (m Province) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	)
}
