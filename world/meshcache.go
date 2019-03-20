package world

import "github.com/faiface/glhf"

type MeshCache struct {
	shader *glhf.Shader
	cacheSlice []*glhf.VertexSlice
}

func (c *MeshCache) Add(x,y,z int, block *Block, showFaces ShowFaces) {
	c.cacheSlice = append(c.cacheSlice, createCubeSlice(c.shader, block, showFaces, float32(x),float32(y),float32(z)))
}

func  createCubeSlice(shader *glhf.Shader, block *Block, showFaces ShowFaces, x,y,z float32) *glhf.VertexSlice {
	vdata := block.VertexData(x,y,z, showFaces)
	s := glhf.MakeVertexSlice(shader, len(vdata)/8, 36)
	s.Begin()
	s.SetVertexData(vdata)
	s.End()

	return s
}

func NewMeshCache(shader *glhf.Shader) *MeshCache {
	return &MeshCache{
		shader: shader,
	}
}