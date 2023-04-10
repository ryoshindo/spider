package spider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
)

func (s *App) Deploy(ctx context.Context) error {
	input := &appmesh.CreateVirtualNodeInput{
		MeshName:        aws.String(appmeshName),     // FIXME
		VirtualNodeName: aws.String(virtualNodeName), // FIXME
		Spec: &types.VirtualNodeSpec{
			Listeners: []types.Listener{},
		},
	}

	_, err := s.appmesh.CreateVirtualNode(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
