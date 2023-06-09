package spider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

type App struct {
	appmesh *appmesh.Client

	config *Config
	loader *configLoader
	logger *log.Logger
}

func New(ctx context.Context, opt *Option) (*App, error) {
	loader := newConfigLoader()

	conf, err := loader.Load(ctx, opt.ConfigFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	conf.AssumeRole(opt.AssumeRoleArn)

	return &App{
		appmesh: appmesh.NewFromConfig(conf.awsConfig),

		config: conf,
		loader: loader,
		logger: newLogger(),
	}, nil
}

type Option struct {
	ConfigFilePath string
	AssumeRoleArn  string
}

func unmarshalJson(src []byte, v interface{}, path string) error {
	strict := json.NewDecoder(bytes.NewReader(src))
	strict.DisallowUnknownFields()

	if err := strict.Decode(v); err != nil {
		if !strings.Contains(err.Error(), "unknown field") {
			return err
		}
		lax := json.NewDecoder(bytes.NewReader(src))
		return lax.Decode(&v)
	}

	return nil
}
