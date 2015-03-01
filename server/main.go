package main

import (
	"fmt"
	fake "github.com/Pallinder/go-randomdata"

	"log"
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/crowdint/grpc-twitter-example/proto"
)

func NewTweet() *pb.Tweet {
	text := fmt.Sprintf("%s%s", fake.Paragraph(), fake.Paragraph())[0:140]
	return &pb.Tweet{
		ID:   uint64(fake.Number(0, 20)),
		Text: text,
		User: &pb.User{
			ID: uint64(fake.Number(0, 20)),
		},
	}
}

type twitterServer struct {
	tweets map[uint64][]*pb.Tweet
	m      sync.RWMutex
}

func (t *twitterServer) Firehose(search *pb.Search, stream pb.Twitter_FirehoseServer) error {
	for {
		// fake grabbing tweets from upstream firehose
		tweet := NewTweet()
		if err := stream.Send(tweet); err != nil {
			return err
		}
	}

	return nil
}

func (t *twitterServer) GetTimeline(ctx context.Context, user *pb.User) (*pb.Timeline, error) {
	t.m.RLock()
	defer t.m.RUnlock()
	return &pb.Timeline{Tweets: t.tweets[user.ID]}, nil
}

func (t *twitterServer) Add(ctx context.Context, tweet *pb.Tweet) (*pb.Ack, error) {
	t.m.Lock()
	defer t.m.Unlock()

	userID := tweet.User.ID

	if t.tweets[userID] == nil {
		t.tweets[userID] = make([]*pb.Tweet, 0)
	}
	t.tweets[userID] = append(t.tweets[userID], tweet)

	return &pb.Ack{}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	twitter := &twitterServer{}
	twitter.tweets = make(map[uint64][]*pb.Tweet)

	pb.RegisterTwitterServer(server, twitter)
	server.Serve(lis)
}
