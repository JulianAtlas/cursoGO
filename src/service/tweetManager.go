package service

import (
	"fmt"

	"github.com/cursoGO/src/domain"
)

//myTweets son todos mis tweets
//var myTweets []*domain.Tweet
var tweetsPorUsuario map[int][]*domain.Tweet

//usuariosRegistrados todos los usuarios registrados
var UsuariosRegistrados []*domain.Usuario

//EstaLogueado consulto si el usuario está logueado
func EstaLogueado(unUsuario domain.Usuario) error {
	respuesta := fmt.Errorf("El usuario no existe")
	for _, user := range UsuariosRegistrados {
		if *user == unUsuario {
			respuesta = nil
		}
	}
	return respuesta
}

//SignUp el usuario se crea una cuenta
func SignUp(unUsuario domain.Usuario) error {
	var respuesta error
	respuesta = EstaLogueado(unUsuario)
	if respuesta != nil {
		unUsuario.ID = len(UsuariosRegistrados)
		UsuariosRegistrados = append(UsuariosRegistrados, &unUsuario)
	}
	return nil
}

//PublishTweet publisher
func PublishTweet(unTweet *domain.Tweet) (int, error) {
	if unTweet.User.Username == "" {
		return 0, fmt.Errorf("User required")
	}
	if unTweet.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	unTweet.ID = len(myTweets)
	myTweets = append(myTweets, unTweet)
	return unTweet.ID, nil
}

//GetTweet getter
func GetTweets() []*domain.Tweet {
	return myTweets
}

//InitializeService init service
func InitializeService() {
	return
}

//GetTweetByID obtiene un tweet por ID
func GetTweetByID(id int) *domain.Tweet {
	for i, tweet := range myTweets {
		if tweet.ID == id {
			return myTweets[i]
		}
	}
	return nil
}
