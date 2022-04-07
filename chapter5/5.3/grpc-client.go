package main

import (
	"context"
	"fmt"
	pb "go-web/chapter5/5.3/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("连接失败:", err)
		return
	}
	defer conn.Close()
	client := pb.NewProgrammerServiceClient(conn)
	req := new(pb.Request)
	req.Name = "jaluik"
	response, err := client.GetProgrammerInfo(context.Background(), req)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("收到服务响应：%v \n", response)
}
