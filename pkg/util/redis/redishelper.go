package redis

import (
	"context"
	"sync"
	"time"

	"avatarmeta.cc/avatar/pkg/setting"
	"avatarmeta.cc/avatar/pkg/util/logger"
	"github.com/go-redis/redis/v8"
)

type RedisHelper struct {
	redisClient *redis.Client
}

var Helper *RedisHelper
var once sync.Once
var ctx = context.Background()

func InitRedisHelper() {
	once.Do(func() {
		Helper = initRedisHelper()
	})
}

// Setup initialize the configuration instance
func initRedisHelper() *RedisHelper {
	redisHelper := &RedisHelper{}
	gRedis := redis.NewClient(&redis.Options{
		Addr:         setting.GetInstance().RedisUrl,
		Password:     setting.GetInstance().RedisPassword,
		DB:           0,
		PoolSize:     setting.GetInstance().RedisPoolSize,
		MinIdleConns: setting.GetInstance().RedisPoolMinIdel,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			logger.Monitor.Debugf("Redis client has connected to the server:%+v", cn)
			return nil
		},
	})

	redisHelper.redisClient = gRedis

	return redisHelper
}

// SET key value
// Set key to hold the string value. If key already holds a value, it is overwritten, regardless of its type.
// Any previous time to live associated with the key is discarded on successful SET operation.
func (rh *RedisHelper) Set(key string, value string) bool {
	err := rh.redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when Set key:%s with value %s, err:%+v", key, value, err)
		return false
	}

	return true
}

// SET key value [EX seconds|PX milliseconds|EXAT unix-time-seconds|PXAT unix-time-milliseconds|KEEPTTL] [NX|XX] [GET]
// Set key to hold the string value. If key already holds a value, it is overwritten, regardless of its type.
// Any previous time to live associated with the key is discarded on successful SET operation.
func (rh *RedisHelper) SetWithExpir(key string, value string, expiration time.Duration) bool {
	err := rh.redisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when SetWithExpir key:%s with value %s and expir %d, err:%+v", key, value, expiration, err)
		return false
	}

	return true
}

// GET key
// Get the value of key. If the key does not exist the special value nil is returned.
// An error is returned if the value stored at key is not a string, because GET only handles string values.
func (rh *RedisHelper) Get(key string) string {
	val, err := rh.redisClient.Get(ctx, key).Result()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when Get key:%s 's value, err:%+v", key, err)
		return ""
	}

	return val
}

// HSET key field value [field value ...]
// Sets field in the hash stored at key to value. If key does not exist, a new key holding a hash is created.
// If field already exists in the hash, it is overwritten.
func (rh *RedisHelper) HSet(key string, field string, value string) bool {
	err := rh.redisClient.HSet(ctx, key, field, value).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when HSet key:%s with field %s and value %s, err:%+v", key, field, value, err)
		return false
	}

	return true
}

// HGET key field
// Returns the value associated with field in the hash stored at key.
func (rh *RedisHelper) HGet(key string, field string) string {
	val, err := rh.redisClient.HGet(ctx, key, field).Result()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when HGet key:%s with field %s 's value, err:%+v", key, field, err)
		return ""
	}

	return val
}

// LPOP key
// Removes and returns the first values of the list stored at key.
// The command pops a single value from the beginning of the list.
func (rh *RedisHelper) LPop(key string) string {
	val, err := rh.redisClient.LPop(ctx, key).Result()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when LPop key:%s 's value, err:%+v", key, err)
		return ""
	}
	return val
}

// LPOP key [count]
// Removes and returns the first values of the list stored at key.
// By default, the command pops a single value from the beginning of the list.
// When provided with the optional count argument, the reply will consist of up to count values, depending on the list's length.
func (rh *RedisHelper) LPopCount(key string, count int64) []string {
	val, err := rh.redisClient.LPopCount(ctx, key, int(count)).Result()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when LPopCount key:%s with count %d 's value, err:%+v", key, count, err)
		return []string{}
	}
	return val
}

// LPUSH key value [value ...]
// Insert all the specified values at the head of the list stored at key.
// If key does not exist, it is created as empty list before performing the push operations.
// When key holds a value that is not a list, an error is returned.
func (rh *RedisHelper) LPush(key string, value ...interface{}) bool {
	err := rh.redisClient.LPush(ctx, key, value...).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when LPush key:%s with value %s, err:%+v", key, value, err)
		return false
	}

	return true
}

// LPUSHX key value [value ...]
// Inserts specified values at the head of the list stored at key, only if key already exists and holds a list.
// In contrary to LPUSH, no operation will be performed when key does not yet exist.
func (rh *RedisHelper) LPushX(key string, value ...interface{}) bool {
	err := rh.redisClient.LPushX(ctx, key, value...).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when LPushX key:%s with value %s, err:%+v", key, value, err)
		return false
	}

	return true
}

// LREM key count value
// Removes the first count occurrences of values equal to value from the list stored at key.
// The count argument influences the operation in the following ways:
// 		count > 0: Remove values equal to value moving from head to tail.
// 		count < 0: Remove values equal to value moving from tail to head.
// 		count = 0: Remove all values equal to value.
// For example, LREM list -2 "hello" will remove the last two occurrences of "hello" in the list stored at list.
// Note that non-existing keys are treated like empty lists, so when key does not exist, the command will always return 0.
func (rh *RedisHelper) LRem(key string, count int64, value string) bool {
	err := rh.redisClient.LRem(ctx, key, count, value).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when LRem key:%s with count %d and value %s, err:%+v", key, count, value, err)
		return false
	}

	return true
}

