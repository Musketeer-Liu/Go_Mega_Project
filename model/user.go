package model

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return user.FollowSelf()
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

// Date and Time
// FormattedLastSeen func
func (u *User) FormattedLastSeen() string {
	return u.LastSeen.Format("2006-01-02 15:04:05")
}

// UpdateAboutMe func
func UpdateAboutMe(username, text string) error {
	contents := map[string]interface{}{"about_me": text}
	return UpdateUserByUsername(username, contents)
}


// Follow func
// follow someone user_id other.id follow_id u.id
// 这里 Followers 可以通过 Gorm 的 Association("Followers")来实现， Following 好像不支持 需要自己实现
func (u *User) Follow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Append(u).Error
}

// Unfollow func
func (u *User) Unfollow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Delete(u).Error
}

// FollowSelf func
func (u *User) FollowSelf() error {
	return db.Model(u).Association("Followers").Append(u).Error
}

// FollowersCount func
func (u *User) FollowersCount() int {
	return db.Model(u).Association("Followers").Count()
}

// FollowingIDs func
func (u *User) FollowingIDs() []int {
	var ids []int
	rows, err := db.Table("follower").Where("follower_id=?", u.ID).Select("user_id, follower_id").Rows()
	if err != nil {
		log.Println("Counting Following error:", err)
		return ids
	}
	defer rows.Close()
	for rows.Next() {
		var id, followerID int
		rows.Scan(&id, &followerID)
		ids = append(ids, id)
	}
	return ids
}

// FollowingCount func
func (u *User) FollowingCount() int {
	ids := u.FollowingIDs()
	return len(ids)
}

// FollowingPosts func
func (u *User) FollowingPosts() (*[]Post, error) {
	var posts []Post
	ids := u.FollowingIDs()
	if err := db.Preload("User").Order("timestamp desc").Where("user_id in (?)", ids).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

// IsFollowedByUser func
func (u *User) IsFollowedByUser(username string) bool {
	user, _ := GetUserByUsername(username)
	ids := user.FollowingIDs()
	for _, id := range ids {
		if u.ID == id {
			return true
		}
	}
	return false
}

// CreatePostfunc
func (u *User) CreatePost(body string) error {
	post := Post{Body: body, UserID: u.ID}
	return db.Create(&post).Error
}

// FollowingPostsByPageAndLimit func
func (u *User) FollowingPostsByPageAndLimit(page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (page - 1) * limit
	ids := u.FollowingIDs()
	if err := db.Preload("User").Order("timestamp desc").Where("user_id in (?)", ids).Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Where("user_id in (?)", ids).Count(&total)
	return &posts, total, nil
}

// GenerateToken func
// 密钥 secret 这里直接写在代码里，其实更优还是通过配置文件配置
func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":	u.Username,
		"exp":		time.Now().Add(time.Hour * 2).Unix(),	//可以添加过期时间
	})
	return token.SignedString([]byte("secret"))
}

// CheckToken func
func CheckToken(tokenString string) (string, error) {
	toekn, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secret"), nil
	})

	if claims, ok := toekn.Claims.(jwt.MapClaims); ok && toekn.Valid {
		return claims["username"].(string), nil
	} else {
		return "", err
	}
}

// GetUserByEmail func
func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := db.Where("email=?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdatePassword func
func UpdatePassword(username, password string) error {
	contents := map[string]interface{}{"password_hash": Md5(password)}
	return UpdateUserByUsername(username, contents)
}

