package libs

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//Translate struct data to Http data
func ToHttpData(i interface{}) (buffer *bytes.Buffer, err error) {
	jsonByte, err := json.Marshal(i)
	if err == nil {
		buffer = bytes.NewBuffer(jsonByte)
	}
	return buffer, err
}

//Check host  string format
func CheckHost(host string) bool {
	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
		return true
	}
	return false
}

//Get Prometheus alert rule string md5
func GetMd5(conf string) string {
	t := md5.New()
	io.WriteString(t, conf)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//Check user input params
func CheckParams(args []string) bool {
	if len(args) != 2 {
		fmt.Printf("Use %s  %s", args[0], "ConfigFile\n")
		fmt.Print("Example:\n prome-auto conf.ini\n")
		return false
	}
	confFile, _ := filepath.Abs(args[1])
	if pathExists(confFile) {
		return true
	}
	return false
}

//Check File Exist
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		fmt.Printf("Config file %s not found\n", path)
		return false
	}
	fmt.Printf("Config file %s not found\n", path)
	return false
}
