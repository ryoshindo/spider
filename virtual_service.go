package spider

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

type VirtualServiceDefinition[T VirtualServiceDefinitionInput] interface {
	Load(path string) (T, error)
}

type VirtualServiceDefinitionInput interface {
	*appmesh.DescribeVirtualServiceInput | *appmesh.CreateVirtualServiceInput | *appmesh.UpdateVirtualServiceInput
}

type describeVirtualService struct {
	*App
}

func (v *describeVirtualService) Load(path string) (*appmesh.DescribeVirtualServiceInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.DescribeVirtualServiceInput{}, err
	}

	c := struct {
		VirtualServiceDefinition json.RawMessage `json:"virtualServiceDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.DescribeVirtualServiceInput{}, err
	}

	if c.VirtualServiceDefinition != nil {
		src = c.VirtualServiceDefinition
	}

	input := appmesh.DescribeVirtualServiceInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}

	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.DescribeVirtualServiceInput{}, err
	}

	return &input, nil
}

type createVirtualService struct {
	*App
}

func (v *createVirtualService) Load(path string) (*appmesh.CreateVirtualServiceInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.CreateVirtualServiceInput{}, err
	}

	c := struct {
		VirtualServiceDefinition json.RawMessage `json:"virtualServiceDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.CreateVirtualServiceInput{}, err
	}

	if c.VirtualServiceDefinition != nil {
		src = c.VirtualServiceDefinition
	}

	input := appmesh.CreateVirtualServiceInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}

	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.CreateVirtualServiceInput{}, err
	}

	return &input, nil
}

type updateVirtualService struct {
	*App
}

func (v *updateVirtualService) Load(path string) (*appmesh.UpdateVirtualServiceInput, error) {
	src, err := v.readDefinitionFile(path)
	if err != nil {
		return &appmesh.UpdateVirtualServiceInput{}, err
	}

	c := struct {
		VirtualServiceDefinition json.RawMessage `json:"virtualServiceDefinition"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(&c); err != nil {
		return &appmesh.UpdateVirtualServiceInput{}, err
	}

	if c.VirtualServiceDefinition != nil {
		src = c.VirtualServiceDefinition
	}

	input := appmesh.UpdateVirtualServiceInput{
		MeshName:  aws.String(v.config.Mesh.Name),
		MeshOwner: aws.String(v.config.Mesh.Owner),
	}

	if err := v.UnmarshalJsonForStruct(src, &input, path); err != nil {
		return &appmesh.UpdateVirtualServiceInput{}, err
	}

	return &input, nil
}
