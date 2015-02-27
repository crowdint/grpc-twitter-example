package main

import (
	"github.com/gin-gonic/gin"

	pb "github.com/crowdint/grpc-twitter-example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var user = &pb.User{ID: 1}

func main() {
	router := gin.Default()
	router.GET("/tweets/:user_id", func(c *gin.Context) {
		// TODO(jpfuentes2): connection pooling or reuse
		conn, _ := grpc.Dial("127.0.0.1:10000")
		defer conn.Close()

		client := pb.NewTwitterClient(conn)
		timeline, _ := client.GetTimeline(context.Background(), user)

		c.JSON(200, timeline.Tweets)
	})
	router.Run(":8080")
}
