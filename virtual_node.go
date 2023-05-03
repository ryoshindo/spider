package spider

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

type VirtualNodeDefinition[T VirtualNodeDefinitionInput] interface {
	Load(path string) (T, error)
}

type VirtualNodeDefinitionInput interface {
	*appmesh.DescribeVirtualNodeInput | *appmesh.CreateVirtualNodeInput | *appmesh.UpdateVirtualNodeInput | *appmesh.DeleteVirtualNodeInput
}

type describeVirtualNode struct {
	*App
}

func (v *describeVirtualNode) Load(path string) (*appmesh.DescribeVirtualNodeInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.DescribeVirtualNodeInput{}, err
	}

	c := struct {
		VirtualNodeDefinition json.RawMessage `json:"virtualNodeDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.DescribeVirtualNodeInput{}, err
	}

	if c.VirtualNodeDefinition != nil {
		src = c.VirtualNodeDefinition
	}

	input := appmesh.DescribeVirtualNodeInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}
	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.DescribeVirtualNodeInput{}, err
	}

	return &input, nil
}

type createVirtualNode struct {
	*App
}

func (v *createVirtualNode) Load(path string) (*appmesh.CreateVirtualNodeInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.CreateVirtualNodeInput{}, err
	}

	c := struct {
		VirtualNodeDefinition json.RawMessage `json:"virtualNodeDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.CreateVirtualNodeInput{}, err
	}

	if c.VirtualNodeDefinition != nil {
		src = c.VirtualNodeDefinition
	}

	input := appmesh.CreateVirtualNodeInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}
	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.CreateVirtualNodeInput{}, err
	}

	if len(input.Tags) == 0 {
		input.Tags = nil
	}

	return &input, nil
}

type updateVirtualNode struct {
	*App
}

func (v *updateVirtualNode) Load(path string) (*appmesh.UpdateVirtualNodeInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.UpdateVirtualNodeInput{}, err
	}

	c := struct {
		VirtualNodeDefinition json.RawMessage `json:"virtualNodeDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.UpdateVirtualNodeInput{}, err
	}

	if c.VirtualNodeDefinition != nil {
		src = c.VirtualNodeDefinition
	}

	input := appmesh.UpdateVirtualNodeInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}
	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.UpdateVirtualNodeInput{}, err
	}

	return &input, nil
}

type deleteVirtualNode struct {
	*App
}

func (v *deleteVirtualNode) Load(path string) (*appmesh.DeleteVirtualNodeInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.DeleteVirtualNodeInput{}, err
	}

	c := struct {
		VirtualNodeDefinition json.RawMessage `json:"virtualNodeDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.DeleteVirtualNodeInput{}, err
	}

	if c.VirtualNodeDefinition != nil {
		src = c.VirtualNodeDefinition
	}

	input := appmesh.DeleteVirtualNodeInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}

	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.DeleteVirtualNodeInput{}, err
	}

	return &input, nil
}
