package handler

type DataResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func NewDataResponseWithError(err error) DataResponse {
	return DataResponse{Success: false, Error: err.Error()}
}

func NewDataResponseWithServerError() DataResponse {
	return DataResponse{Success: false, Error: "An error occurred"}
}

func NewDataResponse() DataResponse {
	return DataResponse{Success: true}
}
