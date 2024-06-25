package modHttp

import (

	//"io"
	"net/http"
	//"github.com/go-chi/chi/v5"
)

func handlerHomepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))

}

func handlerApiList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nothing"))

}
