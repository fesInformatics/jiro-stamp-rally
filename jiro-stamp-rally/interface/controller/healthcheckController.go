package controller

import (
	"github.com/fesInformatics/jiro-stamp-rally/interface/context"
)

type HealthcheckController struct {}

func (c HealthcheckController) Get(ctx context.Context) {
	r := Healthcheck {
		Message: "OK",
	}
	ctx.JSON(r)
}

func NewHealthcheckController () HealthcheckController {
	return HealthcheckController{}
}

type Healthcheck struct {
	Message string `json:"message"`
}
