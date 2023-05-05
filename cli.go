package spider

import (
	"context"
	"os"
	"strings"

	"github.com/alecthomas/kong"
)

type CliOptions struct {
	Config        string `help:"config file" default:"spider.yml"`
	AssumeRoleArn string `help:"the Role ARN to assume" default:""`

	Option *Option

	Apply   *ApplyOption   `cmd:"" help:"apply resources"`
	Destroy *DestroyOption `cmd:"" help:"destroy resources"`
}

func ParseCli(args []string) (string, *CliOptions, func(), error) {
	var opts CliOptions
	parser, err := kong.New(&opts)
	if err != nil {
		return "", nil, nil, err
	}

	c, err := parser.Parse(args)
	if err != nil {
		return "", nil, nil, err
	}

	sub := strings.Fields(c.Command())[0]

	opts.Option = &Option{
		ConfigFilePath: opts.Config,
		AssumeRoleArn:  opts.AssumeRoleArn,
	}

	return sub, &opts, func() { c.PrintUsage(true) }, nil
}

func dispatchCli(ctx context.Context, sub string, usage func(), opts *CliOptions) error {
	app, err := New(ctx, opts.Option)
	if err != nil {
		return err
	}

	switch sub {
	case "apply":
		return app.Apply(ctx, *opts.Apply)
	case "destroy":
		return app.Destroy(ctx, *opts.Destroy)
	default:
		usage()
	}

	return nil
}

type CliParseFunc func(args []string) (string, *CliOptions, func(), error)

func Cli(ctx context.Context, parse CliParseFunc) (int, error) {
	sub, opts, usage, err := parse(os.Args[1:])
	if err != nil {
		return 1, err
	}

	if err := dispatchCli(ctx, sub, usage, opts); err != nil {
		return 1, err
	}

	return 0, nil
}
