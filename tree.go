package spider

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
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

			var virtualServiceName string
			switch {
			case grOutput.GatewayRoute.Spec.HttpRoute != nil:
				virtualServiceName = *grOutput.GatewayRoute.Spec.HttpRoute.Action.Target.VirtualService.VirtualServiceName
				grTreeNode = grTreeNode.Add(fmt.Sprintf("[VirtualService/HttpRoute] %s", *grOutput.GatewayRoute.Spec.HttpRoute.Action.Target.VirtualService.VirtualServiceName))
			case grOutput.GatewayRoute.Spec.Http2Route != nil:
				virtualServiceName = *grOutput.GatewayRoute.Spec.Http2Route.Action.Target.VirtualService.VirtualServiceName
				grTreeNode = grTreeNode.Add(fmt.Sprintf("[VirtualService/Http2Route] %s", *grOutput.GatewayRoute.Spec.Http2Route.Action.Target.VirtualService.VirtualServiceName))
			case grOutput.GatewayRoute.Spec.GrpcRoute != nil:
				virtualServiceName = *grOutput.GatewayRoute.Spec.GrpcRoute.Action.Target.VirtualService.VirtualServiceName
				grTreeNode = grTreeNode.Add(fmt.Sprintf("[VirtualService/GrpcRoute] %s", *grOutput.GatewayRoute.Spec.GrpcRoute.Action.Target.VirtualService.VirtualServiceName))
			}

			virtualService, err := s.appmesh.DescribeVirtualService(ctx, &appmesh.DescribeVirtualServiceInput{
				MeshName:           aws.String(s.config.Mesh.Name),
				MeshOwner:          aws.String(s.config.Mesh.Owner),
				VirtualServiceName: aws.String(virtualServiceName),
			})
			if err != nil {
				return err
			}

			var vsTreeNode *gtree.Node
			switch v := virtualService.VirtualService.Spec.Provider.(type) {
			case *types.VirtualServiceProviderMemberVirtualNode:
				virtualNodeName := *v.Value.VirtualNodeName
				vsTreeNode = grTreeNode.Add(fmt.Sprintf("[VirtualNode] %s", virtualNodeName))
				virtualNode, err := s.appmesh.DescribeVirtualNode(ctx, &appmesh.DescribeVirtualNodeInput{
					MeshName:        aws.String(s.config.Mesh.Name),
					MeshOwner:       aws.String(s.config.Mesh.Owner),
					VirtualNodeName: aws.String(virtualNodeName),
				})
				if err != nil {
					return err
				}
				vsTreeNode.Add(fmt.Sprintf("[VirtualNode] %s", *virtualNode.VirtualNode.VirtualNodeName))
			case *types.VirtualServiceProviderMemberVirtualRouter:
				virtualRouterName := *v.Value.VirtualRouterName
				vsTreeNode = grTreeNode.Add(fmt.Sprintf("[VirtualRouter] %s", virtualRouterName))
				virtualRouter, err := s.appmesh.DescribeVirtualRouter(ctx, &appmesh.DescribeVirtualRouterInput{
					MeshName:          aws.String(s.config.Mesh.Name),
					MeshOwner:         aws.String(s.config.Mesh.Owner),
					VirtualRouterName: aws.String(virtualRouterName),
				})
				if err != nil {
					return err
				}
				vsTreeNode.Add(fmt.Sprintf("[VirtualRouter] %s", *virtualRouter.VirtualRouter.VirtualRouterName))

				virtualRouterRouteMap, err := s.getVirtualRouterRouteMap()
				if err != nil {
					return err
				}
				for _, routes := range virtualRouterRouteMap {
					for _, route := range routes {
						vsTreeNode.Add(fmt.Sprintf("[Route] %s", route))
					}
				}
			default:
				vsTreeNode = grTreeNode
			}
		}
	}

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		return err
	}

	return nil
}
