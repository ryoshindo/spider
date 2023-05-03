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

func (s *App) DescribeVirtualRouter(ctx context.Context) (*appmesh.DescribeVirtualRouterOutput, error) {
	vn := &describeVirtualRouter{s}
	input, err := vn.Load(s.config.VirtualRouters[0]) // FIXME: Allow for multiple file support
	if err != nil {
		return &appmesh.DescribeVirtualRouterOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualRouter(ctx, input)
	if err != nil {
		return &appmesh.DescribeVirtualRouterOutput{}, err
	}

	return output, nil
}

func (s *App) Apply(ctx context.Context) error {
	if err := s.ApplyVirtualNode(ctx); err != nil {
		return err
	}

	if err := s.ApplyVirtualRouter(ctx); err != nil {
		return err
	}

	return nil
}

func (s *App) ApplyVirtualNode(ctx context.Context) error {
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

func (s *App) ApplyVirtualRouter(ctx context.Context) error {
	output, _ := s.DescribeVirtualRouter(ctx)
	if output.VirtualRouter == nil {
		vn := &createVirtualRouter{s}
		input, err := vn.Load(s.config.VirtualRouters[0]) // FIXME: Allow for multiple file support

		_, err = s.appmesh.CreateVirtualRouter(ctx, input)
		if err != nil {
			return err
		}
	} else {
		vn := &updateVirtualRouter{s}
		input, err := vn.Load(s.config.VirtualRouters[0]) // FIXME: Allow for multiple file support

		_, err = s.appmesh.UpdateVirtualRouter(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}
