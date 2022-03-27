package characters

import "github.com/wootra/vampire-survivors-clone/wasm/types"

func DrawCharacter(data *types.Data, xScale, yScale, charSize float64) {

	data.Canvas.Save()
	data.Canvas.Context.Translate(int(data.Character.CharInfo.PosX*xScale), int(data.Character.CharInfo.PosY*yScale))
	if data.Character.LastMovement.Up || data.Character.MovementCode.Up {
		if data.Character.LastMovement.Left || data.Character.MovementCode.Left {
			data.Canvas.Context.Rotate(45)
		} else if data.Character.LastMovement.Right || data.Character.MovementCode.Right {
			data.Canvas.Context.Scale(-1, 1)
			data.Canvas.Context.Rotate(45)
		} else {
			data.Canvas.Context.Rotate(90)
		}
	} else if data.Character.LastMovement.Down || data.Character.MovementCode.Down {
		if data.Character.LastMovement.Left {
			data.Canvas.Context.Rotate(-45)
		} else if data.Character.LastMovement.Right || data.Character.MovementCode.Right {
			data.Canvas.Context.Scale(-1, 1)
			data.Canvas.Context.Rotate(-45)
		} else {
			data.Canvas.Context.Scale(-1, 1)
			data.Canvas.Context.Rotate(-90)
		}
	} else if data.Character.LastMovement.Right || data.Character.MovementCode.Right {
		data.Canvas.Context.Scale(-1, 1)
	} //left doesn't need transform

	data.Canvas.CanvasFuncs.Call("getCharacterImage", "fish", data.Character.CharInfo.FrameIndex, -charSize/2, -charSize/2, charSize, charSize)
	data.Canvas.Restore()

}

func DrawEnemy(data *types.Data, enemy *types.EnemyData, xScale, yScale, charSize float64) {

	data.Canvas.Save()

	data.Canvas.Context.Translate(int(enemy.CharInfo.PosX*xScale), int(enemy.CharInfo.PosY*yScale))

	data.Canvas.CanvasFuncs.Call("getCharacterImage", "cat", enemy.CharInfo.FrameIndex, -charSize/2, -charSize/2, charSize, charSize)
	data.Canvas.Context.Scale(float32(enemy.Direction), 1)
	data.Canvas.Restore()

}
