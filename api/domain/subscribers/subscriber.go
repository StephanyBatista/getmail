package subscribers

import (
	"fmt"
	"getmail/domain"
)

// Subscriber is one person that subscriber
type Subscriber struct {
	Base   domain.Base `json:"base" gorm:"embedded"`
	Email  string      `json:"email" validate:"required,email"`
	Name   string      `json:"name" validate:"required"`
	ListID string      `json:"listid"`
}

//New creates a new subscriber
func New(email, name string) (Subscriber, error) {
	base := domain.NewBase()
	model := Subscriber{
		Base:  base,
		Email: email,
		Name:  name,
	}

	err := domain.ValidateData(model)
	return model, err
}

//PutOnList puts the subscriber in the list choosed
func (s *Subscriber) PutOnList(listID string) error {

	if len(listID) == 0 {
		return fmt.Errorf("The list must be valid to put subscriber in a list")
	}

	s.ListID = listID

	return nil
}
