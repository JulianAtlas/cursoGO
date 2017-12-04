package domain

import (
	"time"
)

//TextTweet struct tweet
type TextTweet struct {
	Text string
	User *Usuario
	Date *Datestruct
	ID   int
}

//NewTextTweet un nuevo tweet
func NewTextTweet(user Usuario, text string) *TextTweet {
	myTweet := new(TextTweet)
	myTweet.Text = text
	myTweet.User = &user
	fecha := time.Now()
	year := fecha.Format("2006")
	month := fecha.Format("01")
	day := fecha.Format("02")
	myTweet.Date = &Datestruct{year, month, day}
	return myTweet
}

//GetText getter text
func (tw *TextTweet) GetText() string {
	return tw.Text
}

//GetID getter id
func (tw *TextTweet) GetID() int {
	return tw.ID
}

//GetUser getter user
func (tw *TextTweet) GetUser() *Usuario {
	return tw.User
}

//SetID setter id
func (tw *TextTweet) SetID(id int) {
	tw.ID = id
}

//SetText setter de text
func (tw *TextTweet) SetText(nuevoTexto string) {
	tw.Text = nuevoTexto
}
