package startup

import (
	"acttos.com/avatar/model/entity"
	"acttos.com/avatar/pkg/util"
	"acttos.com/avatar/pkg/util/logger"
	"acttos.com/avatar/pkg/util/mysql"
	"acttos.com/avatar/pkg/util/redis"
)

// initialize the resource of system
func Initialize() {
	// log system
	logger.InitLoggers()

	//utilities
	util.InitCryptoHelper()
	util.InitTimeHelper()
	util.InitTextHelper()

	// persistence layer
	redis.InitRedisHelper()
	mysql.InitMySQLHelper()

	// data entities
	entity.InitTablesIfNeeded()
}
