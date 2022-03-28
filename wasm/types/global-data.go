package types

import "syscall/js"

//to fix the redline, refer .vscode/settings.json in this workspace

type Data struct {
	Canvas        *Canvas
	Character     *HeroData
	GlueFunctions *js.Value
	Enemies       map[uint64]*EnemyData
}

const WORLD_WIDTH = 10
const WORLD_HEIGHT = 10

const WORLD_SIZE = WORLD_WIDTH * WORLD_HEIGHT // 10x10

type World struct {
	Pt             [10]map[int]map[int][]uint64 //10 x grid of id arrays.
	Active         int                      // current active grid
	Apocalypse     bool                     //if true, end the game
	MainChannel    chan bool
	BufferToDelete chan int
}

func (w *World) AddObject(buffIndex, x,y int, id uint64){
	if w.Pt[buffIndex]==nil {
		w.Pt[buffIndex]=make(map[int]map[int][]uint64)
		
	}
	if (w.Pt[buffIndex][x]==nil) {
		w.Pt[buffIndex][x]=make(map[int][]uint64)
	}
	if (w.Pt[buffIndex][x][y]==nil) {
		w.Pt[buffIndex][x][y] = []uint64{}
	}
	w.Pt[buffIndex][x][y] = append(w.Pt[buffIndex][x][y], id)
}


func (w *World) GetObjects(buffIndex, x,y int) []uint64{
	if w.Pt[buffIndex]==nil || w.Pt[buffIndex][x]==nil || w.Pt[buffIndex][x][y]==nil{
		return nil
	}
	
	return w.Pt[buffIndex][x][y]
	
}


func (data *Data) AddAnEnemy(enemies ...*EnemyData) {
	for _, enemy := range enemies {
		data.Enemies[enemy.CharInfo.Id] = enemy
	}
}

func (data *Data) AddAnEnemies(enemies []*EnemyData) {
	for _, enemy := range enemies {
		data.Enemies[enemy.CharInfo.Id] = enemy
	}
}
