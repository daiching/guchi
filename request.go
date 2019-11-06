package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	GET_GUCHI_ERR_HEADER = "GetGuchiError"
)

type Guchi struct {
	Name    string
	Message string
}

func getGuchi(f *Flags) error {
	res, err := http.Get(getGuchiJsonURL() + "?num=" + strconv.Itoa(f.Num))
	if err != nil {
		return getError(GET_GUCHI_ERR_HEADER, err.Error())
	}

	body, err := httpResponseBodyToBytes(res)
	if err != nil {
		return getError(GET_GUCHI_ERR_HEADER, err.Error())
	}

	var guchis []Guchi
	err = json.Unmarshal(body, &guchis)
	if err != nil {
		return getError(GET_GUCHI_ERR_HEADER, err.Error())
	}

	for i, g := range guchis {
		fmt.Println(g.Name + ":\n" + g.Message)
		if i+1 != len(guchis) {
			fmt.Print("\n")
		}
	}

	return nil
}

func postGuchi(f *Flags) error {
	guchi := Guchi{f.Name, f.Message}
	guchiBytes, err := json.Marshal(guchi)
	if err != nil {
		return getError(GET_GUCHI_ERR_HEADER, err.Error())
	}

	guchiBuf := bytes.NewBuffer(guchiBytes)
	_, err = http.Post(getGuchiJsonURL(), "application/json", guchiBuf)
	if err != nil {
		return getError(GET_GUCHI_ERR_HEADER, err.Error())
	}

	err = getGuchi(f)
	if err != nil {
		return getError(GET_GUCHI_ERR_HEADER, err.Error())
	}

	return nil
}

func getGuchiJsonURL() string {
	var url string
	if Conf.CustomDomain != "" {
		url = "://" + Conf.CustomDomain
	} else {
		url = "://" + DOMAIN_NAME
	}
	if Conf.CustomDomain != "" {
		url = url + ":" + Conf.CustomPort + "/" + GUCHI_JSON_PATH
	} else {
		url = url + "/" + GUCHI_JSON_PATH
	}
	if IS_HTTPS {
		return "https" + url
	} else {
		return "http" + url
	}
}
