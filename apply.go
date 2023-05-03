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

func (s *App) Apply(ctx context.Context) error {
	if err := s.ApplyVirtualNode(ctx); err != nil {
		return err
	}

	if err := s.ApplyVirtualRouter(ctx); err != nil {
		return err
	}

	if err := s.ApplyRoute(ctx); err != nil {
		return err
	}

	if err := s.ApplyVirtualService(ctx); err != nil {
		return err
	}

	if err := s.ApplyVirtualGateway(ctx); err != nil {
		return err
	}

	if err := s.ApplyGatewayRoute(ctx); err != nil {
		return err
	}

	return nil
}

func (s *App) ApplyVirtualNode(ctx context.Context) error {
	for _, virtualNode := range s.config.VirtualNodes {
		output, _ := s.DescribeVirtualNode(ctx, virtualNode)
		if output.VirtualNode == nil {
			vn := &createVirtualNode{s}
			input, err := vn.Load(virtualNode)

			_, err = s.appmesh.CreateVirtualNode(ctx, input)
			if err != nil {
				return err
			}
		} else {
			vn := &updateVirtualNode{s}
			input, err := vn.Load(virtualNode)

			_, err = s.appmesh.UpdateVirtualNode(ctx, input)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *App) ApplyVirtualRouter(ctx context.Context) error {
	for _, virtualRouter := range s.config.VirtualRouters {
		output, _ := s.DescribeVirtualRouter(ctx, virtualRouter.Path)
		if output.VirtualRouter == nil {
			vr := &createVirtualRouter{s}
			input, err := vr.Load(virtualRouter.Path)

			_, err = s.appmesh.CreateVirtualRouter(ctx, input)
			if err != nil {
				return err
			}
		} else {
			vr := &updateVirtualRouter{s}
			input, err := vr.Load(virtualRouter.Path)

			_, err = s.appmesh.UpdateVirtualRouter(ctx, input)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *App) ApplyRoute(ctx context.Context) error {
	for _, virtualRouter := range s.config.VirtualRouters {
		for _, route := range virtualRouter.Routes {
			output, _ := s.DescribeRoute(ctx, route, virtualRouter.Path)
			if output.Route.Spec == nil {
				r := &createRoute{s}
				input, err := r.Load(route, *output.Route.VirtualRouterName)

				_, err = s.appmesh.CreateRoute(ctx, input)
				if err != nil {
					return err
				}
			} else {
				r := &updateRoute{s}
				input, err := r.Load(route, *output.Route.VirtualRouterName)

				_, err = s.appmesh.UpdateRoute(ctx, input)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (s *App) ApplyVirtualService(ctx context.Context) error {
	for _, virtualService := range s.config.VirtualServices {
		output, _ := s.DescribeVirtualService(ctx, virtualService)
		if output.VirtualService == nil {
			vs := &createVirtualService{s}
			input, err := vs.Load(virtualService)

			_, err = s.appmesh.CreateVirtualService(ctx, input)
			if err != nil {
				return err
			}
		} else {
			vs := &updateVirtualService{s}
			input, err := vs.Load(virtualService)

			_, err = s.appmesh.UpdateVirtualService(ctx, input)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *App) ApplyVirtualGateway(ctx context.Context) error {
	for _, virtualGateway := range s.config.VirtualGateways {
		output, _ := s.DescribeVirtualGateway(ctx, virtualGateway.Path)
		if output.VirtualGateway == nil {
			vg := &createVirtualGateway{s}
			input, err := vg.Load(virtualGateway.Path)

			_, err = s.appmesh.CreateVirtualGateway(ctx, input)
			if err != nil {
				return err
			}
		} else {
			vg := &updateVirtualGateway{s}
			input, err := vg.Load(virtualGateway.Path)

			_, err = s.appmesh.UpdateVirtualGateway(ctx, input)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *App) ApplyGatewayRoute(ctx context.Context) error {
	for _, virtualGateway := range s.config.VirtualGateways {
		for _, gatewayRoute := range virtualGateway.GatewayRoutes {
			output, _ := s.DescribeGatewayRoute(ctx, gatewayRoute, virtualGateway.Path)
			if output.GatewayRoute.Spec == nil {
				gr := &createGatewayRoute{s}
				input, err := gr.Load(gatewayRoute, *output.GatewayRoute.VirtualGatewayName)

				_, err = s.appmesh.CreateGatewayRoute(ctx, input)
				if err != nil {
					return err
				}
			} else {
				gr := &updateGatewayRoute{s}
				input, err := gr.Load(gatewayRoute, *output.GatewayRoute.VirtualGatewayName)

				_, err = s.appmesh.UpdateGatewayRoute(ctx, input)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
