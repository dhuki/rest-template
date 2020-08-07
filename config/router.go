package config

import (
	"fmt"
	"net/http"

	"github.com/dhuki/rest-template/common"
	"github.com/gorilla/mux"
)

type router struct {
	Mux *mux.Router
}

func NewRouter() router {
	return router{
		// it will return router under host common.BaseUrl
		Mux: mux.NewRouter().PathPrefix(common.BaseUrl).Subrouter(),
	}
}

func (r router) Start() error {
	// using pointer bcs receiver is pointer
	// actually it's okay to use not pointer even receiver is pointer
	// bcs this struct not return an interface
	// but if struct return an interface you should return as pointer if it's not it will error
	srv := &http.Server{
		Handler: r.Mux,
		Addr:    fmt.Sprintf("%s:%s", common.Host, common.Port),
	}
	return srv.ListenAndServe()
}

// function to get to know that available router
// it is directly from mux router doc
func (r router) GetListRouterAvailable() {
	r.Mux.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
}
