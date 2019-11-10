package vm

import (
	"github.com/musketeer-liu/Go_Mega_Project/model"
)

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM() IndexViewModel {
	// Read data from Database rather than Hard Code
	// Database
	u1, _ := model.GetUserByUsername("Musketeer")
	posts, _ := model.GetPostsByUserID(u1.ID)

	//// Hard Code
	//u1 := model.User{Username: "Musketeer"}
	//u2 := model.User{Username: "Paladin"}

	//posts := []model.Post{
	//	model.Post{User: u1, Body: "Beautiful day in Portland!"},
	//	model.Post{User: u2, Body: "The Avengers movie was so cool!"},
	//}

	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
	return v
}
