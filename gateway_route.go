package spider

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

type GatewayRouteDefinition[T RouteDefinitionInput] interface {
	Load(path, virtualGatewayName string) (T, error)
}

type GatewayRouteDefinitionInput interface {
	*appmesh.DescribeGatewayRouteInput | *appmesh.CreateGatewayRouteInput | *appmesh.UpdateGatewayRouteInput | *appmesh.DeleteGatewayRouteInput
}

type DescribeGatewayRoute struct {
	*App
}

func (r *DescribeGatewayRoute) Load(path, virtualGatewayName string) (*appmesh.DescribeGatewayRouteInput, error) {
	src, err := r.readDefinitionFile(path)
	if err != nil {
		return &appmesh.DescribeGatewayRouteInput{}, err
	}

	c := struct {
		GatewayRouteDefinition json.RawMessage `json:"gatewayRouteDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.DescribeGatewayRouteInput{}, err
	}

	if c.GatewayRouteDefinition != nil {
		src = c.GatewayRouteDefinition
	}

	input := appmesh.DescribeGatewayRouteInput{
		MeshName:           aws.String(r.config.Mesh.Name),
		MeshOwner:          aws.String(r.config.Mesh.Owner),
		VirtualGatewayName: &virtualGatewayName,
	}

	if err := r.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.DescribeGatewayRouteInput{}, err
	}

	return &input, nil
}

type CreateGatewayRoute struct {
	*App
}

func (r *CreateGatewayRoute) Load(path, virtualGatewayName string) (*appmesh.CreateGatewayRouteInput, error) {
	src, err := r.readDefinitionFile(path)
	if err != nil {
		return &appmesh.CreateGatewayRouteInput{}, err
	}

	c := struct {
		GatewayRouteDefinition json.RawMessage `json:"gatewayRouteDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.CreateGatewayRouteInput{}, err
	}

	if c.GatewayRouteDefinition != nil {
		src = c.GatewayRouteDefinition
	}

	input := appmesh.CreateGatewayRouteInput{
		MeshName:           aws.String(r.config.Mesh.Name),
		MeshOwner:          aws.String(r.config.Mesh.Owner),
		VirtualGatewayName: &virtualGatewayName,
	}

	if err := r.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.CreateGatewayRouteInput{}, err
	}

	return &input, nil
}

type UpdateGatewayRoute struct {
	*App
}

func (r *UpdateGatewayRoute) Load(path, virtualGatewayName string) (*appmesh.UpdateGatewayRouteInput, error) {
	src, err := r.readDefinitionFile(path)
	if err != nil {
		return &appmesh.UpdateGatewayRouteInput{}, err
	}

	c := struct {
		GatewayRouteDefinition json.RawMessage `json:"gatewayRouteDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.UpdateGatewayRouteInput{}, err
	}

	if c.GatewayRouteDefinition != nil {
		src = c.GatewayRouteDefinition
	}

	input := appmesh.UpdateGatewayRouteInput{
		MeshName:           aws.String(r.config.Mesh.Name),
		MeshOwner:          aws.String(r.config.Mesh.Owner),
		VirtualGatewayName: &virtualGatewayName,
	}

	if err := r.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.UpdateGatewayRouteInput{}, err
	}

	return &input, nil
}

type DeleteGatewayRoute struct {
	*App
}

func (r *DeleteGatewayRoute) Load(path, virtualGatewayName string) (*appmesh.DeleteGatewayRouteInput, error) {
	src, err := r.readDefinitionFile(path)
	if err != nil {
		return &appmesh.DeleteGatewayRouteInput{}, err
	}

	c := struct {
		GatewayRouteDefinition json.RawMessage `json:"gatewayRouteDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.DeleteGatewayRouteInput{}, err
	}

	if c.GatewayRouteDefinition != nil {
		src = c.GatewayRouteDefinition
	}

	input := appmesh.DeleteGatewayRouteInput{
		MeshName:           aws.String(r.config.Mesh.Name),
		MeshOwner:          aws.String(r.config.Mesh.Owner),
		VirtualGatewayName: &virtualGatewayName,
	}

	if err := r.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.DeleteGatewayRouteInput{}, err
	}

	return &input, nil
}
