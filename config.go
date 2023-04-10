package spider

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

var (
	appmeshName     = "main"           // FIXME
	virtualNodeName = "main-test"      // FIXME
	region          = "ap-northeast-1" // FIXME
	assumeRoleArn   = ""               // FIXME
)

type Config struct {
	Region string

	awsConfig aws.Config
}

type configLoader struct {
}

func newConfigLoader() *configLoader {
	return &configLoader{}
}

func (l *configLoader) Load(ctx context.Context) (*Config, error) {
	conf := &Config{}

	if err := conf.Restrict(ctx); err != nil {
		return nil, err
	}

	return conf, nil
}

func (c *Config) AssumeRole(assumeRoleArn string) {
	if assumeRoleArn == "" {
		return
	}

	stsClient := sts.NewFromConfig(c.awsConfig)
	assumeRoleProvider := stscreds.NewAssumeRoleProvider(stsClient, assumeRoleArn)
	c.awsConfig.Credentials = aws.NewCredentialsCache(assumeRoleProvider)
}

func (c *Config) Restrict(ctx context.Context) error {
	if c.Region == "" {
		c.awsConfig.Region = os.Getenv("AWS_REGION")
	}

	var err error
	c.awsConfig, err = awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load aws config: %w", err)
	}

	return nil
}
