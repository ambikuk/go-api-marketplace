package models

import "github.com/go-ozzo/ozzo-validation"

// Village represents an artist record.
type Village struct {
	Id   		int    `json:"id" db:"id"`
	Name 		string `json:"name" db:"name"`
	DistrictId 	int		`json:"district_id" db:"district_id"`
}

// Validate validates the Village fields.
func (m Village) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	)
}
