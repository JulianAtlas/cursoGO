package service_test

import (
	"testing"

	"github.com/cursoGO/src/domain"
	"github.com/cursoGO/src/service"
)

func TestCanRetriveById(t *testing.T) {
	tm := service.NewTweetManager()
	var id int
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")
	text := "Tweet re loko"
	tweet := domain.NewTweet(user, text)
	id, _ = tm.PublishTweet(tweet)
	publishedTweet := tm.GetTweetByID(id)

	if publishedTweet != tweet {
		t.Error("los tweets no son iguales")
	}
}

func TestGetTweets(t *testing.T) {
	tm := service.NewTweetManager()
	var tweet1 *domain.Tweet
	var tweet2 *domain.Tweet
	var tweet3 *domain.Tweet

	var user1 domain.Usuario
	var user2 domain.Usuario

	user1.SetMail("mercadolibre.com")
	user1.SetUsername("meli-team")
	user2.SetMail("mercadolibre.com")
	user2.SetUsername("meli-team")

	text1 := "Texto1"
	text2 := "Texto2"
	text3 := "Texto3"

	tm.SignUp(user1)
	tm.SignUp(user2)

	tweet1 = domain.NewTweet(user1, text1)
	tweet2 = domain.NewTweet(user1, text2)
	tweet3 = domain.NewTweet(user2, text3)

	tm.PublishTweet(tweet1)
	tm.PublishTweet(tweet2)
	tm.PublishTweet(tweet3)

	tweets := tm.GetTweets()
	if len(tweets) != 3 {
		t.Error("deberian ser 3")
	}

}

func TestSingUp(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")
	tm.SignUp(user)

	if !tm.EstaLogueado(user) {
		t.Error("deberia estar logeado")
	}

	if len(tm.GetUsuariosRegistrados()) < 1 {
		t.Error("deberia haber 1 usuario en la lista de usuarios")
	}

	if tm.GetUsuariosRegistrados()[0].GetMail() != "mercadolibre.com" {
		t.Error("Mail incorrecto")
	}
}

func TestMap(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")

	text := "Tweet re loko"
	tweet := domain.NewTweet(user, text)
	tm.PublishTweet(tweet)

	//test
	tweetTest := tm.GetTweetsPorUsuario()[user.GetID()]
	if tweetTest[0] != tweet {
		t.Error("Los tweets no son iguales")
	}
}

func TestRemoverTweet(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")

	text := "Tweet re loko"
	tweet := domain.NewTweet(user, text)
	id, _ := tm.PublishTweet(tweet)
	//operation
	tm.RemoverTweet(id)

	//test
	if len(tm.GetTweets()) > 0 {
		t.Error("Hay demasiados Tweets")
	}
}
