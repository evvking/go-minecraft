package world

import (
	"fmt"
	"math/rand"

	"minecraft/assets"
)

type World struct {
	Name string
	Map  [][][]*assets.Block
}

func (world World) Make() {
	world.Map = make([][][]*assets.Block, 32)

	for i := range world.Map {
		world.Map[i] = make([][]*assets.Block, 32)

		for j := range world.Map[i] {
			world.Map[i][j] = make([]*assets.Block, 32)

			for k := range world.Map[i][j] {
				var blocksCount = len(assets.BlockList_name)
				var itemID = rand.Intn(blocksCount)

				world.Map[i][j][k] = &assets.Block{
					Name:         assets.GetBlockName(itemID),
					Icon:         assets.GetBlockIcon(itemID),
					BreakingTime: assets.GetBlockBreakingTime(itemID),
				}
			}
		}
	}
}

func (world World) Print() {
	for i := range world.Map {
		for j := range world.Map[i] {
			for k := range world.Map[i][j] {
				fmt.Println(world.Map[i][j][k].Icon)
			}
		}
	}
}
