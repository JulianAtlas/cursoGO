package domain

import (
	"time"
)

//Tweet struct tweet
type Tweet struct {
	text string
	user Usuario
	date *Datestruct
	id   int
}

//NewTweet un nuevo tweet
func NewTweet(user Usuario, text string) *Tweet {
	myTweet := new(Tweet)
	myTweet.text = text
	myTweet.user = user
	fecha := time.Now()
	year := fecha.Format("2006")
	month := fecha.Format("01")
	day := fecha.Format("02")
	myTweet.date = &Datestruct{year, month, day}
	return myTweet
}

//GetText getter text
func (tw *Tweet) GetText() string {
	return tw.text
}

//GetID getter id
func (tw *Tweet) GetID() int {
	return tw.id
}

//GetUser getter user
func (tw *Tweet) GetUser() *Usuario {
	return &tw.user
}

//SetID setter id
func (tw *Tweet) SetID(id int) {
	tw.id = id
}
