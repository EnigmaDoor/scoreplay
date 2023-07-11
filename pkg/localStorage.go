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

func WriteOutput(opts *Options, data *SrData) (err error) {
	file, err := json.MarshalIndent(data, "", " "); if err != nil {
		log.Println("[WriteOutput] Failure on marshalling", err)
		return
	}
	outputPath := path.Join(".", path.Clean(opts.LocalFolder), path.Clean(opts.Output))
	fmt.Println("[WriteOutput] Saving search dataset into ", outputPath)
	err = ioutil.WriteFile(outputPath, file, 0644); if err != nil {
		log.Println("[WriteOutput] Failure on write", outputPath, err)
		return
	}
	return
}

func ReadInput(opts *Options, data *SrData) (err error) {
	inputPath := path.Join(".", path.Clean(opts.LocalFolder), path.Clean(opts.Input))
	file, err := ioutil.ReadFile(inputPath); if err != nil {
		log.Println("[ReadInput] Failure on read", inputPath, err)
		return
	}
	err = json.Unmarshal([]byte(file), data); if err != nil {
		log.Println("[ReadInput] Failure on unmarshalling", err)
		return
	}
	fmt.Println(fmt.Sprintf("[ReadInput] Reusing %s search dataset", inputPath))
	return
}
