package spider

import (
	"context"
	"fmt"
	"os"

	"github.com/ddddddO/gtree"
)

type TreeOption struct{}

func (s *App) Tree(ctx context.Context, opt TreeOption) error {
	var root *gtree.Node
	_, meshOutput, err := s.DescribeMesh(ctx)
	if err != nil {
		return err
	}

	root = gtree.NewRoot(fmt.Sprintf("%s (Owner: %s)", *meshOutput.Mesh.MeshName, *meshOutput.Mesh.Metadata.MeshOwner))

	for _, virtualGateway := range s.config.VirtualGateways {
		_, vgOutput, _ := s.DescribeVirtualGateway(ctx, virtualGateway.Path)
		var vgTreeNode *gtree.Node
		if vgOutput.VirtualGateway != nil {
			vgTreeNode = root.Add(*vgOutput.VirtualGateway.VirtualGatewayName)
		}

		for _, gatewayRoute := range virtualGateway.GatewayRoutes {
			_, grOutput, _ := s.DescribeGatewayRoute(ctx, gatewayRoute.Path, virtualGateway.Path)
			if grOutput.GatewayRoute != nil {
				vgTreeNode.Add(*grOutput.GatewayRoute.GatewayRouteName)
			}
		}
	}

	for _, virtualService := range s.config.VirtualServices {
		_, vsOutput, _ := s.DescribeVirtualService(ctx, virtualService.Path)
		if vsOutput.VirtualService != nil {
			root.Add(*vsOutput.VirtualService.VirtualServiceName)
		}
	}

	for _, virtualRouter := range s.config.VirtualRouters {
		_, vrOutput, _ := s.DescribeVirtualRouter(ctx, virtualRouter.Path)
		var vrTreeNode *gtree.Node
		if vrOutput.VirtualRouter != nil {
			vrTreeNode = root.Add(*vrOutput.VirtualRouter.VirtualRouterName)
		}

		for _, route := range virtualRouter.Routes {
			_, rOutput, _ := s.DescribeRoute(ctx, route.Path, virtualRouter.Path)
			if rOutput.Route != nil {
				vrTreeNode.Add(*rOutput.Route.RouteName)
			}
		}
	}

	for _, virtualNode := range s.config.VirtualNodes {
		_, vnOutput, _ := s.DescribeVirtualNode(ctx, virtualNode.Path)
		if vnOutput.VirtualNode != nil {
			root.Add(*vnOutput.VirtualNode.VirtualNodeName)
		}
	}

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		return err
	}

	return nil
}
