package world

const (
	IBlockWater = iota
	IBlockGrass
	IBlockStone
	IBlockDirt
	IBlockPlanks
)

var (
	BlockWater       = NewBlock(61, 61, 61, 61, 61, 61, true)
	BlockGrass       = NewBlock(243, 243, 60, 243, 243, 243, false)
	BlockStone       = NewBlock(241, 241, 241, 241, 241, 241, false)
	BlockDirt        = NewBlock(242, 242, 242, 242, 242, 242, false)
	BlockPlanks      = NewBlock(244, 244, 244, 244, 244, 244, false)
	BlockTNT         = NewBlock(248, 248, 249, 250, 248, 248, false)
	BlockCobblestone = NewBlock(224, 224, 224, 224, 224, 224, false)
	BlockBedrock     = NewBlock(225, 225, 225, 225, 225, 225, false)
	BlockSand        = NewBlock(226, 226, 226, 226, 226, 226, false)
	BlockGravel      = NewBlock(227, 227, 227, 227, 227, 227, false)
	BlockWood        = NewBlock(228, 228, 229, 229, 228, 228, false)
	BlockIron        = NewBlock(230, 230, 230, 230, 230, 230, false)
	BlockGold        = NewBlock(231, 231, 231, 231, 231, 231, false)
	BlockDiamond     = NewBlock(232, 232, 232, 232, 232, 232, false)
	BlockChest       = NewBlock(234, 234, 233, 233, 235, 234, false)
	BlockGoldOre     = NewBlock(208, 208, 208, 208, 208, 208, false)
	BlockIronOre     = NewBlock(209, 209, 209, 209, 209, 209, false)
	BlockCoalOre     = NewBlock(210, 210, 210, 210, 210, 210, false)
	BlockBookshelf   = NewBlock(211, 211, 211, 211, 211, 211, false)
	BlockFurnace     = NewBlock(221, 221, 206, 206, 220, 221, false)
	BlockPumpkin     = NewBlock(134, 134, 150, 150, 135, 134, false)
	BlockLeaves      = NewBlock(27, 27, 27, 27, 27, 27, true)
)

var BlockMap = map[int]*Block {
	IBlockWater: BlockWater,
	IBlockGrass: BlockGrass,
}

const (
	FaceLeft int = iota
	FaceRight
	FaceUp
	FaceDown
	FaceFront
	FaceBack
)

type ShowFaces [6]bool
type FaceTexture [6][2]float32

func MakeFaceTexture(idx int) FaceTexture {
	const textureColums = 16
	var m = 0.5 / float32(textureColums)
	dx, dy := float32(idx%textureColums)*m, float32(idx/textureColums)*m
	n := float32(1 / 2048.0)
	m -= n
	return [6][2]float32{
		{dx + n, dy + n},
		{dx + m, dy + n},
		{dx + m, dy + m},
		{dx + m, dy + m},
		{dx + n, dy + m},
		{dx + n, dy + n},
	}
}

type BlockTexture struct {
	Left, Right FaceTexture
	Up, Down    FaceTexture
	Front, Back FaceTexture
}

type Block struct {
	// Texture id - sides
	texture *BlockTexture

	transparent bool
}

func NewBlock(tl, tr, tu, td, tf, tb int, transparent bool) *Block {
	return &Block{
		texture: &BlockTexture{
			MakeFaceTexture(tl),
			MakeFaceTexture(tr),
			MakeFaceTexture(tu),
			MakeFaceTexture(td),
			MakeFaceTexture(tf),
			MakeFaceTexture(tb),
		},
		transparent: transparent,
	}
}

func (b *Block) Texture() *BlockTexture {
	return b.texture
}

func (b *Block) Transparent() bool {
	return b.transparent
}

