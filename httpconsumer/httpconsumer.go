package httpconsumer

import (
	"fmt"

	"github.com/consume-rest-api/util"
)

type CatFact struct {
	Fact string `json:"fact"`
	Length int `json:"length"`
}

type RandomUser struct {
	Results []UserResult
}

type UserResult struct {
	Name UserName
	Email string
	Picture UserPicture
}

type UserName struct {
	Title string
	First string
	Last string
}

type UserPicture struct {
	Large string
	Medium string
	Thumbnail string
}

func GetCatFact() {
	url := "https://catfact.ninja/fact"

	var catFact CatFact 
	err := util.GetJson(url, &catFact)
	if err != nil {
		fmt.Printf("error getting cat fact: %s\n", err.Error())
		return
	} else {
		fmt.Printf("A super interesting Cat Fact: %s \n", catFact.Fact)
	}
}

func GetRandomUser() {
	url := "https://randomuser.me/api/?inc=name,email,picture"

	var user RandomUser

	err := util.GetJson(url, &user)
	if err != nil {
		fmt.Printf("error getting json: %s\n", err.Error())
	} else {
		fmt.Printf("user: %s %s %s \nEmail: %s\nThumbnail: %s\n", 
		user.Results[0].Name.Title,
		user.Results[0].Name.First,
		user.Results[0].Name.Last,
		user.Results[0].Email,
		user.Results[0].Picture.Thumbnail)	
	}
}