package vm

import (
	"log"

	"github.com/musketeer-liu/Go_Mega_Project/model"
)

// RegisterVeiwModel struct
type RegisterViewModel struct {
	LoginViewModel
}

// RegisterViewModelOp struct
type RegisterViewModelOp struct {}

// GetVM func
func (RegisterViewModelOp) GetVM() RegisterViewModel {
	v := RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

// CheckUserExist func
func CheckUserExist(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Cannot find username: ", username)
		return true
	}
	return false
}

// addUser func
func AddUser(username, password, email string) error {
	return model.AddUser(username, password, email)
}
