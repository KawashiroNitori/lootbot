package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/KawashiroNitori/lootbot/internal/dao"
	"github.com/KawashiroNitori/lootbot/internal/macro"
	"github.com/KawashiroNitori/lootbot/internal/model"
	"github.com/samber/lo"
	"os"
	"strings"
)

func ImportLootsFromCSV(ctx context.Context, csvPath string, partyID string, needClear bool) {
	pt := dao.DefaultPartyDAO.GetParty(ctx, partyID)
	if pt == nil {
		panic(fmt.Errorf("party not found: %s", partyID))
	}
	f := lo.Must(os.Open(csvPath))
	r := csv.NewReader(f)
	data := lo.Must(r.ReadAll())
	data = data[1:]
	if needClear {
		lo.Must0(dao.DefaultLootDAO.ClearPartyLoots(ctx, partyID))
	}
	for _, row := range data {
		playerName := row[0]
		playerServer := row[1]
		role := lo.Must(macro.RoleString(row[2]))
		job := lo.Must(macro.JobNameString(row[3]))

		for _, cell := range row[4:] {
			fmt.Printf("Processing %s %s -> %s\n", role.String(), playerName, cell)
			itemNames := strings.Split(cell, ",")
			for _, itemName := range itemNames {
				item := SearchXIVItemByName(ctx, itemName)
				if item == nil {
					panic(fmt.Sprintf("item not found: %s", itemName))
				}
				category := DetectCategoryByItem(item)
				lo.Must(dao.DefaultLootDAO.CreateLoot(
					ctx, partyID, playerName, playerServer, role, job,
					category, item.ID, item.Name))
			}
		}
	}
}

func DetectCategoryByItem(item *model.Item) macro.Category {
	switch {
	case strings.Contains(item.Name, "武器箱") || (item.ItemKind != nil && item.ItemKind.Name == "武器") || item.MateriaSlotCount > 0:
		return macro.CategoryWeapon
	case strings.Contains(item.Name, "装备箱"):
		return macro.CategoryCoffer
	case strings.Contains(item.Name, "硬化药"):
		return macro.CategoryCoating
	case strings.Contains(item.Name, "神典石"):
		return macro.CategoryTomestone
	case strings.Contains(item.Name, "强化药"):
		return macro.CategoryRoborant
	case strings.Contains(item.Name, "纤维"):
		return macro.CategorySpool
	case strings.Contains(item.Description, "获得新坐骑"):
		return macro.CategoryMount
	case strings.Contains(item.Name, "管弦乐琴乐谱："):
		return macro.CategoryOrchestraRoll
	case strings.Contains(item.Description, "获得新宠物"):
		return macro.CategoryCompanion
	}
	panic(fmt.Errorf("unknown item: %s", item.Name))
}
