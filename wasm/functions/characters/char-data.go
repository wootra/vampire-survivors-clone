package characters

import (
	"github.com/wootra/vampire-survivors-clone/wasm/functions/weapon"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewCharacterData() *types.CharacterData {
	return &types.CharacterData{
		PosX:         0,
		PosY:         0,
		Speed:        2,
		Shield:       0,
		FrameIndex:   0,
		FrameOffset:  0,
		Armor:        0,
		Life:         100,
		Weapon:       weapon.CreateAWeapon(types.GUN),
		MovementCode: types.MovementType{Up: false, Down: false, Left: false, Right: false}}
}

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
