package main

import (
	"google.golang.org/grpc"
	"log"
	pb "github.com/yuuyamad/grpc-sample/grpcsample"
	"context"
	"io"
	"os"
	"path/filepath"
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

		name := filepath.Join("/Users/yamadayuuta/dev/src/github.com/yuuyamad/grpc-sample", filepath.FromSlash(feature.Name))
		sdown, err := client.Download(ctx, &pb.DownloadRequestType{Name: feature.Name})

		f, err := os.Create(name)
		if err != nil {
			log.Printf("%s: %v", feature.Name, err)
			sdown.CloseSend()
			continue
		}
		for{
			res, err := sdown.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Printf("%s: %v", feature.Name, err)
				break
			}
			f.Write(res.Data)
			if err != nil {
				log.Printf("%s: %v", feature.Name, err)
				break
			}
		}
		f.Close()
		sdown.CloseSend()
	}
}