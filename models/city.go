package models

import "github.com/go-ozzo/ozzo-validation"

// City represents an artist record.
type City struct {
	Id   		int    `json:"id" db:"id"`
	Name 		string `json:"name" db:"name"`
	ProvinceId 	int		`json:"province_id" db:"province_id"`
}

// Validate validates the City fields.
func (m City) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	)
}
