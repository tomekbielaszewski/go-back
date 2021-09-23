package app

type NoBucketAction string

const (
	CREATE NoBucketAction = "CREATE"
	EXIT                  = "EXIT"
	ASK                   = "ASK"
)

type GobackConfig struct {
	Bucket         string
	NoBucketAction NoBucketAction
}

var Config *GobackConfig = &GobackConfig{
	Bucket: "goback-archive",
}
