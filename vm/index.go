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
type IndexViewModelOp struct {}

// GetVM func
func (IndexViewModelOp) GetVM() IndexViewModel {
	u1 := User{Username: "Musketeer"}
	u2 := User{Username: "Paladin"}

	posts := []Post{
		Post{User: u1, Body: "Beautiful day in Portland!"},
		Post{User: u2, Body: "The Avengers movie was so cool!"},
	}

	v := IndexViewModel{Title: "Homepage", User: u1, Posts: posts}
	return v
}
