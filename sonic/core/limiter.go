package core

import (
	"context"
	"fmt"
	log "github.com/ProjectAthenaa/sonic-core/logs"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	license2 "github.com/ProjectAthenaa/sonic-core/sonic/database/ent/license"
	"math"
	"os"
	"strconv"
)

var defaultLimit = 1500

func init(){
	maxDefaultLimitString := os.Getenv("TASK_LIMIT_COUNT")
	maxDefaultLimit, err := strconv.Atoi(maxDefaultLimitString)
	if err == nil{
		defaultLimit = maxDefaultLimit
	}
}

func taskIsEligible(ctx context.Context, task *ent.Task) bool {
	key := fmt.Sprintf("tasks:users:%s", task.ID.String())
	user, err := task.QueryProfileGroup().QueryApp().QueryUser().WithLicense().First(ctx)
	if err != nil{
		log.Errorf("[server] [error retrieving user] [%s] [%s]", task.ID.String(), fmt.Sprint(err))
	}
	if v := rdb.Get(ctx, key).Val(); v != ""{
		count, err := strconv.Atoi(v)
		if err != nil{
			log.Errorf("[server] [error determining eligibility passing through] [%s]", task.ID.String())
			return false
		}

		if count >= getUserLimit(ctx, user){
			return false
		}

		goto incrementByOne

	}


	incrementByOne: rdb.Incr(ctx, key)

	return true
}

func getUserLimit(ctx context.Context, user *ent.User) int {
	license, err := user.License(ctx)
	if err != nil{
		return 0
	}

	switch license.Type {
	case license2.TypeRenewal, license2.TypeLifetime, license2.TypeBeta:
		return defaultLimit
	default:
		return math.MaxInt

	}

}
