package spider

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
)

type MeshDefenition[T MeshDefenitionInput] interface {
	Load(path string) (T, error)
}

type MeshDefenitionInput interface {
	*appmesh.DescribeMeshInput | *appmesh.CreateMeshInput | *appmesh.UpdateMeshInput | *appmesh.DeleteMeshInput
}

type DescribeMesh struct {
	*App
}

func (m *DescribeMesh) Load() (*appmesh.DescribeMeshInput, error) {
	input := appmesh.DescribeMeshInput{
		MeshName:  aws.String(m.config.Mesh.Name),
		MeshOwner: aws.String(m.config.Mesh.Owner),
	}

	return &input, nil
}