func (b *Block) VertexData(x, y, z float32, sf ShowFaces) []float32 {
	var data []float32

	var d = b.texture

	if sf[FaceUp] { // UP
		data = append(data,
			x-0.5, y+0.5, z+0.5, 0.0, 1.0, 0.0, d.Up[0][0], d.Up[0][1],
			x+0.5, y+0.5, z+0.5, 0.0, 1.0, 0.0, d.Up[1][0], d.Up[1][1],
			x+0.5, y+0.5, z-0.5, 0.0, 1.0, 0.0, d.Up[2][0], d.Up[2][1],
			x+0.5, y+0.5, z-0.5, 0.0, 1.0, 0.0, d.Up[3][0], d.Up[3][1],
			x-0.5, y+0.5, z-0.5, 0.0, 1.0, 0.0, d.Up[4][0], d.Up[4][1],
			x-0.5, y+0.5, z+0.5, 0.0, 1.0, 0.0, d.Up[5][0], d.Up[5][1],
		)
	}

	// Down
	if sf[FaceDown] { // DOWN
		data = append(data,
			x-0.5, y-0.5, z-0.5, 0.0, -1.0, 0.0, d.Down[0][0], d.Down[0][1],
			x+0.5, y-0.5, z-0.5, 0.0, -1.0, 0.0, d.Down[1][0], d.Down[1][1],
			x+0.5, y-0.5, z+0.5, 0.0, -1.0, 0.0, d.Down[2][0], d.Down[2][1],
			x+0.5, y-0.5, z+0.5, 0.0, -1.0, 0.0, d.Down[3][0], d.Down[3][1],
			x-0.5, y-0.5, z+0.5, 0.0, -1.0, 0.0, d.Down[4][0], d.Down[4][1],
			x-0.5, y-0.5, z-0.5, 0.0, -1.0, 0.0, d.Down[5][0], d.Down[5][1],
		)
	}

	if sf[FaceLeft] { // LEFT
		data = append(data,
			x-0.5, y-0.5, z-0.5, -1.0, 0.0, 0.0, d.Left[0][0], d.Left[0][1],
			x-0.5, y-0.5, z+0.5, -1.0, 0.0, 0.0, d.Left[1][0], d.Left[1][1],
			x-0.5, y+0.5, z+0.5, -1.0, 0.0, 0.0, d.Left[2][0], d.Left[2][1],
			x-0.5, y+0.5, z+0.5, -1.0, 0.0, 0.0, d.Left[3][0], d.Left[3][1],
			x-0.5, y+0.5, z-0.5, -1.0, 0.0, 0.0, d.Left[4][0], d.Left[4][1],
			x-0.5, y-0.5, z-0.5, -1.0, 0.0, 0.0, d.Left[5][0], d.Left[5][1],
		)
	}

	if sf[FaceRight] { // RIGHT
		data = append(data,
			x+0.5, y-0.5, z+0.5, 1.0, 0.0, 0.0, d.Right[0][0], d.Right[0][1],
			x+0.5, y-0.5, z-0.5, 1.0, 0.0, 0.0, d.Right[1][0], d.Right[1][1],
			x+0.5, y+0.5, z-0.5, 1.0, 0.0, 0.0, d.Right[2][0], d.Right[2][1],
			x+0.5, y+0.5, z-0.5, 1.0, 0.0, 0.0, d.Right[3][0], d.Right[3][1],
			x+0.5, y+0.5, z+0.5, 1.0, 0.0, 0.0, d.Right[4][0], d.Right[4][1],
			x+0.5, y-0.5, z+0.5, 1.0, 0.0, 0.0, d.Right[5][0], d.Right[5][1],
		)
	}

	if sf[FaceFront] { // FRONT
		data = append(data,
			x-0.5, y-0.5, z+0.5, 0.0, 0.0, 1.0, d.Front[0][0], d.Front[0][1],
			x+0.5, y-0.5, z+0.5, 0.0, 0.0, 1.0, d.Front[1][0], d.Front[1][1],
			x+0.5, y+0.5, z+0.5, 0.0, 0.0, 1.0, d.Front[2][0], d.Front[2][1],
			x+0.5, y+0.5, z+0.5, 0.0, 0.0, 1.0, d.Front[3][0], d.Front[3][1],
			x-0.5, y+0.5, z+0.5, 0.0, 0.0, 1.0, d.Front[4][0], d.Front[4][1],
			x-0.5, y-0.5, z+0.5, 0.0, 0.0, 1.0, d.Front[5][0], d.Front[5][1],
		)
	}

	if sf[FaceBack] { // BACK
		data = append(data,
			x+0.5, y-0.5, z-0.5, 0.0, 0.0, -1.0, d.Back[0][0], d.Back[0][1],
			x-0.5, y-0.5, z-0.5, 0.0, 0.0, -1.0, d.Back[1][0], d.Back[1][1],
			x-0.5, y+0.5, z-0.5, 0.0, 0.0, -1.0, d.Back[2][0], d.Back[2][1],
			x-0.5, y+0.5, z-0.5, 0.0, 0.0, -1.0, d.Back[3][0], d.Back[3][1],
			x+0.5, y+0.5, z-0.5, 0.0, 0.0, -1.0, d.Back[4][0], d.Back[4][1],
			x+0.5, y-0.5, z-0.5, 0.0, 0.0, -1.0, d.Back[5][0], d.Back[5][1],
		)
	}

	return data
}
