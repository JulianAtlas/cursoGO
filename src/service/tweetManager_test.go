package service_test

import (
	"testing"

	"github.com/cursoGO/src/domain"
	"github.com/cursoGO/src/service"
)

func TestWithoutTextIsNotPublished(t *testing.T) {

	var tweet *domain.Tweet
	var user domain.Usuario
	user.Mail = "mercadolibre.com"
	user.Username = "meli-team"
	var text string

	tweet = domain.NewTweet(user, text)

	var err error
	_, err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestWithMultipleTweets(t *testing.T) {
	service.InitializeService()
	var id int
	var user domain.Usuario
	user.Mail = "mercadolibre.com"
	user.Username = "meli-team"
	text := "Tweet re loko"
	text2 := "Tweet no tan loko"

	tweet := domain.NewTweet(user, text)
	tweet2 := domain.NewTweet(user, text2)

	id, _ = service.PublishTweet(tweet)
	id, _ = service.PublishTweet(tweet2)

	publishedTweets := service.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("expected size was 2, but was %d ", len(publishedTweets))
		return
	}

	publishedTweet := publishedTweets[0]
	publishedTweet2 := publishedTweets[1]

	if !isValidTweet(t, publishedTweet, id, user, text) {
		return
	}

	if !isValidTweet(t, publishedTweet2, id, user, text) {
		return
	}
}

func isValidTweet(t *testing.T, publishedTweet *domain.Tweet, id int, user domain.Usuario, text string) bool {
	respuesta := true
	if publishedTweet == nil && publishedTweet.ID != id && publishedTweet.User != user &&
		publishedTweet.Text != text {
		respuesta = false
	}
	return respuesta
}

func TestCanRetriveById(t *testing.T) {
	service.InitializeService()
	var id int
	var user domain.Usuario
	user.Mail = "mercadolibre.com"
	user.Username = "meli-team"
	text := "Tweet re loko"
	tweet := domain.NewTweet(user, text)
	id, _ = service.PublishTweet(tweet)
	publishedTweet := service.GetTweetByID(id)
	if !isValidTweet(t, publishedTweet, id, user, text) {
		t.Error("Trajiste un tweet cualca")
	}
}

func TestSignOnDeUnUsuarioLoAgregaALosRegistrados(t *testing.T) {
	//var id int
	var user domain.Usuario
	user.Mail = "mercadolibre.com"
	user.Username = "meli-team"

	//fmt.Println(len(service.UsuariosRegistrados))
	service.SignUp(user)
	//fmt.Println(len(service.UsuariosRegistrados))
	if service.EstaLogueado(user) != nil {
		t.Error("el usuario no se logueo")
	}
}

func TestMap(t *testing.T) {
	var user domain.Usuario
	user.Mail = "mercadolibre.com"
	user.Username = "meli-team"

	text := "Tweet re loko"
	tweet := domain.NewTweet(user, text)
	id, _ = service.PublishTweet(tweet)

	//test
	tweetTest := service.tweetsPorUsuario[user.ID]
	if tweetTest != tweet {
		t.Error("Los tweets no son iguales")
	}
}
