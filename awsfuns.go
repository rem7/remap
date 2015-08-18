package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	PUBLIC_IPV4     = "http://169.254.169.254/latest/meta-data/public-ipv4"
	PRIVATE_IPV4    = "http://169.254.169.254/latest/meta-data/local-ipv4"
	REGION_URL      = "http://169.254.169.254/latest/dynamic/instance-identity/document"
	INSTANCE_ID_URL = "http://169.254.169.254/latest/meta-data/instance-id"
	USER_DATA_URL   = "http://169.254.169.254/latest/user-data"
)

type InstanceConfig struct {
	Region string `json:"region"`
}

func getStringData(endpoint string) (string, error) {

	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func GetInstanceId() (string, error) {
	return getStringData(INSTANCE_ID_URL)
}

func GetUserData() (string, error) {
	return getStringData(USER_DATA_URL)
}

func GetRegion() (string, error) {

	resp, err := http.Get(REGION_URL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)

	v := InstanceConfig{}

	err = d.Decode(&v)
	if err != nil {
		return "", err
	}

	return v.Region, nil

}

func GetPrivateIP() (string, error) {
	return getIP(PRIVATE_IPV4)
}

func GetPublicIP() (string, error) {
	return getIP(PUBLIC_IPV4)
}

func getIP(URL string) (string, error) {

	timeout := time.Duration(REQ_TIMEOUT * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(URL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body), nil

}
