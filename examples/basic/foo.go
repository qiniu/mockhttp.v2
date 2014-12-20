package foo

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// --------------------------------------------------------------------

func reply(w http.ResponseWriter, code int, data interface{}) {

	msg, _ := json.Marshal(data)
	h := w.Header()
	h.Set("Content-Length", strconv.Itoa(len(msg)))
	h.Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(msg)
}

// --------------------------------------------------------------------

type FooRet struct {
	A int    `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

type HandleRet map[string]string

// --------------------------------------------------------------------

type Service struct{}

func (p *Service) foo(w http.ResponseWriter, req *http.Request) {
	reply(w, 200, &FooRet{1, req.Host, req.URL.Path})
}

func (p *Service) handle(w http.ResponseWriter, req *http.Request) {
	reply(w, 200, HandleRet{"foo": "1", "bar": "2"})
}

func (p *Service) RegisterRoute() {
	http.HandleFunc("/foo", p.foo)
	http.HandleFunc("/", p.handle)
}

// --------------------------------------------------------------------
