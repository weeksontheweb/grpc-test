package main

import (
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client :=proto.NewAddServiceClient(conn)

	//Use gin to create a set of api endpoints to call to the gRPC
	//functions
	g := gin.Default()

	g.GET("/add/:a/:b", func (ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx/JSON(http.Statusbadrequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx/JSON(http.Statusbadrequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		reg := &proto.Request{A: int64(a), B: int64(b)}

		if response, err := client.Add(ctx, req); err == nil{
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.result),
			}
		else {
			ctx.JSON(http.StatusinternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/mult/:a/:b", func (ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx/JSON(http.Statusbadrequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx/JSON(http.Statusbadrequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		reg := &proto.Request{A: int64(a), B: int64(b)}

		if response, err := client.Multiply(ctx, req); Err ==nil{
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.result),
			})
		else {
			ctx.JSON(http.StatusinternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8080"; err != nil (
		log.Fatalf("Failed to run server: %v", err)
	}
}
