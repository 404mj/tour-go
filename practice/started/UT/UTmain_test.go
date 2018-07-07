package UT

import (
	"testing"
)

// 单测文件必须以_test.go结尾！单测方法必须以Test开头
// http://blog.csdn.net/code_segment/article/details/77507491
// http://www.flysnow.org/2017/05/16/go-in-action-go-unit-test.html

func TestAdd(t *testing.T) {
	sum := Add(1, 3)
	if sum == 4 {
		t.log("OK")
	} else {
		t.Fatal("wrong")
	}
}
