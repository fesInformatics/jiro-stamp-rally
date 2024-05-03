package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fesInformatics/jiro-stamp-rally/interface/context"
	"github.com/fesInformatics/jiro-stamp-rally/usecase/interactor"
)

type LoginController struct {
	interactor interactor.LoginInteractor
}

func (c LoginController) Login(ctx context.Context) {
	switch ctx.GetHttpMethod() {
	case http.MethodPost:
		var loginAction LoginAction
		if err := json.NewDecoder(ctx.GetBody()).Decode(&loginAction); err != nil {
			ctx.BadRequest(err)
		}

		err := c.interactor.Login(loginAction.MailAddress, loginAction.Password)
		if err != nil && errors.Is(err, interactor.ErrUserDuplicate) {
			ctx.BadRequest(err)
		} else if err != nil {
			ctx.InternalServerError(err)
		}

		ctx.JSON(http.StatusOK, nil)
	default:
		ctx.MethodNotAllowed(fmt.Errorf("MethodNotAllowed"))
	}
}

func NewLoginController(interactor interactor.LoginInteractor) LoginController {
	return LoginController{interactor: interactor}
}

type LoginAction struct {
	MailAddress string `json:"mailAddress"`
	Password    string `json:"password"`
}
