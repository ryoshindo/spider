package spider

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/goccy/go-yaml"
	goConfig "github.com/kayac/go-config"
)

var (
	appmeshName     = "main"           // FIXME
	virtualNodeName = "main-test"      // FIXME
	region          = "ap-northeast-1" // FIXME
	assumeRoleArn   = ""               // FIXME
)

type Config struct {
	Region string `yaml:"region"`
	Mesh   struct {
		Name  string `yaml:"name"`
		Owner string `yaml:"owner"`
	} `yaml:"mesh"`
	VirtualNodes   []string `yaml:"virtual_nodes"`
	VirtualRouters []struct {
		Definition string   `yaml:"definition"`
		Routes     []string `yaml:"routes"`
	} `yaml:"virtual_routers"`
	VirtualServices []string `yaml:"virtual_services"`
	VirtualGateways []struct {
		Definition    string   `yaml:"definition"`
		GatewayRoutes []string `yaml:"gateway_routes"`
	} `yaml:"virtual_gateways"`

	awsConfig aws.Config
}

type configLoader struct {
	*goConfig.Loader
}

func newConfigLoader() *configLoader {
	return &configLoader{}
}

func (l *configLoader) Load(ctx context.Context, path string) (*Config, error) {
	conf := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(content, conf); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

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

	if c.Mesh.Owner == "" {
		stsClient := sts.NewFromConfig(c.awsConfig)
		identity, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
		if err != nil {
			return fmt.Errorf("failed to get caller identity: %w", err)
		}
		c.Mesh.Owner = aws.ToString(identity.Account)
	}

	var err error
	c.awsConfig, err = awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load aws config: %w", err)
	}

	return nil
}
