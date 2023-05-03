package spider_test

import (
	"context"
	"testing"

	"github.com/ryoshindo/spider"
	"github.com/stretchr/testify/assert"
)

func TestLoadDescribeVirtualService(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_service.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vs := spider.DescribeVirtualService{app}
		input, err := vs.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-service", *input.VirtualServiceName)
	}
}

func TestLoadCreateVirtualService(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_service.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vs := spider.CreateVirtualService{app}
		input, err := vs.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-service", *input.VirtualServiceName)
	}
}

func TestLoadUpdateVirtualService(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_service.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vs := spider.UpdateVirtualService{app}
		input, err := vs.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-service", *input.VirtualServiceName)
	}
}

func TestLoadDeleteVirtualService(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_service.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vs := spider.DeleteVirtualService{app}
		input, err := vs.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-service", *input.VirtualServiceName)
	}
}
