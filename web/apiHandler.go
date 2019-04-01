package web

import (
	"go.uber.org/zap"
	"gramya/config"
	"gramya/model"
	"net/http"
)

var logger *zap.Logger = config.GetLogger()

func Expand(w http.ResponseWriter, r *http.Request) model.HttpResponse {
	if r := recover(); r != nil {
		return model.HttpResponse{200, nil, []byte("")}
	}
	logger.Info("HelLo Logger")
	return model.HttpResponse{200, nil, []byte("{}")}
}

func Profile(w http.ResponseWriter, r *http.Request) model.HttpResponse {
	if r := recover(); r != nil {
		return model.HttpResponse{200, nil, []byte("")}
	}
	return model.HttpResponse{200, nil, []byte("")}
}

func ParentChain(w http.ResponseWriter, r *http.Request) model.HttpResponse {
	if r := recover(); r != nil {
		return model.HttpResponse{200, nil, []byte("")}
	}
	return model.HttpResponse{200, nil, []byte("")}
}
