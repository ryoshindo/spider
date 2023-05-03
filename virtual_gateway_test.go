package spider_test

import (
	"context"
	"testing"

	"github.com/ryoshindo/spider"
	"github.com/stretchr/testify/assert"
)

func TestLoadDescribeVirtualGateway(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_gateway.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vg := spider.DescribeVirtualGateway{app}
		input, err := vg.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-gateway", *input.VirtualGatewayName)
	}
}

func TestLoadCreateVirtualGateway(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_gateway.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vg := spider.CreateVirtualGateway{app}
		input, err := vg.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-gateway", *input.VirtualGatewayName)
	}
}

func TestLoadUpdateVirtualGateway(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_gateway.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vg := spider.UpdateVirtualGateway{app}
		input, err := vg.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-gateway", *input.VirtualGatewayName)
	}
}

func TestLoadDeleteVirtualGateway(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_gateway.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vg := spider.DeleteVirtualGateway{app}
		input, err := vg.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-gateway", *input.VirtualGatewayName)
	}
}
