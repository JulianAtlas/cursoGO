package rest

import (
	"net/http"

	"github.com/cursoGO/src/domain"

	"github.com/cursoGO/src/service"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	tweetmanager *service.TweetManager
}

func NewGinServer(tweetManager *service.TweetManager) *GinServer {
	return &GinServer{tweetManager}
}

func (server *GinServer) StartGinServer() {
	router := gin.Default()

	router.GET("/listTweets", server.listTweets)
	// router.GET("/listTweets/:user", server.listTweets)
	router.POST("/publishTweet", server.publishTweet)
	// router.POST("publishImageTweet", server.publishImageTweet)
	// router.POST("publishQuoteTweet", server.publishQuoteTweet)

	router.Run()
}

func (server *GinServer) listTweets(c *gin.Context) {
	c.JSON(http.StatusOK, server.tweetmanager.GetTweets())
}

func (server *GinServer) publishTweet(c *gin.Context) {

	var tweetData domain.TextTweet
	c.Bind(&tweetData)

	tweet := domain.NewTextTweet(*tweetData.User, tweetData.Text)

	id, err := server.tweetmanager.PublishTweet(tweet)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	} else {
		c.JSON(http.StatusOK, struct{ ID int }{id})
	}
}
