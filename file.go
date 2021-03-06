package diff

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func ReadFile(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error while reading %s file: %v\n", path, err)
		// Do I need to forward error?
		return nil, err
	}
	return strings.Split(string(data), "\n"), err
}

func Print(file []string) {
	for _, s := range file {
		fmt.Println(s)
	}
}
