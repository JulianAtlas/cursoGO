package main

import (
	"github.com/cursoGO/src/domain"
	"github.com/cursoGO/src/rest"
	"github.com/cursoGO/src/service"
)

func main() {
	tm := service.NewTweetManager()
	gs := rest.NewGinServer(tm)
	var user domain.Usuario
	user.SetMail("mercadolibre.com")
	user.SetUsername("meli-team")

	text := "Tweet re loko"
	tweet := domain.NewTextTweet(user, text)
	tm.SignUp(&user)
	tm.LogIn(&user)
	tm.PublishTweet(tweet)
	gs.StartGinServer()

	// shell := ishell.New()
	// shell.SetPrompt("Tweeter >> ")
	// shell.Print("Type 'help' to know commands\n")

	// shell.AddCmd(&ishell.Cmd{
	// 	Name: "publishTweet",
	// 	Help: "Publishes a tweet",
	// 	Func: func(c *ishell.Context) {

	// 		defer c.ShowPrompt(true)

	// 		c.Print("Write your tweet: ")

	// 		tweet := c.ReadLine()

	// 		service.PublishTweet(tweet)

	// 		c.Print("Tweet sent\n")

	// 		return
	// 	},
	// })

	// shell.AddCmd(&ishell.Cmd{
	// 	Name: "showTweet",
	// 	Help: "Shows a tweet",
	// 	Func: func(c *ishell.Context) {

	// 		defer c.ShowPrompt(true)

	// 		tweet := service.GetTweet()

	// 		c.Println(tweet)

	// 		return
	// 	},
	// })

	// shell.Run()

}
