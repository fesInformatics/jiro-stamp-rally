package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fesInformatics/jiro-stamp-rally/interface/context"
	"github.com/fesInformatics/jiro-stamp-rally/usecase/interactor"
)

type UserController struct {
	interactor interactor.UserInteractor
}

func (c UserController) Create(ctx context.Context) {
	switch ctx. GetHttpMethod(){
	case http.MethodPost:
		var user UserCreate
		if err := json.NewDecoder(ctx.GetBody()).Decode(&user); err != nil {
			fmt.Printf("エラーで")
		}
		err := c.interactor.Save(user.MailAddress, user.Password)
		fmt.Println(err)
	default:
		fmt.Printf("許可されていないメソットです")
	}
}

func NewUserController (interactor interactor.UserInteractor) UserController {
	return UserController{ interactor: interactor}
}

type UserCreate struct {
	MailAddress string `json:"mailAddress"`
	Password string `json:"password"`
}
