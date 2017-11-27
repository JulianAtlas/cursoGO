package domain

import (
	"time"
)

//Text asdf
var Text string

//User alsdfjk
var User string

//ID unico por tweet
var ID int

//Tweet struct tweet
type Tweet struct {
	Text string
	User Usuario
	Date *Datestruct
	ID   int
}

//NewTweet un nuevo tweet
func NewTweet(user Usuario, text string) *Tweet {
	myTweet := new(Tweet)
	myTweet.Text = text
	myTweet.User = user
	fecha := time.Now()
	year := fecha.Format("2006")
	month := fecha.Format("01")
	day := fecha.Format("02")
	myTweet.Date = &Datestruct{year, month, day}
	return myTweet
}
