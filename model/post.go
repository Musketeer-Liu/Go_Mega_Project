package model

// Post struct
type Post struct {
	User // 这里可以采用匿名简化 不用写User User HTML模板中也可以不用再写.User了
	// User User
	Body string
}
