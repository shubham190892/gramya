package web

import (
	"gramya/model"
	"net/http"
)

func Expand(w http.ResponseWriter, r *http.Request) model.HttpResponse {
	return model.HttpResponse{200, nil, []byte("")}
}

func Profile(w http.ResponseWriter, r *http.Request) model.HttpResponse {
	return model.HttpResponse{200, nil, []byte("")}
}

func ParentChain(w http.ResponseWriter, r *http.Request) model.HttpResponse {
	return model.HttpResponse{200, nil, []byte("")}
}
