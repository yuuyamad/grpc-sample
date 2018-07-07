package service

import (
	pb "github.com/yuuyamad/grpc-sample/grpcsample"
	"path/filepath"
	"os"
)
type MyFileService struct {

}
func(s *MyFileService) GetMyFile(_ *pb.RequestType, stream pb.File_GetMyFileServer) error {

	var filename string
	err := filepath.Walk("/",func(path string, info os.FileInfo, err error) error {

			name, err := filepath.Rel("/", path)
			if err != nil {
				return err
			}
			filename = filepath.ToSlash(name)
			//file, _ := os.Open(path)
			//defer file.Close()
			//filename = getFileName(path)
			f := &pb.MyFileResponse{
				Name: filename,
			}

			return stream.Send(f)
	})

	return err
}
/*
func getFileName(path string) string{
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
*/