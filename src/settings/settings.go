package settings

import (
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

type Settings struct {
	// runtime options
	GrpcUnaryInterceptor grpc.ServerOption
	// env config
	Port                int    `envconfig:"PORT" default:"8080"`
	GrpcPort            int    `envconfig:"GRPC_PORT" default:"8081"`
	DebugPort           int    `envconfig:"DEBUG_PORT" default:"6070"`
	UseStatsd           bool   `envconfig:"USE_STATSD" default:"true"`
	StatsdHost          string `envconfig:"STATSD_HOST" default:"localhost"`
	StatsdPort          int    `envconfig:"STATSD_PORT" default:"8125"`
	RuntimePath         string `envconfig:"RUNTIME_ROOT" default:"/srv/runtime_data/current"`
	RuntimeSubdirectory string `envconfig:"RUNTIME_SUBDIRECTORY"`
	LogLevel            string `envconfig:"LOG_LEVEL" default:"WARN"`
	RedisSocketType     string `envconfig:"REDIS_SOCKET_TYPE" default:"unix"`
	RedisUrl            string `envconfig:"REDIS_URL" default:"/var/run/nutcracker/ratelimit.sock"`
	RedisPoolSize       int    `envconfig:"REDIS_POOL_SIZE" default:"10"`
}

type Option func(*Settings)

func NewSettings() Settings {
	var s Settings

	err := envconfig.Process("", &s)
	if err != nil {
		panic(err)
	}

	return s
}

func GrpcUnaryInterceptor(i grpc.UnaryServerInterceptor) Option {
	return func(s *Settings) {
		s.GrpcUnaryInterceptor = grpc.UnaryInterceptor(i)
	}
}
