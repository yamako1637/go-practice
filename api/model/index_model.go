package model

type Index struct {
	Message string `json:"message"`
}

func GetIndexMessage() Index {
	return Index{
		Message: "Hello",
	}
}
