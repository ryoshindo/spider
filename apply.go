package spider

import (
	"context"
	"fmt"

	"golang.org/x/exp/slog"
)

func (s *App) Apply(ctx context.Context) error {
	s.Log(slog.LevelInfo, "Applying AppMesh resources...")

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

	s.Log(slog.LevelInfo, "AppMesh resources applied successfully!")

	return nil
}

func (s *App) ApplyVirtualNode(ctx context.Context) error {
	for _, virtualNode := range s.config.VirtualNodes {
		_, dOutput, _ := s.DescribeVirtualNode(ctx, virtualNode.Path)
		if dOutput.VirtualNode == nil {
			vn := &CreateVirtualNode{s}
			input, err := vn.Load(virtualNode.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.CreateVirtualNode(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualNode: Virtual Node '%s' in '%s' failed to create", *input.VirtualNodeName, virtualNode.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualNode: Virtual Node '%s' in '%s' created", *input.VirtualNodeName, virtualNode.Path))
		} else {
			vn := &UpdateVirtualNode{s}
			input, err := vn.Load(virtualNode.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.UpdateVirtualNode(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualNode: Virtual Node '%s' in '%s' failed to update", *input.VirtualNodeName, virtualNode.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualNode: Virtual Node '%s' in '%s' updated", *input.VirtualNodeName, virtualNode.Path))
		}
	}

	return nil
}

func (s *App) ApplyVirtualRouter(ctx context.Context) error {
	for _, virtualRouter := range s.config.VirtualRouters {
		_, dOutput, _ := s.DescribeVirtualRouter(ctx, virtualRouter.Path)
		if dOutput.VirtualRouter == nil {
			vr := &CreateVirtualRouter{s}
			input, err := vr.Load(virtualRouter.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.CreateVirtualRouter(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualRouter: Virtual Router '%s' in '%s' failed to create", *input.VirtualRouterName, virtualRouter.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualRouter: Virtual Router '%s' in '%s' created", *input.VirtualRouterName, virtualRouter.Path))
		} else {
			vr := &UpdateVirtualRouter{s}
			input, err := vr.Load(virtualRouter.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.UpdateVirtualRouter(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualRouter: Virtual Router '%s' in '%s' failed to update", *input.VirtualRouterName, virtualRouter.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualRouter: Virtual Router '%s' in '%s' updated", *input.VirtualRouterName, virtualRouter.Path))
		}
	}

	return nil
}

func (s *App) ApplyRoute(ctx context.Context) error {
	for _, virtualRouter := range s.config.VirtualRouters {
		for _, route := range virtualRouter.Routes {
			_, dOutput, _ := s.DescribeRoute(ctx, route.Path, virtualRouter.Path)
			if dOutput.Route.Spec == nil {
				r := &CreateRoute{s}
				input, err := r.Load(route.Path, *dOutput.Route.VirtualRouterName)
				if err != nil {
					return err
				}

				_, err = s.appmesh.CreateRoute(ctx, input)
				if err != nil {
					s.Log(slog.LevelError, fmt.Sprintf("Route: Route '%s' in '%s' failed to create", *input.RouteName, route.Path))
					return err
				}
				s.Log(slog.LevelInfo, fmt.Sprintf("Route: Route '%s' in '%s' created", *input.RouteName, route.Path))
			} else {
				r := &UpdateRoute{s}
				input, err := r.Load(route.Path, *dOutput.Route.VirtualRouterName)
				if err != nil {
					return err
				}

				_, err = s.appmesh.UpdateRoute(ctx, input)
				if err != nil {
					s.Log(slog.LevelError, fmt.Sprintf("Route: Route '%s' in '%s' failed to update", *input.RouteName, route.Path))
					return err
				}
				s.Log(slog.LevelInfo, fmt.Sprintf("Route: Route '%s' in '%s' updated", *input.RouteName, route.Path))
			}
		}
	}

	return nil
}

