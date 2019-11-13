package vm

import (
	"github.com/musketeer-liu/Go_Mega_Project/model"
)

// ProfileViewModel struct
type ProfileViewModel struct {
	BaseViewModel
	Posts			[]model.Post
	ProfileUser		model.User
}

// ProfileViewModelOp struct
type ProfileViewModelOp struct {}

// GetVM func
func (ProfileViewModelOp) GetVM(sUser, pUser string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u1, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, _ := model.GetPostsByUserID(u1.ID)
	v.ProfileUser = *u1
	v.Posts = *posts
	v.SetCurrentUser(sUser)
	return v, nil
}
