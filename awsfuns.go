package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	SELF_CHECK      = "http://ident.me"
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

func getInstanceId() (string, error) {
	return getStringData(INSTANCE_ID_URL)
}

func getUserData() (string, error) {
	return getStringData(USER_DATA_URL)
}

func getRegion() (string, error) {

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

func getPublicIp() (string, error) {

	timeout := time.Duration(REQ_TIMEOUT * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(SELF_CHECK)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body), nil

}

func getLocalIP() (string, error) {

	ipStr := ""
	ifaces, err := net.Interfaces()
	if err != nil {
		return ipStr, err
	}

	for _, i := range ifaces {
		if i.Name == "eth0" {
			addrs, err := i.Addrs()
			if err != nil {
				return ipStr, err
			}
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}

				if ipv4 := ip.To4(); ipv4 != nil {
					ipStr = ipv4.String()
					return ipStr, nil
				}
			}
		}
	}

	return "", errors.New("No IP found.")

}
