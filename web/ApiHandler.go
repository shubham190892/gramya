package web

import (
	"encoding/json"
	"gramya/constants"
	"gramya/model"
	"gramya/services"
	"gramya/utils"
	"net/http"
)

const (
	UnableToParseJson = "Unable to parse JSON"
	MissingHeaders    = "Missing Headers"
)

func Expand(r *http.Request) (mr model.HttpResponse) {
	id := utils.QueryParams(r.RequestURI)["id"]
	return model.HttpResponse{200, nil, []byte("{}")}
}

func Profile(r *http.Request) model.HttpResponse {
	return model.HttpResponse{200, nil, []byte("")}
}

func ParentChain(r *http.Request) model.HttpResponse {
	return model.HttpResponse{200, nil, []byte("")}
}

func Welcome(r *http.Request) model.HttpResponse {
	if r.Header.Get(constants.ContentType) != constants.ApplicationJson {
		return model.HttpResponse{Status: 400, Data: []byte(MissingHeaders)}
	}
	var welcomeRequest model.WelcomeRequest
	err := json.NewDecoder(r.Body).Decode(&welcomeRequest)
	if !welcomeRequest.Validate() {
		return model.HttpResponse{Status: 400, Data: []byte("Required Data missing")}
	}
	if err != nil {
		utils.Logger.Error(err.Error())
		return model.HttpResponse{500, nil, []byte(UnableToParseJson)}
	}
	status := services.AddPerson(welcomeRequest)
	res, _ := json.Marshal(model.ResponseOk{status, "Ok"})
	return model.HttpResponse{200, nil, res}
}
