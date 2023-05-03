package spider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

func (s *App) Describe(ctx context.Context) error {
	if _, err := s.DescribeVirtualNode(ctx); err != nil {
		return err
	}

	return nil
}

func (s *App) DescribeVirtualNode(ctx context.Context) (*appmesh.DescribeVirtualNodeOutput, error) {
	vn := &describeVirtualNode{s}
	input, err := vn.Load(s.config.VirtualNodes[0]) // FIXME: Allow for multiple file support
	if err != nil {
		return &appmesh.DescribeVirtualNodeOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualNode(ctx, input)
	if err != nil {
		return &appmesh.DescribeVirtualNodeOutput{}, err
	}

	return output, nil
}

func (s *App) Deploy(ctx context.Context) error {
	if err := s.DeployVirtualNode(ctx); err != nil {
		return err
	}

	return nil
}

func (s *App) DeployVirtualNode(ctx context.Context) error {
	output, _ := s.DescribeVirtualNode(ctx)
	if output.VirtualNode == nil {
		vn := &createVirtualNode{s}
		input, err := vn.Load(s.config.VirtualNodes[0]) // FIXME: Allow for multiple file support

		_, err = s.appmesh.CreateVirtualNode(ctx, input)
		if err != nil {
			return err
		}
	} else {
		vn := &updateVirtualNode{s}
		input, err := vn.Load(s.config.VirtualNodes[0]) // FIXME: Allow for multiple file support

		_, err = s.appmesh.UpdateVirtualNode(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}