// LSET key index value
// Sets the list value at index to value.
// An error is returned for out of range indexes.
func (rh *RedisHelper) LSet(key string, index int64, value string) bool {
	err := rh.redisClient.LSet(ctx, key, index, value).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when LSet key:%s with index %d and value %s, err:%+v", key, index, value, err)
		return false
	}

	return true
}

// RPUSH key value [value ...]
// Insert all the specified values at the tail of the list stored at key.
// If key does not exist, it is created as empty list before performing the push operation.
// When key holds a value that is not a list, an error is returned.
// It is possible to push multiple values using a single command call just specifying multiple arguments at the end of the command.
// values are inserted one after the other to the tail of the list, from the leftmost value to the rightmost value.
// So for instance the command RPUSH mylist a b c will result into a list containing a as first value, b as second value and c as third value.
func (rh *RedisHelper) RPush(key string, value ...interface{}) bool {
	err := rh.redisClient.RPush(ctx, key, value...).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when RPush key:%s with value %s, err:%+v", key, value, err)
		return false
	}

	return true
}

// RPOP key [count]
// Removes and returns the last values of the list stored at key.
// By default, the command pops a single value from the end of the list.
// When provided with the optional count argument, the reply will consist of up to count values, depending on the list's length.
func (rh *RedisHelper) RPop(key string) string {
	val, err := rh.redisClient.RPop(ctx, key).Result()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when RPop key:%s 's value, err:%+v", key, err)
		return ""
	}

	return val
}

// SADD key member [member ...]
// Add the specified members to the set stored at key.
// Specified members that are already a member of this set are ignored.
// If key does not exist, a new set is created before adding the specified members.
// An error is returned when the value stored at key is not a set.
func (rh *RedisHelper) SAdd(key string, value ...interface{}) bool {
	err := rh.redisClient.SAdd(ctx, key, value...).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when SAdd key:%s with value %s, err:%+v", key, value, err)
		return false
	}

	return true
}

// SPOP key
// Removes and returns one random members from the set value store at key.
// By default, the command pops a single member from the set.
// When provided with the optional count argument, the reply will consist of up to count members, depending on the set's cardinality.
func (rh *RedisHelper) SPop(key string) string {
	val, err := rh.redisClient.SPop(ctx, key).Result()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when SPop key:%s 's value, err:%+v", key, err)
		return ""
	}
	return val
}

// SPOP key [count]
// Removes and returns one or more random members from the set value store at key.
// By default, the command pops a single member from the set.
// When provided with the optional count argument, the reply will consist of up to count members, depending on the set's cardinality.
func (rh *RedisHelper) SPopN(key string, count int64) []string {
	val, err := rh.redisClient.SPopN(ctx, key, count).Result()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when SPop key:%s 's value, err:%+v", key, err)
		return nil
	}
	return val
}

// SREM key member [member ...]
// Remove the specified members from the set stored at key.
// Specified members that are not a member of this set are ignored.
// If key does not exist, it is treated as an empty set and this command returns 0.
// An error is returned when the value stored at key is not a set.
func (rh *RedisHelper) SRem(key string, value ...interface{}) bool {
	err := rh.redisClient.SRem(ctx, key, value...).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when SRem key:%s with value %s, err:%+v", key, value, err)
		return false
	}

	return true
}

// SMEMBERS key
// Returns all the members of the set value stored at key.
// This has the same effect as running SINTER with one argument key.
// Return Array reply: all elements of the set.
func (rh *RedisHelper) SMembers(key string) []string {
	val, err := rh.redisClient.SMembers(ctx, key).Result()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when SMembers key:%s, val:%+v, err:%+v", key, val, err)
		return nil
	}

	return val
}

// INCR key
// Increments the number stored at key by one.
// If the key does not exist, it is set to 0 before performing the operation.
// An error is returned if the key contains a value of the wrong type or contains a string that can not be represented as integer.
// This operation is limited to 64 bit signed integers.
func (rh *RedisHelper) Incr(key string) int64 {
	val, err := rh.redisClient.Incr(ctx, key).Result()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when Incr key:%s 's value, err:%+v", key, err)
		return 0
	}
	return val
}

// INCRBY key increment
// Increments the number stored at key by increment.
// If the key does not exist, it is set to 0 before performing the operation.
// An error is returned if the key contains a value of the wrong type or contains a string that can not be represented as integer.
// This operation is limited to 64 bit signed integers.
func (rh *RedisHelper) IncrBy(key string, value int64) int64 {
	val, err := rh.redisClient.IncrBy(ctx, key, value).Result()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when IncrBy key:%s and value %d, err:%+v", key, value, err)
		return 0
	}
	return val
}

// DEL key [key ...]
// Removes the specified keys. A key is ignored if it does not exist.
func (rh *RedisHelper) Del(key ...string) bool {
	err := rh.redisClient.Del(ctx, key...).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when Del key:%s, err:%+v", key, err)
		return false
	}

	return true
}

// HDEL key field [field ...]
// Removes the specified fields from the hash stored at key.
// Specified fields that do not exist within this hash are ignored.
// If key does not exist, it is treated as an empty hash and this command returns 0.
func (rh *RedisHelper) HDel(key string, field ...string) bool {
	err := rh.redisClient.HDel(ctx, key, field...).Err()
	if err != nil {
		logger.Monitor.Errorf("Error occurs when HDel key:%s with field %s 's value, err:%+v", key, field, err)
		return false
	}

	return true
}
