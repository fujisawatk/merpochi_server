package security

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

// Base64EncodeToString base64エンコードの文字列生成
func Base64EncodeToString(buf *bytes.Buffer) (string, error) {
	data, err := ioutil.ReadAll(buf)
	mine := http.DetectContentType(data)
	uri := "data:" + mine + ";base64," + base64.StdEncoding.EncodeToString(data)
	return uri, err
}
