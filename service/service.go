package service

import (
	pb "github.com/yuuyamad/grpc-sample/grpcsample"
	"path/filepath"
	"os"
	"io"
)
type MyFileService struct {

}
func(s *MyFileService) GetMyFile(_ *pb.RequestType, stream pb.File_GetMyFileServer) error {

	var filename string
	err := filepath.Walk("/Users/yamadayuuta/dev/src/github.com/yuuyamad/grpc-sample/hoge",func(path string, info os.FileInfo, err error) error {

			name, err := filepath.Rel("/Users/yamadayuuta/dev/src/github.com/yuuyamad/grpc-sample/hoge", path)
			if err != nil {
				return err
			}
			filename = filepath.ToSlash(name)
			f := &pb.MyFileResponse{
				Name: filename,
			}

			return stream.Send(f)
	})

	return err
}

func(s *MyFileService) Download(r *pb.DownloadRequestType, stream pb.File_DownloadServer) error {
	f, err := os.Open(filepath.Join("/Users/yamadayuuta/dev/src/github.com/yuuyamad/grpc-sample/hoge", r.Name))
	if err != nil {
		return err
	}
	defer f.Close()

	var b [4096*1000]byte
	for {
		n, err := f.Read(b[:])
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		err = stream.Send(&pb.DownloadFileResponse{
			Data: b[:n],
		})
		if err != nil {
			return err
		}
	}
	return err
}