package spider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
)

func (s *App) Describe(ctx context.Context) error {
	if _, err := s.DescribeVirtualNode(ctx); err != nil {
		return err
	}

	if _, err := s.DescribeVirtualRouter(ctx); err != nil {
		return err
	}

	if _, err := s.DescribeRoute(ctx); err != nil {
		return err
	}

	if _, err := s.DescribeVirtualService(ctx); err != nil {
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
	vr := &describeVirtualRouter{s}
	input, err := vr.Load(s.config.VirtualRouters[0].Definition) // FIXME: Allow for multiple file support
	if err != nil {
		return &appmesh.DescribeVirtualRouterOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualRouter(ctx, input)
	if err != nil {
		return &appmesh.DescribeVirtualRouterOutput{}, err
	}

	return output, nil
}

func (s *App) DescribeRoute(ctx context.Context) (*appmesh.DescribeRouteOutput, error) {
	vrOutput, err := s.DescribeVirtualRouter(ctx)
	if err != nil {
		return &appmesh.DescribeRouteOutput{}, err
	}

	r := &describeRoute{s}
	input, err := r.Load(s.config.VirtualRouters[0].Routes[0], *vrOutput.VirtualRouter.VirtualRouterName) // FIXME: Allow for multiple file support
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

func (s *App) DescribeVirtualService(ctx context.Context) (*appmesh.DescribeVirtualServiceOutput, error) {
	vs := &describeVirtualService{s}
	input, err := vs.Load(s.config.VirtualServices[0]) // FIXME: Allow for multiple file support
	if err != nil {
		return &appmesh.DescribeVirtualServiceOutput{}, err
	}

	output, err := s.appmesh.DescribeVirtualService(ctx, input)
	if err != nil {
		return &appmesh.DescribeVirtualServiceOutput{}, err
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
		vr := &createVirtualRouter{s}
		input, err := vr.Load(s.config.VirtualRouters[0].Definition) // FIXME: Allow for multiple file support

		_, err = s.appmesh.CreateVirtualRouter(ctx, input)
		if err != nil {
			return err
		}
	} else {
		vr := &updateVirtualRouter{s}
		input, err := vr.Load(s.config.VirtualRouters[0].Definition) // FIXME: Allow for multiple file support

		_, err = s.appmesh.UpdateVirtualRouter(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *App) ApplyRoute(ctx context.Context) error {
	output, _ := s.DescribeRoute(ctx)
	if output.Route.Spec == nil {
		r := &createRoute{s}
		input, err := r.Load(s.config.VirtualRouters[0].Routes[0], *output.Route.VirtualRouterName) // FIXME: Allow for multiple file support

		_, err = s.appmesh.CreateRoute(ctx, input)
		if err != nil {
			return err
		}
	} else {
		r := &updateRoute{s}
		input, err := r.Load(s.config.VirtualRouters[0].Routes[0], *output.Route.VirtualRouterName) // FIXME: Allow for multiple file support

		_, err = s.appmesh.UpdateRoute(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *App) ApplyVirtualService(ctx context.Context) error {
	output, _ := s.DescribeVirtualService(ctx)
	if output.VirtualService == nil {
		r := &createVirtualService{s}
		input, err := r.Load(s.config.VirtualServices[0]) // FIXME: Allow for multiple file support

		_, err = s.appmesh.CreateVirtualService(ctx, input)
		if err != nil {
			return err
		}
	} else {
		r := &updateVirtualService{s}
		input, err := r.Load(s.config.VirtualServices[0]) // FIXME: Allow for multiple file support

		_, err = s.appmesh.UpdateVirtualService(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}
