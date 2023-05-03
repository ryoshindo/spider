package spider_test

import (
	"context"
	"testing"

	"github.com/ryoshindo/spider"
	"github.com/stretchr/testify/assert"
)

func TestLoadDescribeRoute(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/route.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		r := spider.DescribeRoute{app}
		input, err := r.Load(path, "main-virtual-router")
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-route", *input.RouteName)
		assert.Equal(t, "main-virtual-router", *input.VirtualRouterName)
	}
}

func TestLoadCreateRoute(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/route.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		r := spider.CreateRoute{app}
		input, err := r.Load(path, "main-virtual-router")
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-route", *input.RouteName)
		assert.Equal(t, "main-virtual-router", *input.VirtualRouterName)
	}
}

func TestLoadUpdateRoute(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/route.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		r := spider.UpdateRoute{app}
		input, err := r.Load(path, "main-virtual-router")
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-route", *input.RouteName)
		assert.Equal(t, "main-virtual-router", *input.VirtualRouterName)
	}
}

func TestLoadDeleteRoute(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/route.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		r := spider.DeleteRoute{app}
		input, err := r.Load(path, "main-virtual-router")
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-route", *input.RouteName)
		assert.Equal(t, "main-virtual-router", *input.VirtualRouterName)
	}
}
