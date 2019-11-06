package main

import (
	"flag"
	"strings"
)

type Flags struct {
	Name    string
	Message string
	Num     int
}

func cmd() error {
	f := Flags{}
	name := flag.String("u", getUserName(), "Name of contributer. If there're no args, this option is ignored.")
	num := flag.Int("n", DEFAULT_COMMENT_NUM, "Number of Comment acquired. If there are args this option is ignored.")
	flag.Parse()
	args := flag.Args()

	f.Name = *name
	f.Num = *num

	if len(args) == 0 {
		return getGuchi(&f)
	}

	f.Message = strings.Join(args, " ")
	return postGuchi(&f)
}
