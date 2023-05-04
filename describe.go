package spider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"golang.org/x/exp/slog"
)

func (s *App) DescribeVirtualNode(ctx context.Context, path string) (*appmesh.DescribeVirtualNodeInput, *appmesh.DescribeVirtualNodeOutput, error) {
	vn := &DescribeVirtualNode{s}
	input, err := vn.Load(path)
	if err != nil {
		return input, &appmesh.DescribeVirtualNodeOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualNode(ctx, input)
	if err != nil {
		s.Log(slog.LevelWarn, fmt.Sprintf("VirtualNode: Virtual Node '%s' in '%s' is not found", *input.VirtualNodeName, path))
		return input, &appmesh.DescribeVirtualNodeOutput{}, err
	}

	return input, output, nil
}

func (s *App) DescribeVirtualRouter(ctx context.Context, path string) (*appmesh.DescribeVirtualRouterInput, *appmesh.DescribeVirtualRouterOutput, error) {
	vr := &DescribeVirtualRouter{s}
	input, err := vr.Load(path)
	if err != nil {
		return input, &appmesh.DescribeVirtualRouterOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualRouter(ctx, input)
	if err != nil {
		s.Log(slog.LevelWarn, fmt.Sprintf("VirtualRouter: Virtual Router '%s' in '%s' is not found", *input.VirtualRouterName, path))
		return input, &appmesh.DescribeVirtualRouterOutput{}, err
	}

	return input, output, nil
}

func (s *App) DescribeRoute(ctx context.Context, path, virtualRouterPath string) (*appmesh.DescribeRouteInput, *appmesh.DescribeRouteOutput, error) {
	vrInput, vrOutput, err := s.DescribeVirtualRouter(ctx, virtualRouterPath)
	if err != nil {
		s.Log(slog.LevelWarn, fmt.Sprintf("Route: Virtual Router '%s' in '%s' is not found", *vrInput.VirtualRouterName, virtualRouterPath))
		return &appmesh.DescribeRouteInput{
			RouteName:         aws.String(""),
			VirtualRouterName: vrInput.VirtualRouterName,
		}, &appmesh.DescribeRouteOutput{}, err
	}

	r := &DescribeRoute{s}
	input, err := r.Load(path, *vrOutput.VirtualRouter.VirtualRouterName)
	if err != nil {
		return input, &appmesh.DescribeRouteOutput{}, err
	}

	output, err := s.appmesh.DescribeRoute(ctx, input)
	if err != nil {
		s.Log(slog.LevelWarn, fmt.Sprintf("Route: Route '%s' in '%s' is not found", *input.RouteName, path))
		return input, &appmesh.DescribeRouteOutput{
			Route: &types.RouteData{
				VirtualRouterName: vrOutput.VirtualRouter.VirtualRouterName,
			},
		}, err
	}

	return input, output, nil
}

func (s *App) DescribeVirtualService(ctx context.Context, path string) (*appmesh.DescribeVirtualServiceInput, *appmesh.DescribeVirtualServiceOutput, error) {
	vs := &DescribeVirtualService{s}
	input, err := vs.Load(path)
	if err != nil {
		return input, &appmesh.DescribeVirtualServiceOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualService(ctx, input)
	if err != nil {
		s.Log(slog.LevelWarn, fmt.Sprintf("VirtualService: Virtual Service '%s' in '%s' is not found", *input.VirtualServiceName, path))
		return input, &appmesh.DescribeVirtualServiceOutput{}, err
	}

	return input, output, nil
}

func (s *App) DescribeVirtualGateway(ctx context.Context, path string) (*appmesh.DescribeVirtualGatewayInput, *appmesh.DescribeVirtualGatewayOutput, error) {
	vg := &DescribeVirtualGateway{s}
	input, err := vg.Load(path)
	if err != nil {
		return input, &appmesh.DescribeVirtualGatewayOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualGateway(ctx, input)
	if err != nil {
		s.Log(slog.LevelWarn, fmt.Sprintf("VirtualGateway: Virtual Gateway '%s' in '%s' is not found", *input.VirtualGatewayName, path))
		return input, &appmesh.DescribeVirtualGatewayOutput{}, err
	}

	return input, output, nil
}

func (s *App) DescribeGatewayRoute(ctx context.Context, path, virtualGatewayPath string) (*appmesh.DescribeGatewayRouteInput, *appmesh.DescribeGatewayRouteOutput, error) {
	vgInput, vgOutput, err := s.DescribeVirtualGateway(ctx, virtualGatewayPath)
	if err != nil {
		s.Log(slog.LevelWarn, fmt.Sprintf("GatewayRoute: Virtual Gateway '%s' in '%s' is not found", *vgInput.VirtualGatewayName, virtualGatewayPath))
		return &appmesh.DescribeGatewayRouteInput{
			GatewayRouteName:   aws.String(""),
			VirtualGatewayName: vgInput.VirtualGatewayName,
		}, &appmesh.DescribeGatewayRouteOutput{}, err
	}

	gr := &DescribeGatewayRoute{s}
	input, err := gr.Load(path, *vgOutput.VirtualGateway.VirtualGatewayName)
	if err != nil {
		return input, &appmesh.DescribeGatewayRouteOutput{}, err
	}

	output, err := s.appmesh.DescribeGatewayRoute(ctx, input)
	if err != nil {
		s.Log(slog.LevelWarn, fmt.Sprintf("GatewayRoute: Gateway Route '%s' in '%s' is not found", *input.GatewayRouteName, path))
		return input, &appmesh.DescribeGatewayRouteOutput{
			GatewayRoute: &types.GatewayRouteData{
				VirtualGatewayName: vgOutput.VirtualGateway.VirtualGatewayName,
			},
		}, err
	}

	return input, output, nil
}
