package spider_test

import (
	"context"
	"testing"

	"github.com/ryoshindo/spider"
	"github.com/stretchr/testify/assert"
)

func TestLoadDescribeMesh(t *testing.T) {
	ctx := context.Background()

	app, err := spider.New(ctx, &spider.Option{ConfigFilePath: "tests/spider.yml"})
	assert.Nil(t, err)

	m := spider.DescribeMesh{app}
	input, err := m.Load()
	assert.Nil(t, err)

	assert.Equal(t, "main-mesh", *input.MeshName)
	assert.Equal(t, "123456789012", *input.MeshOwner)

}
