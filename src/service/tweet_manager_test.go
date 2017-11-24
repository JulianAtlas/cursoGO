package service_test

import (
	"testing"

	"github.com/cursoGO/src/domain"
	"github.com/cursoGO/src/service"
)

func TestPublishedTweetIsSaved3(t *testing.T) {
	var tweet *domain.Tweet

	user := "meli-team"
	text := "los de ama-son unos giles"

	tweet = domain.NewTweet(user, test)

	publishedTweet := service.GetTweet()

	if publishedTweet.User != user && publishedTweet.Test != text {
		t.Errorf("Expected tweet is %s: %s \n but is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}
	// if publishedTweet.Date == nil{
	// 	t.Error("Expected date can´t be nil")
	// }
}
