package vm

import (
	"github.com/musketeer-liu/Go_Mega_Project/model"
)

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	Posts			[]model.Post
	Flash			string

	BasePageViewModel
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
// V3: Add middleware of User Auth
func (IndexViewModelOp) GetVM(username string, flash string, page, limit int) IndexViewModel {
	u, _ := model.GetUserByUsername(username)
	// 顺便将 IndexView 里的 Posts 改成 CurrentUser 的 FollowingPosts
	posts, total, _ := u.FollowingPostsByPageAndLimit(page, limit)
	v := IndexViewModel{}
	v.SetTitle("Homepage")
	v.Posts = *posts
	v.Flash = flash
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v
}

// CreatePost func
func CreatePost(username, post string) error {
	u, _ := model.GetUserByUsername(username)
	return u.CreatePost(post)
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
