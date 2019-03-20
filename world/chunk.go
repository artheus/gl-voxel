package world

import "math"

const (
	ChunkHeight int = 256
	ChunkWidth  int = 16

	// Do not change, should always be same
	ChunkDepth = ChunkWidth
)

type Vec2 struct {
	X, Z int
}

type Vec3 struct {
	X, Y, Z int
}

func (v Vec3) Left() Vec3 {
	return Vec3{v.X - 1, v.Y, v.Z}
}
func (v Vec3) Right() Vec3 {
	return Vec3{v.X + 1, v.Y, v.Z}
}
func (v Vec3) Up() Vec3 {
	return Vec3{v.X, v.Y + 1, v.Z}
}
func (v Vec3) Down() Vec3 {
	return Vec3{v.X, v.Y - 1, v.Z}
}
func (v Vec3) Front() Vec3 {
	return Vec3{v.X, v.Y, v.Z + 1}
}
func (v Vec3) Back() Vec3 {
	return Vec3{v.X, v.Y, v.Z - 1}
}
func (v Vec3) Chunkid() Vec2 {
	return Vec2{
		int(math.Floor(float64(v.X) / float64(ChunkWidth))),
		int(math.Floor(float64(v.Z) / float64(ChunkWidth))),
	}
}

type Chunk struct {
	Blocks   map[Vec3]int
}

func NewChunk(pos Vec2, cache *MeshCache) *Chunk {
	return &Chunk{
		Blocks: makeChunkMap(pos, cache),
	}
}

func makeChunkMap(cid Vec2, cache *MeshCache) map[Vec3]int {
	m := make(map[Vec3]int)
	p, q := cid.X, cid.Z
	for dx := 0; dx < ChunkWidth; dx++ {
		for dz := 0; dz < ChunkWidth; dz++ {
			x, z := p*ChunkWidth+dx, q*ChunkWidth+dz
			f := noise2(float32(x)*0.005, float32(z)*0.005, 1, 2.5, 2)

			mh := 30
			h := int(f * float32(mh))
			w := IBlockGrass
			if h <= 12 {
				h = 12
				w = IBlockWater
			}
			// grass and sand
			//for y := 0; y < h; y++ {
			//	m[Pos3{x, y, z}] = w
			//}
			for y := 0; y < h; y++ {
				m[Vec3{x,y,z}] = w
				cache.Add(x,y,z,BlockMap[w],ShowFaces{false,false,true,false,false,false})
			}

			m[Vec3{x,h,z}] = w
			cache.Add(x,h,z,BlockMap[w],ShowFaces{true,true,true,false,true,true})


			// flowers
			//if w == BlockGrass {
			//	if noise2(-float32(x)*0.1, float32(z)*0.1, 4, 0.8, 2) > 0.6 {
			//		m[Vec3{x, h, z}] = BlockPumpkin
			//	}
			//	if noise2(float32(x)*0.05, float32(-z)*0.05, 4, 0.8, 2) > 0.7 {
			//		w := 18 + int(noise2(float32(x)*0.1, float32(z)*0.1, 4, 0.8, 2)*7)
			//		m[Vec3{x, h, z}] = BlockFurnace
			//	}
			//}

			// tree
			//if w == BlockGrass {
			//	ok := true
			//	if dx-4 < 0 || dz-4 < 0 ||
			//		dx+4 > ChunkWidth || dz+4 > ChunkWidth {
			//		ok = false
			//	}
			//	if ok && noise2(float32(x), float32(z), 6, 0.5, 2) > 0.79 {
			//		for y := h + 3; y < h+8; y++ {
			//			for ox := -3; ox <= 3; ox++ {
			//				for oz := -3; oz <= 3; oz++ {
			//					d := ox*ox + oz*oz + (y-h-4)*(y-h-4)
			//					if d < 11 {
			//						m[Pos3{x + ox, y, z + oz}] = BlockLeaves
			//					}
			//				}
			//			}
			//		}
			//		for y := h; y < h+7; y++ {
			//			m[Pos3{x, y, z}] = BlockWood
			//		}
			//	}
			//}

			// cloud
			//for y := 64; y < 72; y++ {
			//	if noise3(float32(x)*0.01, float32(y)*0.1, float32(z)*0.01, 8, 0.5, 2) > 0.69 {
			//		m[Pos3{x, y, z}] = BlockGoldOre
			//	}
			//}
		}
	}
	return m
}