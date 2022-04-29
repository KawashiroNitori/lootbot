package loot

import (
	"context"

	"github.com/KawashiroNitori/lootbot/internal/dao"
	"github.com/KawashiroNitori/lootbot/internal/model"
	"github.com/samber/lo"
)

type Dispatcher interface {
	GetRecommendedPlayersByLootItem(ctx context.Context, partyID, itemID int64) []*model.PlayerNeedInfo
	ObtainLootItem(ctx context.Context, partyID int64, playerName, playerServer string, itemID int64) error
}

type DispatcherImpl struct {
	lootDAO dao.LootDAO
}

func NewDispatcher() *DispatcherImpl {
	return &DispatcherImpl{
		lootDAO: dao.DefaultLootDAO,
	}
}

var (
	DefaultDispatcher            = NewDispatcher()
	_                 Dispatcher = (*DispatcherImpl)(nil)
)

func (d *DispatcherImpl) GetRecommendedPlayersByLootItem(ctx context.Context, partyID, itemID int64) []*model.PlayerNeedInfo {
	lootInfos := d.lootDAO.GetNeedInfoByItemID(ctx, partyID, itemID)
	if len(lootInfos) == 0 || len(lootInfos[0].Loots) == 0 {
		return nil
	}
	category := lootInfos[0].Loots[0].Category
	needInfos := lo.Map(lootInfos, func(lt *model.PlayerLootInfo, _ int) *model.PlayerNeedInfo {
		return lt.GetPlayerNeedInfo(category)
	})
	return needInfos
}

func (d *DispatcherImpl) ObtainLootItem(ctx context.Context, partyID int64, playerName, playerServer string, itemID int64) error {
	_, err := d.lootDAO.ObtainLoot(ctx, partyID, playerName, playerServer, itemID)
	return err
}
