package console

import (
	"encoding/json"
	"fmt"
	"log"
)

// Pretty インデントを適用し標準出力
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
