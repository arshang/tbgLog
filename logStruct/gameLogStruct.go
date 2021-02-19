package logStruct

import (
	"time"
)

type GameLog struct {
	ArcadeID   int
	GameType   int
	InsertCoin int
	WinCoin    int
	ChanceL    int
	ChanceR    int
	ChestAward int
	Time       int
}

func NewLog(arcadeID int, insertCoin int, winCoin int, chanceL int, chanceR int, chestAward int, gameType int) *GameLog {

	timeSecond := time.Now().Unix()
	log := &GameLog{
		ArcadeID:   arcadeID,
		GameType:   gameType,
		InsertCoin: insertCoin,
		WinCoin:    winCoin,
		ChanceL:    chanceL,
		ChanceR:    chanceR,
		ChestAward: chestAward,
		Time:       int(timeSecond),
	}

	return log
}
