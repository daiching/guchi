package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func getError(header string, body string) error {
	return errors.New(header + ": " + body)
}

func getUserName() string {
	if Conf.Name != "" {
		return Conf.Name
	}
	return NO_NAME
}

func httpResponseBodyToBytes(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func getAbsolutePath(path string) string {
	wd, _ := os.Getwd()
	usr, _ := user.Current()
	tempPath, _ := filepath.Abs(path)
	aPath := strings.Replace(tempPath, fmt.Sprintf("%s/~", wd), usr.HomeDir, 1)
	return aPath
}

func exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
