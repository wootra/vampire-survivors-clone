package functions

import (
	"fmt"
	"sort"
	"strconv"
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	"github.com/wootra/vampire-survivors-clone/wasm/functions/characters"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewCanvas() *types.Canvas {
	return &types.Canvas{0, 0, false, nil, types.Context2D{}, types.BitmapImage{}}
}

func CheckIfImageLoaded(data *types.Data) bool {
	if data.GlueFunctions == nil {
		return false
	}
	total := data.GlueFunctions.Call("getLoadStatus").Get("total").Int()
	loaded := data.GlueFunctions.Call("getLoadStatus").Get("loaded").Int()
	if loaded != total {
		fmt.Println("Image is loading... (" + strconv.Itoa(loaded) + "/" + strconv.Itoa(total) + ")")
		return false
	}
	return true
}

type Enemies []*types.EnemyData

func (a Enemies) Len() int           { return len(a) }
func (a Enemies) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Enemies) Less(i, j int) bool { return a[i].PosY < a[j].PosY || a[i].PosX < a[j].PosX }

func DrawInCanvas(data *types.Data) {
	if data.Canvas.CanvasFuncs == nil {
		return
	}

	data.Canvas.CanvasFuncs.Call("getBackground", "back-1")

	xScale := float64(data.Canvas.Width) / 100
	yScale := float64(data.Canvas.Height) / 100
	var charSize float64 = types.CHAR_SIZE * xScale

	characters.DrawCharacter(data, xScale, yScale, charSize)
	enemies := []*types.EnemyData{}
	for _, en := range data.Enemies {
		enemies = append(enemies, en)
	}
	sort.Slice(enemies, func(i, j int) bool {
		return enemies[i].PosY < enemies[j].PosY || enemies[i].PosX < enemies[j].PosX
	})

	for _, enemy := range enemies {
		characters.DrawEnemy(data, enemy, xScale, yScale, charSize)
	}
}

func SetCanvas(data *types.Data, width, height int, funcs *js.Value) {

	data.Canvas.Width = width
	data.Canvas.Height = height
	data.Canvas.CanvasFuncs = funcs
	data.Canvas.Context = types.Context2D(data.Canvas.CanvasFuncs.Call("getContext"))
	data.Canvas.Init = true
	data.Canvas.CanvasFuncs.Call("getBackground", "back-1")
}
