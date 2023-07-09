package scoreplay

import (
	"fmt"
	"log"
	"path"
	"io/ioutil"
	"enconding/json"
)

func WriteOutput(data *SrData, opts *Options) (err error) {
	file, err := json.MarshalIndent(data, "", " "); if err != nil {
		log.Println("[WriteOutput] Failure on marshalling", err)
		return
	}
	outputPath := path.Join(path.Clean(opts.LocalFolder), path.clean(opts.Output))
	fmt.Println("[WriteOutput] Saving search dataset into ", outputPath)
	err = ioutil.WriteFile(outputPath, file, 0644); if err != {
		log.Println("[WriteOutput] Failure on write", err)
		return
	}
	return
}
