// +build !jsoniter

package json

import (
	"encoding/json"
	"fmt"
)

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	fmt.Println("Use [encoding/json] package")
	return json.MarshalIndent(v, prefix, indent)
}

// 使用条件编译的关键！！！以上的注释！
// go run -tags=jsoniter main.go
