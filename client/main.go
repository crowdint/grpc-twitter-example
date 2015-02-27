package main

import (
	"fmt"
	"log"

	fake "github.com/Pallinder/go-randomdata"
	pb "github.com/crowdint/grpc-twitter-example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var user = &pb.User{ID: 1}

func NewTweet() *pb.Tweet {
	text := fmt.Sprintf("%s%s", fake.Paragraph(), fake.Paragraph())[0:140]
	return &pb.Tweet{
		ID:   uint64(fake.Number(0, 20)),
		Text: text,
		User: user,
	}
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:10000")
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterClient(conn)

	for i := 0; i < 5; i++ {
		_, err = client.Add(context.Background(), NewTweet())
		if err != nil {
			log.Fatalf("%#v.Add(_) = _, %v: ", client, err)
		}
	}

	timeline, err := client.GetTimeline(context.Background(), user)
	if err != nil {
		log.Fatalf("%#v.GetTimeline(_) = _, %v: ", client, err)
	}

	for _, t := range timeline.Tweets {
		fmt.Println(t)
	}

	fmt.Printf("\nGot %d tweets\n", len(timeline.Tweets))
}
