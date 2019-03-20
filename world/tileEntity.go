package world

// Metadata connected to certain blocks in the world
// Example usages:
// * Inventory in chests
// * Keeping state in machines like furnaces
// * Switching textures on blocks with multiple colors (like wool)
type TileEntity interface {
	Meta() int32
	SetMeta(int32) int32
	Data() map[int32]interface{}
	GetData(int32) interface{}
	SetData(int32, interface{})
}

// Default implementation of TileEntity interface
// will keep no data, only meta id!
type DefaultTileEntity struct {
	metaid int32
}

func (me *DefaultTileEntity) SetMeta(meta int32) int32 {
	me.metaid = meta
	return meta
}

func NewTileEntity(meta int32) TileEntity  {
	return &DefaultTileEntity{meta}
}

func (me *DefaultTileEntity) Meta() int32 {
	return me.metaid
}

func (me *DefaultTileEntity) Data() map[int32]interface{} {
	return nil
}
func (me *DefaultTileEntity) GetData(int32) interface{} {
	return nil
}

func (me *DefaultTileEntity) SetData(int32, interface{}) {
	return
}