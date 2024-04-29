package modHttp

import (
	"domcloud/modUtility"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CChi struct {
	router *chi.Mux
}

var g_singleChi *CChi = &CChi{}

func getSingleChi() *CChi {
	return g_singleChi
}

func (pInst *CChi) Initialize() error {
	r := chi.NewRouter()
	r.Get("/", handlerHomepage)
	r.Get("/apilist", handlerApiList)
	pInst.router = r

	return nil
}

func (pInst *CChi) start() error {
	listenStr := fmt.Sprintf(":%d", modUtility.G_HttpPort)
	fmt.Println("start listen: ", listenStr)
	err := http.ListenAndServe(listenStr, pInst.router)
	return err
}
