package vm

import (
	"github.com/musketeer-liu/Go_Mega_Project/model"
	"strings"
)

// ProfileViewModel struct
type ProfileViewModel struct {
	BaseViewModel
	Posts				[]model.Post
	Editable			bool
	IsFollow			bool
	FollowersCount		int
	FollowingCount		int
	ProfileUser			model.User
	BasePageViewModel
}

// ProfileViewModelOp struct
type ProfileViewModelOp struct {}

// GetVM func
func (ProfileViewModelOp) GetVM(sUser, pUser string, page, limit int) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, total, _ := model.GetPostsByUserIDPageAndLimit(u.ID, page, limit)
	v.ProfileUser = *u
	v.Editable = (sUser == pUser)
	v.SetBasePageViewModel(total, page, limit)
	if !v.Editable {
		v.IsFollow = u.IsFollowedByUser(sUser)
	}
	v.FollowersCount = u.FollowersCount()
	v.FollowingCount = u.FollowingCount()

	v.Posts = *posts
	v.SetCurrentUser(sUser)
	return v, nil
}

// Follow func: A follow B
func Follow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}
	return u.Follow(b)
}

// UnFollow func: A unfollow B
func UnFollow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}
	return u.Unfollow(b)
}

// 由于Popup的 viewmodel 和 Profile 的相似，我们直接在 vm/profile.go中加入 GetPopupVM 来获得 Popup 的 vm
// GetPopupVM func
func (ProfileViewModelOp) GetPopupVM(sUser, pUser string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	v.ProfileUser = *u
	v.Editable = (strings.ToLower(sUser) == strings.ToLower(pUser))
	if !v.Editable {
		v.IsFollow = u.IsFollowedByUser(sUser)
	}
	v.FollowersCount = u.FollowersCount()
	v.FollowingCount = u.FollowingCount()
	v.SetCurrentUser(sUser)
	return v, nil
}