func (s *App) ApplyVirtualService(ctx context.Context) error {
	for _, virtualService := range s.config.VirtualServices {
		_, dOutput, _ := s.DescribeVirtualService(ctx, virtualService.Path)
		if dOutput.VirtualService == nil {
			vs := &CreateVirtualService{s}
			input, err := vs.Load(virtualService.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.CreateVirtualService(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualService: Virtual Service '%s' in '%s' failed to create", *input.VirtualServiceName, virtualService.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualService: Virtual Service '%s' in '%s' created", *input.VirtualServiceName, virtualService.Path))
		} else {
			vs := &UpdateVirtualService{s}
			input, err := vs.Load(virtualService.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.UpdateVirtualService(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualService: Virtual Service '%s' in '%s' failed to update", *input.VirtualServiceName, virtualService.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualService: Virtual Service '%s' in '%s' updated", *input.VirtualServiceName, virtualService.Path))
		}
	}

	return nil
}

func (s *App) ApplyVirtualGateway(ctx context.Context) error {
	for _, virtualGateway := range s.config.VirtualGateways {
		_, dOutput, _ := s.DescribeVirtualGateway(ctx, virtualGateway.Path)
		if dOutput.VirtualGateway == nil {
			vg := &CreateVirtualGateway{s}
			input, err := vg.Load(virtualGateway.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.CreateVirtualGateway(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualGateway: Virtual Gateway '%s' in '%s' failed to create", *input.VirtualGatewayName, virtualGateway.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualGateway: Virtual Gateway '%s' in '%s' created", *input.VirtualGatewayName, virtualGateway.Path))
		} else {
			vg := &UpdateVirtualGateway{s}
			input, err := vg.Load(virtualGateway.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.UpdateVirtualGateway(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualGateway: Virtual Gateway '%s' in '%s' failed to update", *input.VirtualGatewayName, virtualGateway.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualGateway: Virtual Gateway '%s' in '%s' updated", *input.VirtualGatewayName, virtualGateway.Path))
		}
	}

	return nil
}

func (s *App) ApplyGatewayRoute(ctx context.Context) error {
	for _, virtualGateway := range s.config.VirtualGateways {
		for _, gatewayRoute := range virtualGateway.GatewayRoutes {
			_, dOutput, _ := s.DescribeGatewayRoute(ctx, gatewayRoute.Path, virtualGateway.Path)
			if dOutput.GatewayRoute.Spec == nil {
				gr := &CreateGatewayRoute{s}
				input, err := gr.Load(gatewayRoute.Path, *dOutput.GatewayRoute.VirtualGatewayName)
				if err != nil {
					return err
				}

				_, err = s.appmesh.CreateGatewayRoute(ctx, input)
				if err != nil {
					s.Log(slog.LevelError, fmt.Sprintf("GatewayRoute: Gateway Route '%s' in '%s' failed to create", *input.GatewayRouteName, gatewayRoute.Path))
					return err
				}
				s.Log(slog.LevelInfo, fmt.Sprintf("GatewayRoute: Gateway Route '%s' in '%s' created", *input.GatewayRouteName, gatewayRoute.Path))
			} else {
				gr := &UpdateGatewayRoute{s}
				input, err := gr.Load(gatewayRoute.Path, *dOutput.GatewayRoute.VirtualGatewayName)
				if err != nil {
					return err
				}

				_, err = s.appmesh.UpdateGatewayRoute(ctx, input)
				if err != nil {
					s.Log(slog.LevelError, fmt.Sprintf("GatewayRoute: Gateway Route '%s' in '%s' failed to update", *input.GatewayRouteName, gatewayRoute.Path))
					return err
				}
				s.Log(slog.LevelInfo, fmt.Sprintf("GatewayRoute: Gateway Route '%s' in '%s' updated", *input.GatewayRouteName, gatewayRoute.Path))
			}
		}
	}

	return nil
}
