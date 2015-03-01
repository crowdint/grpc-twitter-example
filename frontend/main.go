package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"

	pb "github.com/crowdint/grpc-twitter-example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
)

var user = &pb.User{ID: 1}

type AddTweetRequest struct {
	UserID uint64 `json:"userId"`
	Text   string `json:"text"`
}

// Pool is a naive connection pool to service
type Pool struct {
	clients []pb.TwitterClient
	sync.Mutex
}

func newPool(host string, conns int) *Pool {
	pool := &Pool{}
	for i := 0; i < conns; i++ {
		conn, _ := grpc.Dial("127.0.0.1:10000")
		pool.clients = append(pool.clients, pb.NewTwitterClient(conn))
	}
	return pool
}

func (p *Pool) Pop() pb.TwitterClient {
	p.Lock()
	defer p.Unlock()
	client := p.clients[0]
	p.clients = p.clients[1:]
	return client
}

func (p *Pool) Put(client pb.TwitterClient) {
	p.Lock()
	defer p.Unlock()
	p.clients = append(p.clients, client)
}

func (p *Pool) PopAndLock(fn func(client pb.TwitterClient)) {
	client := p.Pop()
	defer p.Put(client)
	fn(client)
}

func main() {
	pool := newPool("127.0.0.1:10000", 2)

	router := gin.Default()
	router.GET("/timeline/:user_id", func(c *gin.Context) {
		pool.PopAndLock(func(client pb.TwitterClient) {
			timeline, _ := client.GetTimeline(context.Background(), user)
			c.JSON(200, timeline.Tweets)
		})
	})

	router.POST("/tweet/:user_id", func(c *gin.Context) {
		pool.PopAndLock(func(client pb.TwitterClient) {
			request := AddTweetRequest{}
			if !c.Bind(&request) {
				c.Fail(http.StatusBadRequest, errors.New("could not add tweet because reasons"))
			} else {
				tweet := &pb.Tweet{Text: request.Text, User: user}
				// Add to BIG DATA STORAGE

				// Cache it
				if _, err := client.Add(context.Background(), tweet); err != nil {
					c.Fail(500, err)
				} else {
					c.String(204, "")
				}
			}
		})
	})

	router.Run(":8080")
}
