package spider

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

type RouteDefinition[T RouteDefinitionInput] interface {
	Load(path, virtualRouterName string) (T, error)
}

type RouteDefinitionInput interface {
	*appmesh.DescribeRouteInput | *appmesh.CreateRouteInput | *appmesh.UpdateRouteInput | *appmesh.DeleteRouteInput
}

type DescribeRoute struct {
	*App
}

func (r *DescribeRoute) Load(path, virtualRouterName string) (*appmesh.DescribeRouteInput, error) {
	src, err := r.readDefinitionFile(path)
	if err != nil {
		return &appmesh.DescribeRouteInput{}, err
	}

	c := struct {
		RouteDefinition json.RawMessage `json:"routeDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.DescribeRouteInput{}, err
	}

	if c.RouteDefinition != nil {
		src = c.RouteDefinition
	}

	input := appmesh.DescribeRouteInput{
		MeshName:          aws.String(r.config.Mesh.Name),
		MeshOwner:         aws.String(r.config.Mesh.Owner),
		VirtualRouterName: &virtualRouterName,
	}

	if err := r.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.DescribeRouteInput{}, err
	}

	return &input, nil
}

type CreateRoute struct {
	*App
}

func (r *CreateRoute) Load(path, virtualRouterName string) (*appmesh.CreateRouteInput, error) {
	src, err := r.readDefinitionFile(path)
	if err != nil {
		return &appmesh.CreateRouteInput{}, err
	}

	c := struct {
		RouteDefinition json.RawMessage `json:"routeDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.CreateRouteInput{}, err
	}

	if c.RouteDefinition != nil {
		src = c.RouteDefinition
	}

	input := appmesh.CreateRouteInput{
		MeshName:          aws.String(r.config.Mesh.Name),
		MeshOwner:         aws.String(r.config.Mesh.Owner),
		VirtualRouterName: &virtualRouterName,
	}

	if err := r.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.CreateRouteInput{}, err
	}

	return &input, nil
}

type UpdateRoute struct {
	*App
}

func (r *UpdateRoute) Load(path, virtualRouterName string) (*appmesh.UpdateRouteInput, error) {
	src, err := r.readDefinitionFile(path)
	if err != nil {
		return &appmesh.UpdateRouteInput{}, err
	}

	c := struct {
		RouteDefinition json.RawMessage `json:"routeDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.UpdateRouteInput{}, err
	}

	if c.RouteDefinition != nil {
		src = c.RouteDefinition
	}

	input := appmesh.UpdateRouteInput{
		MeshName:          aws.String(r.config.Mesh.Name),
		MeshOwner:         aws.String(r.config.Mesh.Owner),
		VirtualRouterName: &virtualRouterName,
	}

	if err := r.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.UpdateRouteInput{}, err
	}

	return &input, nil
}

type DeleteRoute struct {
	*App
}

func (r *DeleteRoute) Load(path, virtualRouterName string) (*appmesh.DeleteRouteInput, error) {
	src, err := r.readDefinitionFile(path)
	if err != nil {
		return &appmesh.DeleteRouteInput{}, err
	}

	c := struct {
		RouteDefinition json.RawMessage `json:"routeDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.DeleteRouteInput{}, err
	}

	if c.RouteDefinition != nil {
		src = c.RouteDefinition
	}

	input := appmesh.DeleteRouteInput{
		MeshName:          aws.String(r.config.Mesh.Name),
		MeshOwner:         aws.String(r.config.Mesh.Owner),
		VirtualRouterName: &virtualRouterName,
	}

	if err := r.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.DeleteRouteInput{}, err
	}

	return &input, nil
}
