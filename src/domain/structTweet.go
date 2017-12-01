package domain

import (
	"time"
)

//Tweet struct tweet
type TextTweet struct {
	text string
	user Usuario
	date *Datestruct
	id   int
}

//NewTweet un nuevo tweet
func NewTextTweet(user Usuario, text string) *TextTweet {
	myTweet := new(TextTweet)
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
func (tw *TextTweet) GetText() string {
	return tw.text
}

//GetID getter id
func (tw *TextTweet) GetID() int {
	return tw.id
}

//GetUser getter user
func (tw *TextTweet) GetUser() *Usuario {
	return &tw.user
}

//SetID setter id
func (tw *TextTweet) SetID(id int) {
	tw.id = id
}

//SetText setter de text
func (tw *TextTweet) SetText(nuevoTexto string) {
	tw.text = nuevoTexto
}
