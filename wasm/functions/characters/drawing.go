package characters

import "github.com/wootra/vampire-survivors-clone/wasm/types"

func DrawCharacter(data *types.Data) {
	scale:= data.Canvas.GetObjectScale()
	charSize:=data.Character.CharInfo.Detail.Size * scale
	data.Canvas.Save()
	data.Canvas.Context.Translate(int(data.Character.CharInfo.PosX*scale), int(data.Character.CharInfo.PosY*scale))
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
	data.Canvas.Context.StrokeRect(data.Character.CharInfo.PosX*scale-charSize/2, data.Character.CharInfo.PosY*scale-charSize/2, charSize, charSize, 0, 255, 0, 255)
	data.Canvas.Context.StrokeCircle(data.Character.CharInfo.PosX*scale, data.Character.CharInfo.PosY*scale, charSize/2, 0, 255, 0, 255)

}

func DrawEnemy(data *types.Data, enemy *types.EnemyData) {
	scale:= data.Canvas.GetObjectScale()
	charSize:=data.Character.CharInfo.Detail.Size * scale
	data.Canvas.Save()

	data.Canvas.Context.Translate(int(enemy.CharInfo.PosX*scale), int(enemy.CharInfo.PosY*scale))

	data.Canvas.CanvasFuncs.Call("getCharacterImage", "cat", enemy.CharInfo.FrameIndex, -charSize/2, -charSize/2, charSize, charSize)
	data.Canvas.Context.Scale(enemy.Direction, 1)
	
	data.Canvas.Restore()
	
	data.Canvas.Context.StrokeRect(enemy.CharInfo.PosX*scale-charSize/2, enemy.CharInfo.PosY*scale-charSize/2, charSize, charSize, 255, 0, 0, 255)
	data.Canvas.Context.StrokeCircle(enemy.CharInfo.PosX*scale, enemy.CharInfo.PosY*scale, charSize/2, 255, 0, 0, 255)
}
