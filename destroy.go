package spider

import "context"

func (s *App) Destroy(ctx context.Context) error {
	if err := s.DestroyGatewayRoute(ctx); err != nil {
		return err
	}

	if err := s.DestroyVirtualGateway(ctx); err != nil {
		return err
	}

	if err := s.DestroyVirtualService(ctx); err != nil {
		return err
	}

	if err := s.DestroyRoute(ctx); err != nil {
		return err
	}

	if err := s.DestroyVirtualRouter(ctx); err != nil {
		return err
	}

	if err := s.DestroyVirtualNode(ctx); err != nil {
		return err
	}

	return nil
}

func (s *App) DestroyVirtualNode(ctx context.Context) error {
	for _, virtualNode := range s.config.VirtualNodes {
		vn := &DeleteVirtualNode{s}
		input, err := vn.Load(virtualNode)
		if err != nil {
			return err
		}

		_, err = s.appmesh.DeleteVirtualNode(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *App) DestroyVirtualRouter(ctx context.Context) error {
	for _, virtualRouter := range s.config.VirtualRouters {
		vr := &DeleteVirtualRouter{s}
		input, err := vr.Load(virtualRouter.Path)
		if err != nil {
			return err
		}

		_, err = s.appmesh.DeleteVirtualRouter(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *App) DestroyRoute(ctx context.Context) error {
	for _, virtualRouter := range s.config.VirtualRouters {
		for _, route := range virtualRouter.Routes {
			output, _ := s.DescribeRoute(ctx, route, virtualRouter.Path)
			r := &DeleteRoute{s}
			input, err := r.Load(route, *output.Route.VirtualRouterName)
			if err != nil {
				return err
			}

			_, err = s.appmesh.DeleteRoute(ctx, input)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *App) DestroyVirtualService(ctx context.Context) error {
	for _, virtualService := range s.config.VirtualServices {
		vs := &DeleteVirtualService{s}
		input, err := vs.Load(virtualService)
		if err != nil {
			return err
		}

		_, err = s.appmesh.DeleteVirtualService(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *App) DestroyVirtualGateway(ctx context.Context) error {
	for _, virtualGateway := range s.config.VirtualGateways {
		vg := &DeleteVirtualGateway{s}
		input, err := vg.Load(virtualGateway.Path)
		if err != nil {
			return err
		}

		_, err = s.appmesh.DeleteVirtualGateway(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *App) DestroyGatewayRoute(ctx context.Context) error {
	for _, virtualGateway := range s.config.VirtualGateways {
		for _, gatewayRoute := range virtualGateway.GatewayRoutes {
			output, _ := s.DescribeGatewayRoute(ctx, gatewayRoute, virtualGateway.Path)
			gr := &DeleteGatewayRoute{s}
			input, err := gr.Load(gatewayRoute, *output.GatewayRoute.VirtualGatewayName)
			if err != nil {
				return err
			}

			_, err = s.appmesh.DeleteGatewayRoute(ctx, input)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
