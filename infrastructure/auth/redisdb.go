package auth

import (
	"github.com/go-redis/redis/v7"
	"github.com/redis/rueidis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type RedisService struct {
	Client  *redis.Client
	RClient rueidis.Client
}

func NewRedisDB(host, port, password string) (*RedisService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       viper.GetInt("REDIS_DB_NAMESPACE"),
	})

	rueidisClient, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{host + ":" + port},
		Password:    password,
		SelectDB:    viper.GetInt("REDIS_DB_NAMESPACE"),
	})
	if err != nil {
		logrus.Error(err.Error())
	}

	return &RedisService{
		Client:  redisClient,
		RClient: rueidisClient,
	}, nil
}
