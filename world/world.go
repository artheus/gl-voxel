package world

import "github.com/faiface/glhf"

type World struct {
	Chunks map[Vec2]*Chunk
	MeshCache *MeshCache
}

func NewWorld(blockShader *glhf.Shader) *World {
	chunkmap := make(map[Vec2]*Chunk)
	meshCache := NewMeshCache(blockShader)

	for x := -5; x < 5; x++ {
		for z := -5; z < 5; z++ {
			chunkmap[Vec2{x,z}] = NewChunk(Vec2{x, z}, meshCache)
		}
	}

	return &World{
		Chunks: chunkmap,
		MeshCache: meshCache,
	}
}

func (w *World) EmptyAt(pos Vec3) bool {
	chunkId := pos.Chunkid()
	if c, cok := w.Chunks[chunkId]; cok {
		if _, ok := c.Blocks[pos]; ok {
			return false
		}
	}

	return true
}
