package vm

import (
	"log"

	"github.com/musketeer-liu/Go_Mega_Project/model"
)

// ResetPasswordReqeustViewModel struct
type ResetPasswordRequestViewModel struct {
	LoginViewModel
}

// ResetPasswordRequestViewModelOp struct
type ResetPasswordRequestViewModelOp struct {}

// getVM func
func (ResetPasswordRequestViewModelOp) GetVM() ResetPasswordRequestViewModel {
	v := ResetPasswordRequestViewModel{}
	v.SetTitle("Forget Password")
	return v
}

// CheckEmailExistfunc
func CheckEmailExist(email string) bool {
	_, err := model.GetUserByEmail(email)
	if err != nil {
		log.Println("Cannot find email:", email)
		return false
	}
	return true
}

