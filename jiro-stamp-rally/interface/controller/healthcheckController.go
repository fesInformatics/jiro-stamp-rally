package controller

type HealthcheckController struct {}

func (c HealthcheckController) Get() Healthcheck {
	return Healthcheck {
		Message: "OK",
	}
}

func NewHealthcheckController () HealthcheckController {
	return HealthcheckController{}
}

type Healthcheck struct {
	Message string `json:"message"`
}
