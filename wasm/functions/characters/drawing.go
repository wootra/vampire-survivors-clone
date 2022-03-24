package characters

import "github.com/wootra/vampire-survivors-clone/wasm/types"

func DrawCharacter(data *types.Data, xScale, yScale, charSize float32) {

	data.Canvas.Save()
	data.Canvas.Context.Translate(int(data.Character.PosX*xScale), int(data.Character.PosY*yScale))
	if data.Character.MovementCode.Up {
		if data.Character.MovementCode.Left {
			data.Canvas.Context.Rotate(45)
		} else if data.Character.MovementCode.Right {
			data.Canvas.Context.Scale(-1, 1)
			data.Canvas.Context.Rotate(45)
		} else {
			data.Canvas.Context.Rotate(90)
		}
	} else if data.Character.MovementCode.Down {
		if data.Character.MovementCode.Left {
			data.Canvas.Context.Rotate(-45)
		} else if data.Character.MovementCode.Right {
			data.Canvas.Context.Scale(-1, 1)
			data.Canvas.Context.Rotate(-45)
		} else {
			data.Canvas.Context.Scale(-1, 1)
			data.Canvas.Context.Rotate(-90)
		}
	} else if data.Character.MovementCode.Right {
		data.Canvas.Context.Scale(-1, 1)
	} //left doesn't need transform

	// data.Canvas.GetContext().FillRect(data.Character.PosX*xScale-charSize/2, data.Character.PosY*yScale-charSize/2, charSize, charSize, 255, 0, 0, 255)
	data.Canvas.CanvasFuncs.Call("getCharacterImage", "fish", data.Character.FrameIndex, -charSize/2, -charSize/2, charSize, charSize)
	data.Canvas.Restore()

}

func DrawEnemy(data *types.Data, enemyNo int, xScale, yScale, charSize float32) {

	data.Canvas.Save()
	enemy := data.Enemies[enemyNo]

	data.Canvas.Context.Translate(int(enemy.PosX*xScale), int(enemy.PosY*yScale))
	data.Canvas.Context.Rotate(enemy.Direction)
	// if enemy.Direction > 90 && enemy.Direction < 270 {
	// 	data.Canvas.Context.Scale(-1, 1)
	// 	data.Canvas.Context.Rotate(enemy.Direction - 90)
	// } else {
	// 	data.Canvas.Context.Rotate(enemy.Direction)
	// }

	data.Canvas.CanvasFuncs.Call("getCharacterImage", "cat", enemy.FrameIndex, -charSize/2, -charSize/2, charSize, charSize)
	data.Canvas.Context.Rotate(-enemy.Direction)
	// data.Canvas.GetContext().FillRect(enemy.PosX*xScale-charSize/2, enemy.PosY*yScale-charSize/2, charSize, charSize, 255, 0, 0, 255)
	// data.Canvas.CanvasFuncs.Call("getCharacterImage", "fish", enemy.FrameIndex, -charSize/2, -charSize/2, charSize, charSize)
	data.Canvas.Restore()

}
