// +build jsoniter

package json

import (
	"fmt"
	"github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	fmt.Println("Use [jsoniter] package")
	return json.MarshalIndent(v, prefix, indent)
}
