package model

import (
	"github.com/KawashiroNitori/lootbot/ent"
	"github.com/KawashiroNitori/lootbot/internal/macro"
)

type Player struct {
	Name   string
	Server string
	Role   macro.Role
	Job    macro.Job
}

type PlayerLootInfo struct {
	Player       *Player
	Loots        []*ent.Loot
	CategoryLoot map[macro.Category][]*ent.Loot
}

func (p *PlayerLootInfo) GetPlayerNeedInfo(category macro.Category) *PlayerNeedInfo {
	var obtained int
	all := len(p.CategoryLoot[category])
	for _, loot := range p.CategoryLoot[category] {
		if loot.IsObtained {
			obtained++
		}
	}
	return &PlayerNeedInfo{
		Player:   p.Player,
		Category: category,
		Obtained: obtained,
		All:      all,
	}
}

type PlayerNeedInfo struct {
	Player   *Player
	Category macro.Category
	Obtained int
	All      int
}

func (p *PlayerNeedInfo) NeedScore() float64 {
	if p.All == 0 {
		return 0
	}
	left := p.All - p.Obtained
	return float64(left) / float64(p.All)
}
