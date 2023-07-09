package scoreplay

import (
	"os"
	"fmt"
	"log"
	"path"
	"io/ioutil"
	"encoding/json"
)

func EnsureLocalFolder(opts *Options) (err error) {
	// MkdirAll return nil if folder already exists, otherwise creates it
	localPath := path.Join(".", path.Clean(opts.LocalFolder))
	err = os.MkdirAll(localPath, os.ModePerm)
	return
}

func WriteOutput(data *SrData, opts *Options) (err error) {
	file, err := json.MarshalIndent(data, "", " "); if err != nil {
		log.Println("[WriteOutput] Failure on marshalling", err)
		return
	}
	outputPath := path.Join(".", path.Clean(opts.LocalFolder), path.Clean(opts.Output))
	fmt.Println("[WriteOutput] Saving search dataset into ", outputPath)
	err = ioutil.WriteFile(outputPath, file, 0644); if err != nil {
		log.Println("[WriteOutput] Failure on write", err)
		return
	}
	return
}
