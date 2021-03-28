package helper

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// Save 保存文件
func Save(path string, bytes []byte) error {
	_ = os.MkdirAll(path, 0666)
	fileName := fmt.Sprintf("%s/%d.jpg", path, randName())
	return ioutil.WriteFile(fileName, bytes, 0666)
}

func randName() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000)
}
