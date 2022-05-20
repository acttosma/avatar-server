package startup

import (
	"acttos.com/avatar/model/entity"
	"acttos.com/avatar/pkg/util"
	"acttos.com/avatar/pkg/util/logger"
	"acttos.com/avatar/pkg/util/mysql"
	"acttos.com/avatar/pkg/util/redis"
)

// 初始化系统资源
func Initialize() {
	// 日志系统
	logger.InitLoggers()

	//工具化类
	util.InitCryptoHelper()
	util.InitTimeHelper()
	util.InitTextHelper()

	// 持久化层
	redis.InitRedisHelper()
	mysql.InitMySQLHelper()

	// 数据模型
	entity.InitTablesIfNeeded()
}
