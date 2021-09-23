package app

type GobackConfig struct {
	Bucket string
}

var Config *GobackConfig = &GobackConfig{
	Bucket: "goback-archive",
}
