package spider_test

import (
	"context"
	"testing"

	"github.com/ryoshindo/spider"
	"github.com/stretchr/testify/assert"
)

func TestLoadDescribeGatewayRoute(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/gateway_route.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		gr := spider.DescribeGatewayRoute{app}
		input, err := gr.Load(path, "main-virtual-gateway")
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-gateway-route", *input.GatewayRouteName)
		assert.Equal(t, "main-virtual-gateway", *input.VirtualGatewayName)
	}
}

func TestLoadCreateGatewayRoute(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/gateway_route.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		gr := spider.CreateGatewayRoute{app}
		input, err := gr.Load(path, "main-virtual-gateway")
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-gateway-route", *input.GatewayRouteName)
		assert.Equal(t, "main-virtual-gateway", *input.VirtualGatewayName)
	}
}

func TestLoadUpdateGatewayRoute(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/gateway_route.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		gr := spider.UpdateGatewayRoute{app}
		input, err := gr.Load(path, "main-virtual-gateway")
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-gateway-route", *input.GatewayRouteName)
		assert.Equal(t, "main-virtual-gateway", *input.VirtualGatewayName)
	}
}

func TestLoadDeleteGatewayRoute(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/gateway_route.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		gr := spider.DeleteGatewayRoute{app}
		input, err := gr.Load(path, "main-virtual-gateway")
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-gateway-route", *input.GatewayRouteName)
		assert.Equal(t, "main-virtual-gateway", *input.VirtualGatewayName)
	}
}
