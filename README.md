# grpc-twitter-example

## Usage

`go get github.com/crowdint/grpc-twitter-example`

Run the server:
```bash
$ go run server/main.go
```

Run the client:
```bash
$ go run client/main.go
```

You should see something like:
```bash
ID:7 text:"One dog rolled before him, well-nigh slashed in half; but a second had him by the thigh, a third gripped his collar be- hind, and a fourth h" user:<ID:1 >
ID:19 text:"Near it in the field, I remember, were three faint points of light, three telescopic stars infinitely remote, and all around it was the unfa" user:<ID:1 >
ID:1 text:"Near it in the field, I remember, were three faint points of light, three telescopic stars infinitely remote, and all around it was the unfa" user:<ID:1 >
ID:9 text:"The secular cooling that must someday overtake our planet has already gone far indeed with our neighbour.It was at this time that the meetin" user:<ID:1 >
ID:3 text:"The Nellie, a cruising yawl, swung to her anchor without a flutter of the sails, and was at rest.One dog rolled before him, well-nigh slashe" user:<ID:1 >

Got 5 tweets
```

## Web front-end example

This example demonstrates an HTTP front-end, using [gin](https://github.com/gin-gonic/gin), which connects to the backend Twitter server via grpc.

```bash
$ go run web/main.go
[GIN-debug] GET   /tweets/:user_id          --> main.func·001 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

Now, let's get some tweets!
```bash
$ curl http://localhost:8080/tweets/1
[{"ID":7,"text":"One dog rolled before him, well-nigh slashed in half; but a second had him by the thigh, a third gripped his collar be- hind, and a fourth h","user":{"ID":1}},{"ID":19,"text":"Near it in the field, I remember, were three faint points of light, three telescopic stars infinitely remote, and all around it was the unfa","user":{"ID":1}},{"ID":1,"text":"Near it in the field, I remember, were three faint points of light, three telescopic stars infinitely remote, and all around it was the unfa","user":{"ID":1}},{"ID":9,"text":"The secular cooling that must someday overtake our planet has already gone far indeed with our neighbour.It was at this time that the meetin","user":{"ID":1}},{"ID":3,"text":"The Nellie, a cruising yawl, swung to her anchor without a flutter of the sails, and was at rest.One dog rolled before him, well-nigh slashe","user":{"ID":1}}]
```
