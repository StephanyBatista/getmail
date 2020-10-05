package fake

import (
	"bytes"
	"getmail/httpd"
	"getmail/infra/data"
	"net/http"
	"net/http/httptest"
)

func NewJsonRequest(repo data.IRepository, method string, url string, json []byte) (code int, body string) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(json))
	req.Header.Add("content-type", "application/json")

	router := httpd.RegisterHTTPHandlers(repo)
	router.ServeHTTP(w, req)

	return w.Code, w.Body.String()
}
