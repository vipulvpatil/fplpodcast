package model

import "strings"

type Episode struct {
	AudioUrl string
	Title    string
}

func (e Episode) Filename() string {
	underscored := strings.Replace(e.Title, " ", "_", -1)
	sanitized := strings.Replace(underscored, "/", "_", -1)
	return sanitized
}
