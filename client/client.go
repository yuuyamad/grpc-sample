package main

import (
	"google.golang.org/grpc"
	"log"
	pb "github.com/yuuyamad/grpc-sample/grpcsample"
	"context"
	"io"
	"os"
	"path/filepath"
	"flag"
)

func main(){

	var (

		host = flag.String("h", "127.0.0.1", "hostname")
	)

	conn, err := grpc.Dial(*host + ":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewFileClient(conn)
	ctx := context.Background()

	res, err := client.GetMyFile(ctx, &pb.RequestType{})
	if err != nil {
		panic(err)
	}

	for {
		feature, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.FileList(_) = _, %v", client, err)
		}
		log.Println(feature)

		if feature.Name == "."{
			continue
		}
		name := filepath.Join("./", filepath.FromSlash(feature.Name))

		if os.FileMode(feature.Mode).IsDir() {
			err := os.MkdirAll(name, os.FileMode(feature.Mode))

			if err != nil {
				log.Printf("%s: %v", feature.Name, err)
			}
			continue
		}
		sdown, err := client.Download(ctx, &pb.DownloadRequestType{Name: feature.Name})

		log.Println(name)
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