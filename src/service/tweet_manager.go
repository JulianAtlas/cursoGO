package service

import "github.com/cursoGO/src/domain"

var myTweet *domain.Tweet

//PublishTweet publisher
func PublishTweet(unTweet *domain.Tweet) {
	myTweet = unTweet
}

//GetTweet getter
func GetTweet() *domain.Tweet {
	return myTweet
}
