package model

type ErrorMessage struct {
	Message string `json:"message"`
}

func GetErrorMessage() ErrorMessage {
	return ErrorMessage{
		Message: "Error",
	}
}
