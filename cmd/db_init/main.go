package main

import (
	"log"
	//"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/musketeer-liu/Go_Mega_Project/model"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{}, "follower")
	db.CreateTable(model.User{}, model.Post{})

	model.AddUser("Musketeer", "abc123", "musketeer.liu@gmail.com")
	model.AddUser("Paladin", "abc123", "paladin@test.com")

	u1, _ := model.GetUserByUsername("musketeer")
	u1.CreatePost("Beautiful day in Portland")
	u1.CreatePost("Beautiful day in California")
	u1.CreatePost("Beautiful day in Texas")
	u1.CreatePost("Beautiful day in New York")
	u1.CreatePost("Beautiful day in Arizona")
	u1.CreatePost("Beautiful day in Illinois")
	u1.CreatePost("Beautiful day in Shanghai")
	u1.CreatePost("Beautiful day in Beijing")
	u1.CreatePost("Beautiful day in Guangdong")
	u1.CreatePost("Beautiful day in Hongkong")
	model.UpdateAboutMe(u1.Username, `I'm the author of this Go Mega Project`)

	u2, _ := model.GetUserByUsername("paladin")
	u2.CreatePost("The Avengers movie was so cool!")
	u2.CreatePost("Sun shine is beautiful")

	u1.Follow(u2.Username)

	//db.DropTableIfExists(model.User{}, model.Post{})
	//db.CreateTable(model.User{}, model.Post{})

	//users := []model.User{
	//	{
	//		Username:		"Musketeer",
	//		PasswordHash:	model.GeneratePasswordHash("abc123"),
	//		Email:			"musketeer@test.com",
	//		Avatar:			fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("musketeer@test.com")),
	//		Posts: 			[]model.Post{
	//			{Body: "Beautiful day in Portland!"},
	//		},
	//	},
	//	{
	//		Username:     "Paladin",
	//		PasswordHash: model.GeneratePasswordHash("abc123"),
	//		Email:        "paladin@test.com",
	//		Avatar:			fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("paladin@test.com")),
	//		Posts: []model.Post{
	//			{Body: "The Avengers movie was so cool!"},
	//			{Body: "Sun shine is beautiful"},
	//		},
	//	},
	//}

	//for _, u := range users {
	//	db.Debug().Create(&u)
	//}
}



