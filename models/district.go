package models

import "github.com/go-ozzo/ozzo-validation"

// District represents an artist record.
type District struct {
	Id   		int    `json:"id" db:"id"`
	Name 		string `json:"name" db:"name"`
	RegencyId 	int		`json:"regency_id" db:"regency_id"`
}

// Validate validates the District fields.
func (m District) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	)
}
