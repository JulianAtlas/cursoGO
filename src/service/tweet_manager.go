package service

//Tweet una variable re loka asd
var tweet string

//PublishTweet functionasd
func PublishTweet(tweetToPublish string) {
	tweet = tweetToPublish
}

//GetTweet getter de tweet
func GetTweet() string {
	return tweet
}
