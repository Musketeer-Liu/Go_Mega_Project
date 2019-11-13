package vm

import (
	"github.com/musketeer-liu/Go_Mega_Project/model"
)

// ExploreViewModel struct
type ExploreViewModel struct {
	BaseViewModel
	Posts				[]model.Post
	BasePageViewModel
}

// ExploreViewModelOp struct
type ExploreViewModelOp struct {}

// GetVM func
func (ExploreViewModelOp) GetVM(username string, page, limit int) ExploreViewModel {
	// posts, _ := model.GetAllPosts()
	posts, total, _ := model.GetPostsByPageAndLimit(page, limit)
	v := ExploreViewModel{}
	v.SetTitle("Explore")
	v.Posts = *posts
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v
}
