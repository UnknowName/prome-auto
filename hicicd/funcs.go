package hicicd

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"errors"
	"strings"
	"encoding/json"
	"time"
	"github.com/unknowname/prome-auto/libs"
)

//Login get token
func Login(host, username, password string) (token string, err error) {
	url := host + "/user/login/"
	myAuth := Token{Username: username, Password: password}
	jsonByte, err := json.Marshal(myAuth)
	if err != nil {
		fmt.Println("Login Failed ", err)
		return token, err
	}
	myToken := Response{}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonByte))
	if err == nil {
		defer resp.Body.Close()
		byteResp, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(byteResp, &myToken)
		if err == nil {
			if myToken.Code == 200 {
				token = myToken.Data
			} else {
				err = errors.New(myToken.Message)
			}
		}
	} else {
		//Hidden login url
		errs := strings.Split(err.Error(), ":")
		err = errors.New(errs[len(errs)-1])
	}

	return token, err
}

//Create ConfigMap on OpenShift
func CreateConf(token, server, name, namespace, confStr string) error {
	if name == "" {
		name = "prometheus-rule"

	}
	if namespace == "" {
		namespace = "istio-system"
	}
	fileName := name + ".yaml"
	serverResp := Response{}
	data := make(map[string]string)
	data[fileName] = confStr
	url := server + "/configMap/add"
	confMap := ConfMap{
		Name:      name,
		NameSpace: namespace,
		Data:      data,
	}
	client := &http.Client{Timeout: 600 * time.Second}
	postData, err := libs.ToHttpData(confMap)
	req, err := http.NewRequest("POST", url, postData)
	if err == nil {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("Error occur ", err)
		return err
	}
	respByte, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respByte, &serverResp)
	if err == nil {
		if serverResp.Code == 200 && serverResp.Message == "Success" {
			err = nil
		} else {
			fmt.Print("Error occur at ", err)
			err = errors.New(serverResp.Message)
		}
	}
	defer resp.Body.Close()
	return err
}
