package subscribers

import (
	"getmail/domain"
	"getmail/util"

	guuid "github.com/google/uuid"
)

// Subscriber is one person that subscriber
type Subscriber struct {
	Base   domain.Base `json:"base" gorm:"embedded"`
	Email  string      `json:"email" validate:"required,email"`
	Name   string      `json:"name" validate:"required"`
	ListID guuid.UUID  `json:"listid"`
}

//New creates a new subscriber
func New(email, name string) (*Subscriber, error) {
	base := domain.NewBase()
	model := &Subscriber{
		Base:  base,
		Email: email,
		Name:  name,
	}

	err := util.ValidateData(model)
	return model, err
}

//PutOnList put the subscriber in the list choosed
func (s *Subscriber) PutOnList(ListID guuid.UUID) {
	s.ListID = ListID
}
