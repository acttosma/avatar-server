package startup

import (
	"avatarmeta.cc/avatar/model/entity"
	"avatarmeta.cc/avatar/pkg/util"
	"avatarmeta.cc/avatar/pkg/util/logger"
	"avatarmeta.cc/avatar/pkg/util/mysql"
	"avatarmeta.cc/avatar/pkg/util/redis"
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
