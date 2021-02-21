package db

import (
	"github.com/luobin998877/go_db/mongo"
	"github.com/luobin998877/go_db/mysql"
	"github.com/luobin998877/go_db/redis"
	"github.com/spf13/viper"
)

// ConnectRedis simple
func ConnectRedis() {
	addr := viper.GetString("Redis.addr")
	password := viper.GetString("Redis.password")
	db := viper.GetInt("Redis.db")
	redis.Connect(addr, password, db)
}

// ConnectRedisCluster cluster
func ConnectRedisCluster() {
	addrs := viper.GetStringSlice("RedisCluster.addrs")
	password := viper.GetString("RedisCluster.password")
	poolSize := viper.GetInt("RedisCluster.poolSize")
	redis.ConnectCluster(addrs, poolSize, password)
}

// ConnectMysql simple
func ConnectMysql() {
	addr := viper.GetString("Mysql.addrs")
	port := viper.GetInt("Mysql.port")
	user := viper.GetString("Mysql.user")
	password := viper.GetString("Mysql.password")
	dbname := viper.GetString("Mysql.dbname")
	mysql.Connect(addr, uint16(port), user, password, dbname)
}

// ConnectMongo simple
func ConnectMongo() {
	addr := viper.GetString("Mongo.addrs")
	port := viper.GetInt("Mongo.port")
	user := viper.GetString("Mongo.user")
	password := viper.GetString("Mongo.password")
	mongo.Connect(addr, port, user, password)
}
