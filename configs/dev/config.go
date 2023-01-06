package configs

import (
	"github.com/gowebspider/renderpool/pkg/mongodbstorage"
	"github.com/gowebspider/renderpool/pkg/redisstorage"
)

type RenderBackend struct {
	RedisConfig redisstorage.Storage
	MongoConfig mongodbstorage.Storage
}

type RenderClient struct {
	RenderServerURI string
	RenderPoolSize  int
}

type RenderConfig struct {
	RenderClient  RenderClient
	RenderBackend RenderBackend
}
