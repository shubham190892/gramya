package model

type HttpResponse struct {
	Status  int
	Headers map[string]string
	Data    []byte
}
