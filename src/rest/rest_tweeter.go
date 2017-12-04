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
	router.POST("publishTweet", server.publishTweet)
	// router.POST("publishImageTweet", server.publishImageTweet)
	// router.POST("publishQuoteTweet", server.publishQuoteTweet)

	router.Run()
}

func (server *GinServer) listTweets(c *gin.Context) {
	c.JSON(http.StatusOK, server.tweetmanager.GetTweets())
}

func (server *GinServer) publishTweet(c *gin.Context) {

	var tweet domain.TextTweet
	c.Bind(&tweet)

	tweetToPublish := domain.NewTextTweet(tweetdata.User, tweetdata.Text)

	id, err := server.tweetManager.PublishTweet(tweetToPublish, quit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet "+err.Error())
	} else {
		c.JSON(http.StatusOK, struct{ Id int }{id})
	}
}
