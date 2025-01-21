package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserFindByID struct {
	ID string
}

func (a UserFindByID) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ID, validation.Required.Error("IDは必須入力です")),
	)
}
