package domain

//Text asdf
var Text string

//User alsdfjk
var User string

//Tweet struct tweet
type Tweet struct {
	Text string
	User string
	// Date string
}

//NewTweet un nuevo tweet
func NewTweet(user string, text string) *Tweet {
	myTweet := new(Tweet)
	myTweet.Text = text
	myTweet.User = user
	// myTweet.Date = time.Now()
	return myTweet
}
