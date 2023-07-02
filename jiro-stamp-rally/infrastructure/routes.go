package infrastructure

import (
	"github.com/fesInformatics/jiro-stamp-rally/jiro-stamp-rally/interface/controller"
	"github.com/fesInformatics/jiro-stamp-rally/jiro-stamp-rally/interface/context"
	"time"
	"net/http"
	"fmt"
)

func Run() {
	http.HandleFunc("/health", healthCheck)

	s := http.Server{
		Addr: ":8080",
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout: 90 * time.Second,
	}

	fmt.Println("Server Start")

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	controller := controller.NewHealthcheckController()
	ctx := context.NewContext(w, r)
	controller.Get(ctx)
}
