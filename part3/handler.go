package practical

import "net/http"

type HelloWorldHandler struct{}

func (h HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
