package spider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

type App struct {
	appmesh *appmesh.Client

	config *Config
}

func New(ctx context.Context) (*App, error) {
	loader := newConfigLoader()

	conf, err := loader.Load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	conf.AssumeRole(assumeRoleArn) // FIXME

	return &App{
		appmesh: appmesh.NewFromConfig(conf.awsConfig),

		config: conf,
	}, nil
}
