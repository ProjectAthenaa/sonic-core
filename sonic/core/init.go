package core

import (
	"github.com/ProjectAthenaa/sonic-core/sonic/frame"
	"github.com/ProjectAthenaa/sonic-core/sonic/tools"
	"github.com/prometheus/common/log"
	"os"
)

var (
	coreInit int
	Base = &frame.CoreContext{}
)



func init() {
	coreInit = 1
	defer func() {
		coreInit = 0
	}()
	log.Infoln("start connect core databases")

	//pgql
	pgURL := os.Getenv("PG_URL")
	pgURL = tools.DefaultStr(pgURL, "postgresql://doadmin:olwikzizq24jj9ni@test-pg-do-user-9223163-0.b.db.ondigitalocean.com:25060/defaultdb")
	_, err := Base.NewPg("pg", pgURL)
	if err != nil {
		log.Fatalln("pg.connect", pgURL, err)
	}

	//redis
	rdbURL := os.Getenv("REDIS_URL")
	rdbURL = tools.DefaultStr(rdbURL, "rediss://default:n6luoc78ac44pgs0@test-redis-do-user-9223163-0.b.db.ondigitalocean.com:25061")
	_, err = Base.NewRedis("cache", rdbURL)
	if err != nil {
		log.Fatalln("redis.connect", rdbURL, err)
	}
	startRuntimeStats()
}
