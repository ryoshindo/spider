package spider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
)

func (s *App) DescribeVirtualNode(ctx context.Context, path string) (*appmesh.DescribeVirtualNodeOutput, error) {
	vn := &describeVirtualNode{s}
	input, err := vn.Load(path)
	if err != nil {
		return &appmesh.DescribeVirtualNodeOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualNode(ctx, input)
	if err != nil {
		return &appmesh.DescribeVirtualNodeOutput{}, err
	}

	return output, nil
}

func (s *App) DescribeVirtualRouter(ctx context.Context, path string) (*appmesh.DescribeVirtualRouterOutput, error) {
	vr := &describeVirtualRouter{s}
	input, err := vr.Load(path)
	if err != nil {
		return &appmesh.DescribeVirtualRouterOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualRouter(ctx, input)
	if err != nil {
		return &appmesh.DescribeVirtualRouterOutput{}, err
	}

	return output, nil
}

func (s *App) DescribeRoute(ctx context.Context, path, virtualRouterPath string) (*appmesh.DescribeRouteOutput, error) {
	vrOutput, err := s.DescribeVirtualRouter(ctx, virtualRouterPath)
	if err != nil {
		return &appmesh.DescribeRouteOutput{}, err
	}

	r := &describeRoute{s}
	input, err := r.Load(path, *vrOutput.VirtualRouter.VirtualRouterName)
	if err != nil {
		return &appmesh.DescribeRouteOutput{}, err
	}

	output, err := s.appmesh.DescribeRoute(ctx, input)
	if err != nil {
		return &appmesh.DescribeRouteOutput{
			Route: &types.RouteData{
				VirtualRouterName: vrOutput.VirtualRouter.VirtualRouterName,
			},
		}, err
	}

	return output, nil
}

func (s *App) DescribeVirtualService(ctx context.Context, path string) (*appmesh.DescribeVirtualServiceOutput, error) {
	vs := &describeVirtualService{s}
	input, err := vs.Load(path)
	if err != nil {
		return &appmesh.DescribeVirtualServiceOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualService(ctx, input)
	if err != nil {
		return &appmesh.DescribeVirtualServiceOutput{}, err
	}

	return output, nil
}

func (s *App) DescribeVirtualGateway(ctx context.Context, path string) (*appmesh.DescribeVirtualGatewayOutput, error) {
	vg := &describeVirtualGateway{s}
	input, err := vg.Load(path)
	if err != nil {
		return &appmesh.DescribeVirtualGatewayOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualGateway(ctx, input)
	if err != nil {
		return &appmesh.DescribeVirtualGatewayOutput{}, err
	}

	return output, nil
}

func (s *App) DescribeGatewayRoute(ctx context.Context, path, virtualGatewayPath string) (*appmesh.DescribeGatewayRouteOutput, error) {
	vgOutput, err := s.DescribeVirtualGateway(ctx, virtualGatewayPath)
	if err != nil {
		return &appmesh.DescribeGatewayRouteOutput{}, err
	}

	gr := &describeGatewayRoute{s}
	input, err := gr.Load(path, *vgOutput.VirtualGateway.VirtualGatewayName)
	if err != nil {
		return &appmesh.DescribeGatewayRouteOutput{}, err
	}

	output, err := s.appmesh.DescribeGatewayRoute(ctx, input)
	if err != nil {
		return &appmesh.DescribeGatewayRouteOutput{
			GatewayRoute: &types.GatewayRouteData{
				VirtualGatewayName: vgOutput.VirtualGateway.VirtualGatewayName,
			},
		}, err
	}

	return output, nil
}
