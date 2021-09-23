package app

import "time"

type archive struct {
	id   string `json:"id"`
	name string `json:"name"`
	md5  string `json:"md5"`
	size int32  `json:"size"`
}

//func (a *archive) ToString() string {
//
//}

type state struct {
	updated  time.Time `json:"updated"`
	archives []archive `json:"archives"`
}
