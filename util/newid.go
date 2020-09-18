package util

import guuid "github.com/google/uuid"

//NewID generates a new ID with 8 chars
func NewID() string {

	guid := guuid.New().String()
	return guid[0:8]
}
