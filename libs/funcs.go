package libs

import (
	"encoding/json"
	"bytes"
	"strings"
	"crypto/md5"
	"io"
	"fmt"
)

//传入待JSON化的struct数据，转换成HTTP POST Data
func ToHttpData(i interface{}) (buffer *bytes.Buffer, err error) {
	jsonByte, err := json.Marshal(i)
	if err == nil {
		buffer = bytes.NewBuffer(jsonByte)
	}
	return buffer, err
}

//检查用户提供的Prometheus服务器格式正确
func CheckHost(host string) bool {
	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
		return true
	}
	return false
}

//获取alertRule配置文件的MD5值
func GetMd5(conf string) string {
	t := md5.New()
	io.WriteString(t,conf)
	return fmt.Sprintf("%x",t.Sum(nil))
}