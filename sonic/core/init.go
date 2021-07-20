package core

import (
	"github.com/ProjectAthenaa/sonic-core/sonic/frame"
	"os"
)

var Base = &frame.CoreContext{}

func init() {
	_, err := Base.NewPg("pg", os.Getenv("PG_URL"))
	if err != nil{
		panic(err)
	}
	_, err = Base.NewRedis("default", os.Getenv("REDIS_URL"))
	if err != nil{
		panic(err)
	}
}
