package campaings

//Email represents an email to be sent
type Email struct {
	FromName  string `json:"fromname"`
	FromEmail string `json:"fromemail"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}
