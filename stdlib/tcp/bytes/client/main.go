package main

import (
	"flag"
	"io/fs"
	"net"
	"os"
	"path/filepath"

	"testgo/stdlib/tcp/bytes/proto"

	log "github.com/sirupsen/logrus"
)

var file string

func init() {
	flag.StringVar(&file, "f", "/Users/user/Downloads", "file path, eg: -f config.yaml")
}

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", "192.168.213.183:9999")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("connect success")

	if err := filepath.Walk(file, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		name := path + "-1"

		f := &proto.File{
			Name: name,
			Data: data,
		}

		if err := f.WriteFile(conn); err != nil {
			return err
		}
		log.Infof("%s send", path)
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
