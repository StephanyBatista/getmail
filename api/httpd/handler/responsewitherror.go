package handler

import "getmail/domain"

type DataResponse struct {
	Error string `json:"error"`
}

//ResponseWithError return the status code and the error message to client
func ResponseWithError(err error) (int, DataResponse) {

	if _, ok := err.(*domain.Error); ok {
		return 400, DataResponse{Error: err.Error()}
	}
	return 500, DataResponse{Error: "An error occurred"}
}
