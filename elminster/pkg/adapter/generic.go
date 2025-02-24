package adapter

type GenericResponseAdapter struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
