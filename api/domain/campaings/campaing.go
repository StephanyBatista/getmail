package campaings

import (
	"fmt"
	"getmail/domain"
	"time"
)

const (
	RegularType   = "Regular"
	AutomatedType = "Automated"
	DraftStatus   = "Draft"
	ReadyStatus   = "Ready"
	StartedStatus = "Started"
	SentStatus    = "Sent"
)

//Campaing represents a campaing
type Campaing struct {
	Base      domain.Base `json:"base" gorm:"embedded"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	ListID    string      `json:"listid"`
	Status    string      `json:"status"`
	Email     Email       `json:"email" gorm:"embedded"`
	StartedAt time.Time
}

//New creates a new campaing
func New(name string, campaingType string) (*Campaing, error) {

	if len(name) == 0 {
		return nil, fmt.Errorf("The Name field is required")
	}

	switch campaingType {
	case AutomatedType, RegularType:
		base := domain.NewBase()
		model := &Campaing{
			Base:   base,
			Name:   name,
			Type:   campaingType,
			Status: DraftStatus,
		}

		return model, nil
	}

	return nil, fmt.Errorf("Invalid campaing type")
}

//SendToList tells which list should be sent
func (c *Campaing) SendToList(listID string) error {
	if len(listID) == 0 {
		return fmt.Errorf("Subscriber List is invalid")
	}

	c.ListID = listID
	return nil
}

//ConfigureEmail configures the email to be sent
func (c *Campaing) ConfigureEmail(fromName, fromEmail, subject, body string) error {

	if err := validateEmail(fromName, fromEmail, subject, body); err != nil {
		return err
	}

	c.Email = Email{
		FromName:  fromName,
		FromEmail: fromEmail,
		Subject:   subject,
		Body:      body,
	}

	return nil
}

func validateEmail(fromName, fromEmail, subject, body string) error {

	if len(fromName) == 0 {
		return fmt.Errorf("Name is required")
	} else if len(fromEmail) == 0 {
		return fmt.Errorf("Email is required")
	} else if len(subject) == 0 {
		return fmt.Errorf("Subject is required")
	} else if len(body) == 0 {
		return fmt.Errorf("Body is required")
	}

	return nil
}
