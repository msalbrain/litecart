package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Page is ...
type Page struct {
	Core
	Name     string  `json:"name"`
	Slug     string  `json:"slug"`
	Position string  `json:"position,omitempty"`
	Content  *string `json:"content,omitempty"`
	Active   bool    `json:"active"`
	Seo      *Seo    `json:"seo,omitempty"`
}

// Validate is ...
func (v Page) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.ID, validation.Length(15, 15)),
		validation.Field(&v.Name, validation.Length(3, 50)),
		validation.Field(&v.Slug, validation.Length(3, 20)),
	)
}
