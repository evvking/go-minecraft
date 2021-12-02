package assets

type Block struct {
	Name         string
	Icon         string
	BreakingTime int
}

var BlockList_name = []string{"Grass", "Dirt", "Stone", "Cobblestone"}
var BlockList_icon = []string{"[GR]", "[DI]", "[ST]", "[CS]"}
var BlockList_breakingTime = []int{1, 2, 3, 4}

func GetBlockName(id int) string {
	return BlockList_name[id]
}

func GetBlockIcon(id int) string {
	return BlockList_icon[id]
}

func GetBlockBreakingTime(id int) int {
	return BlockList_breakingTime[id]
}
