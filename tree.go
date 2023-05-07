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

	root = gtree.NewRoot(fmt.Sprintf("[Mesh] %s (Owner: %s)", *meshOutput.Mesh.MeshName, *meshOutput.Mesh.Metadata.MeshOwner))

	for _, virtualGateway := range s.config.VirtualGateways {
		_, vgOutput, err := s.DescribeVirtualGateway(ctx, virtualGateway.Path)
		if err != nil {
			return err
		}

		var vgTreeNode *gtree.Node
		if vgOutput.VirtualGateway != nil {
			vgTreeNode = root.Add(fmt.Sprintf("[VirtualGateway] %s", *vgOutput.VirtualGateway.VirtualGatewayName))
		}

		for _, gatewayRoute := range virtualGateway.GatewayRoutes {
			_, grOutput, err := s.DescribeGatewayRoute(ctx, gatewayRoute.Path, virtualGateway.Path)
			if err != nil {
				return err
			}

			var grTreeNode *gtree.Node
			if grOutput.GatewayRoute != nil {
				grTreeNode = vgTreeNode.Add(fmt.Sprintf("[GatewayRoute] %s", *grOutput.GatewayRoute.GatewayRouteName))
			}

			if grOutput.GatewayRoute.Spec.HttpRoute != nil {
				grTreeNode.Add(fmt.Sprintf("[VirtualService/HttpRoute] %s", *grOutput.GatewayRoute.Spec.HttpRoute.Action.Target.VirtualService.VirtualServiceName))
			}
			if grOutput.GatewayRoute.Spec.Http2Route != nil {
				grTreeNode.Add(fmt.Sprintf("[VirtualService/Http2Route] %s", *grOutput.GatewayRoute.Spec.Http2Route.Action.Target.VirtualService.VirtualServiceName))
			}
			if grOutput.GatewayRoute.Spec.GrpcRoute != nil {
				grTreeNode.Add(fmt.Sprintf("[VirtualService/GrpcRoute] %s", *grOutput.GatewayRoute.Spec.GrpcRoute.Action.Target.VirtualService.VirtualServiceName))
			}
		}
	}

	for _, virtualService := range s.config.VirtualServices {
		_, vsOutput, _ := s.DescribeVirtualService(ctx, virtualService.Path)
		// var vsTreeNode *gtree.Node
		if vsOutput.VirtualService != nil {
			// vsTreeNode = root.Add(fmt.Sprintf("[VirtualService] %s", *vsOutput.VirtualService.VirtualServiceName))
			root.Add(fmt.Sprintf("[VirtualService] %s", *vsOutput.VirtualService.VirtualServiceName))
		}

		provider := vsOutput.VirtualService.Spec.Provider
		fmt.Println(provider)
	}

	for _, virtualRouter := range s.config.VirtualRouters {
		_, vrOutput, _ := s.DescribeVirtualRouter(ctx, virtualRouter.Path)
		var vrTreeNode *gtree.Node
		if vrOutput.VirtualRouter != nil {
			vrTreeNode = root.Add(fmt.Sprintf("[VirtualRouter] %s", *vrOutput.VirtualRouter.VirtualRouterName))
		}

		for _, route := range virtualRouter.Routes {
			_, rOutput, _ := s.DescribeRoute(ctx, route.Path, virtualRouter.Path)
			if rOutput.Route != nil {
				vrTreeNode.Add(fmt.Sprintf("[Route] %s", *rOutput.Route.RouteName))
			}
		}
	}

	for _, virtualNode := range s.config.VirtualNodes {
		_, vnOutput, _ := s.DescribeVirtualNode(ctx, virtualNode.Path)
		if vnOutput.VirtualNode != nil {
			root.Add(fmt.Sprintf("[VirtualNode] %s", *vnOutput.VirtualNode.VirtualNodeName))
		}
	}

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		return err
	}

	return nil
}
