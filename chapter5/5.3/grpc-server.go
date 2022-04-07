package main

import (
	"context"
	pb "go-web/chapter5/5.3/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ProgrammerServiceServer struct {
	pb.UnimplementedProgrammerServiceServer
}

func (p ProgrammerServiceServer) GetProgrammerInfo(_ context.Context, req *pb.Request) (resp *pb.Response, err error) {
	name := req.Name
	if name == "jaluik" {
		resp = &pb.Response{
			Uid:      6,
			Username: name,
			Job:      "FE",
			GoodAt:   []string{"Go", "JS"},
		}
	}
	err = nil
	return
}

func main() {
	addr := ":8080"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("监听失败了:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterProgrammerServiceServer(s, &ProgrammerServiceServer{})
	err = s.Serve(l)
	if err != nil {
		log.Fatal("Serve Error:", err)
		return
	}

}
