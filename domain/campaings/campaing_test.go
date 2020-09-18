package campaings_test

import (
	"getmail/domain/campaings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name         = "Teste"
	campaingType = "Regular"
	listId       = "XA2334"
)

func BuildNewCampaing(name, campaingType string) (*campaings.Campaing, error) {

	return campaings.New(name, campaingType)
}

func Test_MustCreateANewCampaing(t *testing.T) {

	campaing, _ := BuildNewCampaing(name, campaingType)

	assert.Equal(t, name, campaing.Name)
	assert.Equal(t, campaingType, campaing.Type)
}

func Test_MustValidateNameWhenCreateCampaing(t *testing.T) {

	const errorExpected = "The Name field is required"
	const invalidName = ""

	_, err := campaings.New(invalidName, campaingType)

	assert.Equal(t, err.Error(), errorExpected)
}

func Test_MustValidateCampaingTypeWhenCreateCampaing(t *testing.T) {

	const errorExpected = "Invalid campaing type"
	const invalidCampaingType = "Teste"

	_, err := campaings.New(name, invalidCampaingType)

	assert.Equal(t, err.Error(), errorExpected)
}

func Test_MustNewCampaingBeCreatedAsDraft(t *testing.T) {

	campaing, _ := campaings.New(name, campaingType)

	assert.Equal(t, campaings.DraftStatus, campaing.Status)
}

func Test_MustCampaingHasASubscriberList(t *testing.T) {

	campaing, _ := campaings.New(name, campaingType)

	campaing.SendToList(listId)

	assert.Equal(t, listId, campaing.ListID)
}

func Test_MustValidateSubscriberList(t *testing.T) {

	const errorExpected = "Subscriber List is invalid"
	const invalidListId = ""
	campaing, _ := campaings.New(name, campaingType)

	err := campaing.SendToList(invalidListId)

	assert.Equal(t, errorExpected, err.Error())
}

func Test_MustConfigureEmail(t *testing.T) {

	campaing, _ := campaings.New(name, campaingType)

	campaing.ConfigureEmail("teste", "teste@teste.com.br", "Hello", "Hello my friend")

	assert.NotNil(t, campaing.Email)
}

func Test_MustValidateFromNameWhenConfigureEmail(t *testing.T) {

	const errorExpected = "Name is required"
	const invalidFromName = ""
	campaing, _ := campaings.New(name, campaingType)

	err := campaing.ConfigureEmail(invalidFromName, "teste@teste.com.br", "Hello", "Hello my friend")

	assert.Equal(t, errorExpected, err.Error())
}

func Test_MustValidateFromEmailWhenConfigureEmail(t *testing.T) {

	const errorExpected = "Email is required"
	const invalidFromEmail = ""
	campaing, _ := campaings.New(name, campaingType)

	err := campaing.ConfigureEmail("teste", invalidFromEmail, "Hello", "Hello my friend")

	assert.Equal(t, errorExpected, err.Error())
}

func Test_MustValidateSubjectWhenConfigureEmail(t *testing.T) {

	const errorExpected = "Subject is required"
	const invalidSubject = ""
	campaing, _ := campaings.New(name, campaingType)

	err := campaing.ConfigureEmail("teste", "teste@teste.com.br", invalidSubject, "Hello my friend")

	assert.Equal(t, errorExpected, err.Error())
}

func Test_MustValidateBodyWhenConfigureEmail(t *testing.T) {

	const errorExpected = "Body is required"
	const invalidBody = ""
	campaing, _ := campaings.New(name, campaingType)

	err := campaing.ConfigureEmail("teste", "teste@teste.com.br", "Hello", invalidBody)

	assert.Equal(t, errorExpected, err.Error())
}
