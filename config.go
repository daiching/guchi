package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

var Conf struct {
	Name         string
	CustomDomain string
	CustomPort   string
}

func readConf() {
	aPath := getAbsolutePath(CONFIG_PATH)
	if !exist(aPath) {
		return
	}

	confBytes, err := ioutil.ReadFile(aPath)
	if err != nil {
		return
	}

	re := regexp.MustCompile(`\r\n|\r|\n`)
	confString := string(confBytes)
	confLines := re.Split(confString, -1)

	for _, l := range confLines {
		if tl := strings.Trim(l, " "); tl != "" && tl[1:] == "#" {
			continue
		}
		param := strings.Split(l, "=")
		if param[0] == "Name" {
			Conf.Name = param[1]
			continue
		}
		if param[0] == "CustomDomain" {
			Conf.CustomDomain = param[1]
			continue
		}
		if param[0] == "CustomPort" {
			Conf.CustomPort = param[1]
			continue
		}
	}

	return
}
