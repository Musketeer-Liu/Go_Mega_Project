package model

import (
	"fmt"
	"time"
)

// User struct
// 由于 Go 原生的Template 虽然是支持自定义函数的 但是不像 Jinja2 那么好支持函数
// 需要在template.New().Funcs() 中预先传入，与我们已有的 PopulateTemplates函数集成上有点难度
// Go Template 支持类的 Func，不用预先传入，这里 Avatar 字段不是特别的必要
// 直接将 Avatar 作为字段放入数据库中，等于冗余了Avatar数据，但是减少了我们coding的难度
// 在增加User 的时候，直接设置Avatar，并加入AboutMe、LastSeen 字段
type User struct {
	ID				int			`gorm:"primary_key"`
	Username		string		`gorm:"type:varchar(64)"`
	Email			string		`gorm:"type:varchar(120)"`
	PasswordHash	string		`gorm:"type:varchar(128)"`
	LastSeen		*time.Time
	AboutMe			string		`gorm:"type:varchar(140)"`
	Avatar			string		`gorm:"type:varchar(200)"`
	Posts			[]Post
	Followers		[]*User		`gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

// SetPassword func: Set PasswordHash
func (u *User) SetPassword(password string) {
	u.PasswordHash = GeneratePasswordHash(password)
}

// CheckPassword func
func (u *User) CheckPassword(password string) bool {
	return GeneratePasswordHash(password) == u.PasswordHash
}

// GetUserByUsername func
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// SetAvatar func
func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprint("https://www.gravatar.com/avatar/%s?d=identicon", Md5(email))
}

// AddUser func
func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPassword(password)
	user.SetAvatar(email)
	return db.Create(&user).Error
}

// UpdateUserbyUsername func
func UpdateUserByUsername(username string, contents map[string]interface{}) error {
	item, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(item).Updates(contents).Error
}

// UpdateLastSeen func
func UpdateLastSeen(username string) error {
	contents := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserByUsername(username, contents)
}

// UpdateAboutMe func
func UpdateAboutMe(username, text string) error {
	contents := map[string]interface{}{"about_me": text}
	return UpdateUserByUsername(username, contents)
}
