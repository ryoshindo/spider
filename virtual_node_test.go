package spider_test

import (
	"context"
	"testing"

	"github.com/ryoshindo/spider"
	"github.com/stretchr/testify/assert"
)

func TestLoadDescribeVirtualNode(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_node.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vn := spider.DescribeVirtualNode{app}
		input, err := vn.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-node", *input.VirtualNodeName)
	}
}

func TestLoadCreateVirtualNode(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_node.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vn := spider.CreateVirtualNode{app}
		input, err := vn.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-node", *input.VirtualNodeName)
	}
}

func TestLoadUpdateVirtualNode(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_node.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vn := spider.UpdateVirtualNode{app}
		input, err := vn.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-node", *input.VirtualNodeName)
	}
}

func TestLoadDeleteVirtualNode(t *testing.T) {
	ctx := context.Background()

	for _, path := range []string{
		"tests/virtual_node.json",
	} {
		app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
		assert.Nil(t, err)

		vn := spider.DeleteVirtualNode{app}
		input, err := vn.Load(path)
		assert.Nil(t, err)

		assert.Equal(t, "main-mesh", *input.MeshName)
		assert.Equal(t, "123456789012", *input.MeshOwner)
		assert.Equal(t, "main-virtual-node", *input.VirtualNodeName)
	}
}
