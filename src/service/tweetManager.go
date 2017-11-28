package service

import (
	"fmt"

	"github.com/cursoGO/src/domain"
)

//TweetManager el master de nuestra structura
type TweetManager struct {
	tweetsPorUsuario    map[int][]*domain.Tweet
	cantTweets          int
	usuariosRegistrados []*domain.Usuario
}

//GetTweetsPorUsuario getter mapa
func (tm *TweetManager) GetTweetsPorUsuario() map[int][]*domain.Tweet {
	return tm.tweetsPorUsuario
}

//GetUsuariosRegistrados getter usuarios registrados
func (tm *TweetManager) GetUsuariosRegistrados() []*domain.Usuario {
	return tm.usuariosRegistrados
}

//NewTweetManager constructor
func NewTweetManager() *TweetManager {
	tm := new(TweetManager)
	tm.tweetsPorUsuario = make(map[int][]*domain.Tweet)
	return tm
}

//RemoverTweet remueve un tweet
func (tm *TweetManager) RemoverTweet(idTweet int) {
	for key, tweetsDeUnUsuario := range tm.tweetsPorUsuario {
		//fmt.Printf("key[%s] value[%s]\n", k, v)
		for index, tweet := range tweetsDeUnUsuario {
			if tweet.GetID() == idTweet {
				tm.removerElemento(index, key)
			}
		}
	}
}

func (tm *TweetManager) removerElemento(index int, key int) {
	tm.tweetsPorUsuario[key] = append(tm.tweetsPorUsuario[key][:index], tm.tweetsPorUsuario[key][index+1:]...)
}

//EstaLogueado consulto si el usuario está logueado
func (tm *TweetManager) EstaLogueado(unUsuario domain.Usuario) bool {
	respuesta := false
	for _, user := range tm.usuariosRegistrados {
		if *user == unUsuario {
			respuesta = true
		}
	}
	return respuesta
}

//SignUp el usuario se crea una cuenta
func (tm *TweetManager) SignUp(unUsuario domain.Usuario) {
	var respuesta bool
	respuesta = tm.EstaLogueado(unUsuario)

	if !respuesta {
		unUsuario.SetID(len(tm.usuariosRegistrados))
		tm.usuariosRegistrados = append(tm.usuariosRegistrados, &unUsuario)
	}
}

//PublishTweet publisher
func (tm *TweetManager) PublishTweet(unTweet *domain.Tweet) (int, error) {
	if unTweet.GetUser().GetUsername() == "" {
		return 0, fmt.Errorf("User required")
	}
	if unTweet.GetText() == "" {
		return 0, fmt.Errorf("text is required")
	}

	unTweet.SetID(tm.cantTweets)
	tm.cantTweets++

	idUsuario := unTweet.GetUser().GetID()
	tm.tweetsPorUsuario[idUsuario] = append(tm.tweetsPorUsuario[idUsuario], unTweet)
	return unTweet.GetID(), nil
}

//GetTweets getter
func (tm *TweetManager) GetTweets() []*domain.Tweet {
	var allTweets []*domain.Tweet
	for _, tweetsDeUnUsuario := range tm.tweetsPorUsuario {
		for _, tweet := range tweetsDeUnUsuario {
			allTweets = append(allTweets, tweet)
		}
	}
	return allTweets
}

//InitializeService init service
func InitializeService() {
	return
}

//GetTweetByID obtiene un tweet por ID
func (tm *TweetManager) GetTweetByID(id int) *domain.Tweet {

	for _, tweetsDeUnUsuario := range tm.tweetsPorUsuario {
		//fmt.Printf("key[%s] value[%s]\n", k, v)
		for _, tweet := range tweetsDeUnUsuario {
			if tweet.GetID() == id {
				return tweet
			}
		}
	}
	return nil
}
