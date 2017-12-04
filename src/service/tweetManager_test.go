package service_test

import (
	"testing"

	"github.com/cursoGO/src/domain"
	"github.com/cursoGO/src/service"
)

func TestSignUp(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")

	tm.SignUp(&user)

	if !tm.EstaRegistrado(&user) {
		t.Error("deberia estar registrado")
	}

	if len(tm.GetUsuariosRegistrados()) != 1 {
		t.Error("deberia haber 1 usuario en la lista de usuarios")
	}
}

func TestLogIn(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")
	tm.SignUp(&user)
	tm.LogIn(&user)

	if !tm.EstaLogueado(&user) {
		t.Error("Deberia estar logueado")
	}
}

func TestLogOut(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")
	tm.SignUp(&user)
	tm.LogIn(&user)
	tm.LogOut(&user)

	if tm.EstaLogueado(&user) {
		t.Error("Debería no estar logueado")
	}
}

func TestCanRetriveById(t *testing.T) {
	tm := service.NewTweetManager()
	var id int
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")
	text := "Tweet re loko"
	tweet := domain.NewTextTweet(user, text)
	tm.SignUp(&user)
	tm.LogIn(&user)
	id, _ = tm.PublishTweet(tweet)
	publishedTweet, _ := tm.GetTweetByID(id)

	if publishedTweet != tweet {
		t.Error("los tweets no son iguales")
	}
}

func TestGetTweets(t *testing.T) {
	tm := service.NewTweetManager()
	var tweet1 domain.Tweet
	var tweet2 domain.Tweet
	var tweet3 domain.Tweet

	var user1 domain.Usuario
	var user2 domain.Usuario

	user1.SetMail("mercadolibre.com")
	user1.SetUsername("meli-team")
	user2.SetMail("mercadolibre.com")
	user2.SetUsername("meli-team")

	text1 := "Texto1"
	text2 := "Texto2"
	text3 := "Texto3"

	tm.SignUp(&user1)
	tm.SignUp(&user2)

	tweet1 = domain.NewTextTweet(user1, text1)
	tweet2 = domain.NewTextTweet(user1, text2)
	tweet3 = domain.NewTextTweet(user2, text3)

	tm.LogIn(&user1)
	tm.LogIn(&user2)

	tm.PublishTweet(tweet1)
	tm.PublishTweet(tweet2)
	tm.PublishTweet(tweet3)

	tweets := tm.GetTweets()
	if len(tweets) != 3 {
		t.Error("deberian ser 3")
	}

}

func TestMap(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")

	text := "Tweet re loko"
	tweet := domain.NewTextTweet(user, text)

	tm.SignUp(&user)
	tm.LogIn(&user)
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
	tweet := domain.NewTextTweet(user, text)
	tm.SignUp(&user)
	tm.LogIn(&user)
	id, _ := tm.PublishTweet(tweet)

	//operation
	tm.RemoverTweet(id)

	//test
	if len(tm.GetTweets()) != 0 {
		t.Error("Hay demasiados Tweets")
	}
}

func TestEditarUnTweet(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")

	text := "Tweet re loko"
	tweet := domain.NewTextTweet(user, text)
	tm.SignUp(&user)
	tm.LogIn(&user)
	id, _ := tm.PublishTweet(tweet)

	nuevo_texto := "Tweet modificado"
	tm.EditarTweet(id, nuevo_texto)

	if twt, _ := tm.GetTweetByID(id); twt.GetText() != nuevo_texto {
		t.Error("No se modifico el texto del tweet")
	}
}

func TestNoPermitoTweetsDuplicados(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")

	text := "Tweet re loko"
	tweet := domain.NewTextTweet(user, text)
	tm.SignUp(&user)
	tm.LogIn(&user)
	tm.PublishTweet(tweet)

	if _, err := tm.PublishTweet(tweet); err == nil {
		t.Error("El tweet no se debería haber agregado")
	}
}

