package model

import "strings"

type Episode struct {
	AudioUrl string
	Title    string
}

func (e Episode) Filename() string {
	return strings.Replace(e.Title, " ", "_", -1)
}
