package handler

type DataResponse struct {
	Error string `json:"error"`
}

func NewDataResponseWithError(err error) DataResponse {
	return DataResponse{Error: err.Error()}
}

func NewDataResponseWithServerError() DataResponse {
	return DataResponse{Error: "An error occurred"}
}
