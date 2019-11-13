package vm

import (
	"github.com/musketeer-liu/Go_Mega_Project/model"
)

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
// V3: Add middleware of User Auth
func (IndexViewModelOp) GetVM(username string) IndexViewModel {
	u1, _ := model.GetUserByUsername(username)
	posts, _ := model.GetPostsByUserID(u1.ID)
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *posts}
	v.SetCurrentUser(username)
	return v
}

//func (IndexViewModelOp) GetVM() IndexViewModel {
	//// V2: Database
	//u1, _ := model.GetUserByUsername("Musketeer")
	//posts, _ := model.GetPostsByUserID(u1.ID)
	//v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
	//return v

	//// V1: Hard Code
	//u1 := model.User{Username: "Musketeer"}
	//u2 := model.User{Username: "Paladin"}
	//
	//posts := []model.Post{
	//	model.Post{User: u1, Body: "Beautiful day in Portland!"},
	//	model.Post{User: u2, Body: "The Avengers movie was so cool!"},
	//}
	//v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
	//return v
//}
