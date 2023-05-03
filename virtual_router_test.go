package spider_test

import (
	"context"
	"testing"

	"github.com/ryoshindo/spider"
	"github.com/stretchr/testify/assert"
)

func TestLoadDescribeVirtualRouter(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_router.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vr := spider.DescribeVirtualRouter{app}
		input, err := vr.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-router", *input.VirtualRouterName)
	}
}

func TestLoadCreateVirtualRouter(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_router.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vr := spider.CreateVirtualRouter{app}
		input, err := vr.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-router", *input.VirtualRouterName)
	}
}

func TestLoadUpdateVirtualRouter(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_router.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vr := spider.UpdateVirtualRouter{app}
		input, err := vr.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-router", *input.VirtualRouterName)
	}
}

func TestLoadDeleteVirtualRouter(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_router.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vr := spider.DeleteVirtualRouter{app}
		input, err := vr.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-router", *input.VirtualRouterName)
	}
}
