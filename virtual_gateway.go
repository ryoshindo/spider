package spider

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

type VirtualGatewayDefinition[T VirtualGatewayDefinitionInput] interface {
	Load(path string) (T, error)
}

type VirtualGatewayDefinitionInput interface {
	*appmesh.DescribeVirtualGatewayInput | *appmesh.CreateVirtualGatewayInput | *appmesh.UpdateVirtualGatewayInput | *appmesh.DeleteVirtualGatewayInput
}

type describeVirtualGateway struct {
	*App
}

func (v *describeVirtualGateway) Load(path string) (*appmesh.DescribeVirtualGatewayInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.DescribeVirtualGatewayInput{}, err
	}

	c := struct {
		VirtualGatewayDefinition json.RawMessage `json:"virtualGatewayDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.DescribeVirtualGatewayInput{}, err
	}

	if c.VirtualGatewayDefinition != nil {
		src = c.VirtualGatewayDefinition
	}

	input := appmesh.DescribeVirtualGatewayInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}

	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.DescribeVirtualGatewayInput{}, err
	}

	return &input, nil
}

type createVirtualGateway struct {
	*App
}

func (v *createVirtualGateway) Load(path string) (*appmesh.CreateVirtualGatewayInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.CreateVirtualGatewayInput{}, err
	}

	c := struct {
		VirtualGatewayDefinition json.RawMessage `json:"virtualGatewayDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.CreateVirtualGatewayInput{}, err
	}

	if c.VirtualGatewayDefinition != nil {
		src = c.VirtualGatewayDefinition
	}

	input := appmesh.CreateVirtualGatewayInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}

	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.CreateVirtualGatewayInput{}, err
	}

	return &input, nil
}

type updateVirtualGateway struct {
	*App
}

func (v *updateVirtualGateway) Load(path string) (*appmesh.UpdateVirtualGatewayInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.UpdateVirtualGatewayInput{}, err
	}

	c := struct {
		VirtualGatewayDefinition json.RawMessage `json:"virtualGatewayDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.UpdateVirtualGatewayInput{}, err
	}

	if c.VirtualGatewayDefinition != nil {
		src = c.VirtualGatewayDefinition
	}

	input := appmesh.UpdateVirtualGatewayInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}

	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.UpdateVirtualGatewayInput{}, err
	}

	return &input, nil
}

type deleteVirtualGateway struct {
	*App
}

func (v *deleteVirtualGateway) Load(path string) (*appmesh.DeleteVirtualGatewayInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.DeleteVirtualGatewayInput{}, err
	}

	c := struct {
		VirtualGatewayDefinition json.RawMessage `json:"virtualGatewayDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.DeleteVirtualGatewayInput{}, err
	}

	if c.VirtualGatewayDefinition != nil {
		src = c.VirtualGatewayDefinition
	}

	input := appmesh.DeleteVirtualGatewayInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}

	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.DeleteVirtualGatewayInput{}, err
	}

	return &input, nil
}
