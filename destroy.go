package spider

import (
	"context"
	"fmt"

	"golang.org/x/exp/slog"
)

type DestroyOption struct{}

func (s *App) Destroy(ctx context.Context, _ DestroyOption) error {
	s.Log(slog.LevelInfo, "Destroying AppMesh resources...")

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

	s.Log(slog.LevelInfo, "AppMesh resources destroyed successfully!")

	return nil
}

func (s *App) DestroyVirtualNode(ctx context.Context) error {
	for _, virtualNode := range s.config.VirtualNodes {
		dInput, dOutput, _ := s.DescribeVirtualNode(ctx, virtualNode.Path)
		if dOutput.VirtualNode == nil {
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualNode: Virtual Node '%s' in '%s' has been already deleted", *dInput.VirtualNodeName, virtualNode.Path))
			continue
		} else {
			vn := &DeleteVirtualNode{s}
			input, err := vn.Load(virtualNode.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.DeleteVirtualNode(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualNode: Virtual Node '%s' in '%s' failed to delete", *input.VirtualNodeName, virtualNode.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualNode: Virtual Node '%s' in '%s' deleted", *input.VirtualNodeName, virtualNode.Path))
		}
	}

	return nil
}

func (s *App) DestroyVirtualRouter(ctx context.Context) error {
	for _, virtualRouter := range s.config.VirtualRouters {
		dInput, dOutput, _ := s.DescribeVirtualRouter(ctx, virtualRouter.Path)
		if dOutput.VirtualRouter == nil {
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualRouter: Virtual Router '%s' in '%s' has been already deleted", *dInput.VirtualRouterName, virtualRouter.Path))
			continue
		} else {
			vr := &DeleteVirtualRouter{s}
			input, err := vr.Load(virtualRouter.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.DeleteVirtualRouter(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualRouter: Virtual Router '%s' in '%s' failed to delete", *input.VirtualRouterName, virtualRouter.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualRouter: Virtual Router '%s' in '%s' deleted", *input.VirtualRouterName, virtualRouter.Path))
		}
	}

	return nil
}

func (s *App) DestroyRoute(ctx context.Context) error {
	for _, virtualRouter := range s.config.VirtualRouters {
		for _, route := range virtualRouter.Routes {
			dInput, dOutput, _ := s.DescribeRoute(ctx, route.Path, virtualRouter.Path)
			if dOutput.Route == nil {
				s.Log(slog.LevelInfo, fmt.Sprintf("Route: Route '%s' in '%s' has been already deleted", *dInput.RouteName, route.Path))
				continue
			} else {
				r := &DeleteRoute{s}
				input, err := r.Load(route.Path, *dOutput.Route.VirtualRouterName)
				if err != nil {
					return err
				}

				_, err = s.appmesh.DeleteRoute(ctx, input)
				if err != nil {
					s.Log(slog.LevelError, fmt.Sprintf("Route: Route '%s' in '%s' failed to delete", *input.RouteName, route.Path))
					return err
				}
				s.Log(slog.LevelInfo, fmt.Sprintf("Route: Route '%s' in '%s' deleted", *input.RouteName, route.Path))
			}
		}
	}

	return nil
}

func (s *App) DestroyVirtualService(ctx context.Context) error {
	for _, virtualService := range s.config.VirtualServices {
		dInput, dOutput, _ := s.DescribeVirtualService(ctx, virtualService.Path)
		if dOutput.VirtualService == nil {
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualService: Virtual Service '%s' in '%s' has been already deleted", *dInput.VirtualServiceName, virtualService.Path))
			continue
		} else {
			vs := &DeleteVirtualService{s}
			input, err := vs.Load(virtualService.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.DeleteVirtualService(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualService: Virtual Service '%s' in '%s' failed to delete", *input.VirtualServiceName, virtualService.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualService: Virtual Service '%s' in '%s' deleted", *input.VirtualServiceName, virtualService.Path))
		}
	}

	return nil
}

func (s *App) DestroyVirtualGateway(ctx context.Context) error {
	for _, virtualGateway := range s.config.VirtualGateways {
		dInput, dOutput, _ := s.DescribeVirtualGateway(ctx, virtualGateway.Path)
		if dOutput.VirtualGateway == nil {
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualGateway: Virtual Gateway '%s' in '%s' has been already deleted", *dInput.VirtualGatewayName, virtualGateway.Path))
			continue
		} else {
			vg := &DeleteVirtualGateway{s}
			input, err := vg.Load(virtualGateway.Path)
			if err != nil {
				return err
			}

			_, err = s.appmesh.DeleteVirtualGateway(ctx, input)
			if err != nil {
				s.Log(slog.LevelError, fmt.Sprintf("VirtualGateway: Virtual Gateway '%s' in '%s' failed to delete", *input.VirtualGatewayName, virtualGateway.Path))
				return err
			}
			s.Log(slog.LevelInfo, fmt.Sprintf("VirtualGateway: Virtual Gateway '%s' in '%s' deleted", *input.VirtualGatewayName, virtualGateway.Path))
		}
	}

	return nil
}

func (s *App) DestroyGatewayRoute(ctx context.Context) error {
	for _, virtualGateway := range s.config.VirtualGateways {
		for _, gatewayRoute := range virtualGateway.GatewayRoutes {
			dInput, dOutput, _ := s.DescribeGatewayRoute(ctx, gatewayRoute.Path, virtualGateway.Path)
			if dOutput.GatewayRoute == nil {
				s.Log(slog.LevelInfo, fmt.Sprintf("GatewayRoute: Gateway Route '%s' in '%s' has been already deleted", *dInput.GatewayRouteName, gatewayRoute.Path))
				continue
			} else {
				gr := &DeleteGatewayRoute{s}
				input, err := gr.Load(gatewayRoute.Path, *dOutput.GatewayRoute.VirtualGatewayName)
				if err != nil {
					return err
				}

				_, err = s.appmesh.DeleteGatewayRoute(ctx, input)
				if err != nil {
					s.Log(slog.LevelError, fmt.Sprintf("GatewayRoute: Gateway Route '%s' in '%s' failed to delete", *input.GatewayRouteName, gatewayRoute.Path))
					return err
				}
				s.Log(slog.LevelInfo, fmt.Sprintf("GatewayRoute: Gateway Route '%s' in '%s' deleted", *input.GatewayRouteName, gatewayRoute.Path))
			}
		}
	}

	return nil
}
