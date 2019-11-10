package vm

import (
	"log"

	"github.com/musketeer-liu/Go_Mega_Project/model"
)

// LoginViewModel struct
type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

// LoginViewModelOp struct
type LoginViewModelOp struct {}

// GetVM func
func (LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}

// AddError func
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

// CheckLogin func
func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Cannot find username: ", username)
		log.Println("Error", err)
		return false
	}
	return user.CheckPassword(password)
}