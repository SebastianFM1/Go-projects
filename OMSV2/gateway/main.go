package main

import (
	"log"
	"net/http"

	pb "github.com/SebastianFM1/Go-projects/OMSV2/commons/api"
	common "github.com/SebastianFM1/Go-projects/OMSV2/commons"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
	ordersServiceAddr = "localhost:2000"
)



func main(){
	conn, err := grpc.Dial(ordersServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Failed to dial orders service: %v", err)
	}

	defer conn.Close()

	log.Printf("Dialing orders service at %s", ordersServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Println("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil{
		log.Fatal("Failed to start http server")
	}

}