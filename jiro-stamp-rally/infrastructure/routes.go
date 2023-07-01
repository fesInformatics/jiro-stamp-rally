package infrastructure

import (
	"encoding/json"
	"github.com/fesInformatics/jiro-stamp-rally/jiro-stamp-rally/interface/controller"
	"time"
	"net/http"
	"fmt"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	controller := controller.NewHealthcheckController()
	
	if err := json.NewEncoder(w).Encode(controller.Get()); err != nil {
		fmt.Printf("エンコードエラー")
	}
}

func Run() {


	http.HandleFunc("/health", HealthCheck)

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
