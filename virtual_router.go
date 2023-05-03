package spider

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

type VirtualRouterDefinition[T VirtualRouterDefinitionInput] interface {
	Load(path string) (T, error)
}

type VirtualRouterDefinitionInput interface {
	*appmesh.DescribeVirtualRouterInput | *appmesh.CreateVirtualRouterInput | *appmesh.UpdateVirtualRouterInput
}

type describeVirtualRouter struct {
	*App
}

func (v *describeVirtualRouter) Load(path string) (*appmesh.DescribeVirtualRouterInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.DescribeVirtualRouterInput{}, err
	}

	c := struct {
		VirtualRouterDefinition json.RawMessage `json:"virtualRouterDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.DescribeVirtualRouterInput{}, err
	}

	if c.VirtualRouterDefinition != nil {
		src = c.VirtualRouterDefinition
	}

	input := appmesh.DescribeVirtualRouterInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}
	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.DescribeVirtualRouterInput{}, err
	}

	return &input, nil
}

type createVirtualRouter struct {
	*App
}

func (v *createVirtualRouter) Load(path string) (*appmesh.CreateVirtualRouterInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.CreateVirtualRouterInput{}, err
	}

	c := struct {
		VirtualRouterDefinition json.RawMessage `json:"virtualRouterDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.CreateVirtualRouterInput{}, err
	}

	if c.VirtualRouterDefinition != nil {
		src = c.VirtualRouterDefinition
	}

	input := appmesh.CreateVirtualRouterInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}

	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.CreateVirtualRouterInput{}, err
	}

	return &input, nil
}

type updateVirtualRouter struct {
	*App
}

func (v *updateVirtualRouter) Load(path string) (*appmesh.UpdateVirtualRouterInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.UpdateVirtualRouterInput{}, err
	}

	c := struct {
		VirtualRouterDefinition json.RawMessage `json:"virtualRouterDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.UpdateVirtualRouterInput{}, err
	}

	if c.VirtualRouterDefinition != nil {
		src = c.VirtualRouterDefinition
	}

	input := appmesh.UpdateVirtualRouterInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}

	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.UpdateVirtualRouterInput{}, err
	}

	return &input, nil
}