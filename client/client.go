package main

import (
	"google.golang.org/grpc"
	"log"
	pb "github.com/yuuyamad/grpc-sample/grpcsample"
	"context"
	"io"
)

func main(){
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewFileClient(conn)

	ctx := context.Background()
	res, err := client.GetMyFile(ctx, &pb.RequestType{})

	for {
		feature, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.FileList(_) = _, %v", client, err)
		}
		log.Println(feature)
	}
	//fmt.Printf("result:%#v \n", res)
	//fmt.Printf("error::%#v \n", err)
}