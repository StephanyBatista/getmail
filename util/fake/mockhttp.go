package fake

import (
	"bytes"
	"getmail/httpd"
	"net/http"
	"net/http/httptest"
)

var router = httpd.RegisterHTTPHandlers()

func NewJsonRequest(method string, url string, json []byte) (code int, body string) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(json))
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)

	return w.Code, w.Body.String()
}