func TestSeguirUsuario(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")
	tm.SignUp(&user)
	tm.LogIn(&user)

	var user_que_sigue domain.Usuario
	user_que_sigue.SetMail("gmail.com")
	user_que_sigue.SetUsername("meli")
	tm.SignUp(&user_que_sigue)
	tm.LogIn(&user_que_sigue)

	tm.SeguirUsuario(&user_que_sigue, &user)

	if len(user_que_sigue.GetSeguidos()) != 1 {
		t.Error("Debería seguir a un user")
	}
	if len(user.GetSeguidores()) != 1 {
		t.Error("Debería seguirme un user")
	}
}

func TestTimeline(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")
	tm.SignUp(&user)
	tm.LogIn(&user)

	var user1 domain.Usuario
	user1.SetMail("mercadolibre.com")
	user1.SetUsername("meli-team")
	tm.SignUp(&user1)
	tm.LogIn(&user1)

	tm.SeguirUsuario(&user, &user1)

	text := "Tweet re loko"
	tweet := domain.NewTextTweet(user, text)

	text1 := "Tweet muy loko"
	tweet1 := domain.NewTextTweet(user1, text1)

	text2 := "Tweet realmente loko"
	tweet2 := domain.NewTextTweet(user1, text2)

	tm.PublishTweet(tweet)
	tm.PublishTweet(tweet1)
	tm.PublishTweet(tweet2)

	timeline, _ := tm.Timeline(&user)

	if len(timeline) != 3 {
		t.Error("Deberían ser 3 tweets")
	}
}

func TestRetweet(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")
	tm.SignUp(&user)
	tm.LogIn(&user)

	var user1 domain.Usuario
	user1.SetMail("mercadolibre.com")
	user1.SetUsername("meli-team")
	tm.SignUp(&user1)
	tm.LogIn(&user1)

	text := "Tweet re loko"
	tweet := domain.NewTextTweet(user, text)

	text1 := "Tweet muy loko"
	tweet1 := domain.NewTextTweet(user1, text1)

	tm.PublishTweet(tweet)
	tm.PublishTweet(tweet1)

	tm.Retweet(&user, tweet1.GetID())

	if len(tm.GetTweetsFromUser(&user)) != 2 {
		t.Error("Deberian ser 2")
	}
}

func TestFavoritos(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")
	tm.SignUp(&user)
	tm.LogIn(&user)

	var user1 domain.Usuario
	user1.SetMail("mercadolibre.com")
	user1.SetUsername("meli-team")
	tm.SignUp(&user1)
	tm.LogIn(&user1)

	text := "Tweet re loko"
	tweet := domain.NewTextTweet(user, text)

	text1 := "Tweet muy loko"
	tweet1 := domain.NewTextTweet(user1, text1)

	tm.PublishTweet(tweet)
	tm.PublishTweet(tweet1)

	tm.AgregarAFavoritos(&user, tweet1.GetID())

	if len(user.GetFavoritos()) != 1 {
		t.Error("No agrego el tweet a favoritos")
	}
}

func TestDameFavoritosDeUser(t *testing.T) {
	tm := service.NewTweetManager()
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")
	tm.SignUp(&user)
	tm.LogIn(&user)

	var user1 domain.Usuario
	user1.SetMail("mercadolibre.com")
	user1.SetUsername("meli-team")
	tm.SignUp(&user1)
	tm.LogIn(&user1)

	text := "Tweet re loko"
	tweet := domain.NewTextTweet(user, text)

	text1 := "Tweet muy loko"
	tweet1 := domain.NewTextTweet(user1, text1)

	tm.PublishTweet(tweet)
	tm.PublishTweet(tweet1)

	tm.AgregarAFavoritos(&user, tweet1.GetID())

	if res, _ := tm.DameFavoritos(&user); len(res) != 1 {
		t.Error("No me agrego el favorito")
	}
}
