/*
	Copyright (C) 2021-2022  The YNOproject Developers

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package server

type Minigame struct {
	Id             string `json:"minigameId"` // TODO: make this `json:"id"`
	VarId          int    `json:"varId"`
	InitialVarSync bool   `json:"initialVarSync"`
	SwitchId       int    `json:"switchId"`
	SwitchValue    bool   `json:"switchValue"`
	Dev            bool   `json:"dev"`
}

func getRoomMinigames(roomId int) (minigames []*Minigame) {
	switch serverConfig.GameName {
	case "yume":
		if roomId == 155 {
			minigames = append(minigames, &Minigame{Id: "nasu", VarId: 88, SwitchId: 215})
		}
	case "2kki":
		switch roomId {
		case 102:
			minigames = append(minigames, &Minigame{Id: "rby", VarId: 1010, InitialVarSync: true})
		case 618:
			minigames = append(minigames, &Minigame{Id: "rby_ex", VarId: 79, InitialVarSync: true})
		case 344:
			minigames = append(minigames, &Minigame{Id: "fuji_ex", VarId: 3218, SwitchId: 3219, SwitchValue: true})
		}
	}
	return minigames
}
