package app

import "time"

type Archive struct {
	Id     string `json:"id"`
	Path   string `json:"path"`
	Bucket string `json:"bucket"`
	Size   int64  `json:"size"`
}

type State struct {
	Updated  time.Time `json:"updated"`
	Archives []Archive `json:"archives"`
}
