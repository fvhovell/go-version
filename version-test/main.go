// go_version.go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fvhovell/go-version/version"
)

func main() {
	log.Printf("Hello World:\n%v", version.MetaData)

	metaJson, err := json.Marshal(version.MetaData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(metaJson))
}
