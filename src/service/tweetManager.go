package service

import (
	"errors"
	"fmt"

	"github.com/cursoGO/src/domain"
)

//TweetManager el master de nuestra structura
type TweetManager struct {
	tweetsPorUsuario    map[int][]*domain.Tweet
	cantTweets          int
	usuariosRegistrados []*domain.Usuario
	usuariosLogueados   []*domain.Usuario
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

//EstaRegistrado consulto si el usuario está logueado
func (tm *TweetManager) EstaRegistrado(unUsuario *domain.Usuario) bool {
	isRegistered := false
	for _, user := range tm.usuariosRegistrados {
		if user == unUsuario {
			isRegistered = true
		}
	}
	return isRegistered
}

//SignUp el usuario se crea una cuenta
func (tm *TweetManager) SignUp(unUsuario *domain.Usuario) {
	if tm.EstaRegistrado(unUsuario) {
		return
	}
	unUsuario.SetID(len(tm.usuariosRegistrados))
	tm.usuariosRegistrados = append(tm.usuariosRegistrados, unUsuario)
}

//EstaLogueado devuelve un bool que dice si el usuario esta logueado en el sitio
func (tm *TweetManager) EstaLogueado(user *domain.Usuario) bool {
	isLogued := false
	if !tm.EstaRegistrado(user) {
		isLogued = false
	}
	for _, userTmp := range tm.usuariosLogueados {
		if user.GetID() == userTmp.GetID() {
			isLogued = true
		}
	}
	return isLogued
}

//LogIn en usuario que ya esta registrado en el sitio se loguea
func (tm *TweetManager) LogIn(user *domain.Usuario) {
	if tm.EstaLogueado(user) {
		return
	}
	tm.usuariosLogueados = append(tm.usuariosLogueados, user)
}

//LogOut cierro la seción del usuario
func (tm *TweetManager) LogOut(user *domain.Usuario) {
	if !tm.EstaLogueado(user) {
		return
	}
	for indice, usr := range tm.usuariosLogueados {
		if usr == user {
			tm.usuariosLogueados = append(tm.usuariosLogueados[:indice], tm.usuariosLogueados[indice+1:]...)
		}
	}
}

//RemoverTweet remueve un tweet
func (tm *TweetManager) RemoverTweet(idTweet int) {
	for key, tweetsDeUnUsuario := range tm.tweetsPorUsuario {
		//fmt.Printf("key[%s] value[%s]\n", k, v)
		for index, tweet := range tweetsDeUnUsuario {
			if tweet.GetID() == idTweet {
				tm.removerTweet(index, key)
			}
		}
	}
}

func (tm *TweetManager) removerTweet(index int, key int) {
	tm.tweetsPorUsuario[key] = append(tm.tweetsPorUsuario[key][:index], tm.tweetsPorUsuario[key][index+1:]...)
}

func (tm *TweetManager) alreadyExists(tw *domain.Tweet) bool {
	allTweets := tm.GetTweets()
	respuesta := false
	for _, tweet := range allTweets {
		if tweet.GetText() == tw.GetText() {
			respuesta = true
		}
	}
	return respuesta
}

//PublishTweet publisher
func (tm *TweetManager) PublishTweet(unTweet *domain.Tweet) (int, error) {

	if unTweet.GetUser().GetUsername() == "" {
		return 0, fmt.Errorf("User required")
	}
	if unTweet.GetText() == "" {
		return 0, fmt.Errorf("text is required")
	}
	if !tm.EstaLogueado(unTweet.GetUser()) {
		return 0, errors.New("El usuario no esta logueado")
	}
	if tm.alreadyExists(unTweet) {
		return 0, errors.New("El tweet ya existe")
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
func (tm *TweetManager) GetTweetByID(id int) (*domain.Tweet, error) {

	for _, tweetsDeUnUsuario := range tm.tweetsPorUsuario {
		for _, tweet := range tweetsDeUnUsuario {
			if tweet.GetID() == id {
				return tweet, nil
			}
		}
	}
	return nil, errors.New("El tweet no existe")
}

//EditarTweet modifico el texto de un tweet
func (tm *TweetManager) EditarTweet(id int, nuevoTexto string) error {
	tweet, err := tm.GetTweetByID(id)
	if err != nil {
		return errors.New("El tweet no exite")
	}
	tweet.SetText(nuevoTexto)
	return nil
}

//SeguirUsuario comienzo a seguir a un usuario
func (tm *TweetManager) SeguirUsuario(usuarioQueSigue *domain.Usuario, usuarioSeguido *domain.Usuario) error {
	if !tm.EstaRegistrado(usuarioQueSigue) {
		return errors.New("El primer usuario no existe")
	}
	if !tm.EstaRegistrado(usuarioSeguido) {
		return errors.New("El segundo usuario no existe")
	}
	usuarioQueSigue.AddSeguidos(usuarioSeguido)
	usuarioSeguido.AddSeguidor(usuarioQueSigue)
	return nil
}
