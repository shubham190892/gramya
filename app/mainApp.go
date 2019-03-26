package main

import (
	"fmt"
	"gramya/model"
	"gramya/web"
	"log"
	"net/http"
)

type Response = model.HttpResponse

func main() {
	server()
}

func handleGet(w http.ResponseWriter, r *http.Request) Response {
	switch r.URL.Path {
	case "/gram-service/expand":
		return web.Expand(w, r)
	case "/gram-service/profile":
		return web.Profile(w, r)
	case "/gram-service/parentChain":
		return web.ParentChain(w, r)
	default:
		return Response{Status: 404, Headers: nil, Data: []byte("")}
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) Response {
	switch r.URL.Path {
	case "/gramya/testPost":
		return Response{200, nil, []byte("Post Test OK")}
	default:
		return Response{404, nil, []byte("")}
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	var res Response

	switch r.Method {
	case http.MethodGet:
		res = handleGet(w, r)
	case http.MethodPost:
		res = handlePost(w, r)
	default:
		res = Response{Status: 405, Headers: nil, Data: []byte("")}
	}

	w.WriteHeader(res.Status)
	if res.Headers != nil {
		for k, v := range res.Headers {
			w.Header().Set(k, v)
		}
	}
	_, err := w.Write(res.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func server() {
	port := ":8383"
	http.HandleFunc("/", defaultHandler)
	fmt.Printf("Http server listening at %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
