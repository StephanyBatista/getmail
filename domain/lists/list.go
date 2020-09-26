package lists

import (
	"getmail/domain"
)

// List is where one subscriber can stay
type List struct {
	Base domain.Base `json:"base" gorm:"embedded"`
	Name string      `json:"name" validate:"required"`
}

//New creates a new subscriber
func New(name string) (List, error) {
	base := domain.NewBase()
	model := List{
		Base: base,
		Name: name,
	}

	err := domain.ValidateData(model)
	return model, err
}
