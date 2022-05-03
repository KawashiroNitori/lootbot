package service

import (
	"context"
	"fmt"
	"github.com/phuslu/log"
	"sort"

	"github.com/KawashiroNitori/lootbot/internal/dao"
	"github.com/KawashiroNitori/lootbot/internal/model"
	"github.com/samber/lo"
)

type Dispatcher interface {
	SendRecommendedPlayersByLootItem(ctx context.Context, partyID string, itemID int64)
	ObtainLootItem(ctx context.Context, partyID, playerName, playerServer string, itemID int64)
}

type DispatcherImpl struct {
	lootDAO  dao.LootDAO
	partyDAO dao.PartyDAO
}

func NewDispatcher() *DispatcherImpl {
	return &DispatcherImpl{
		lootDAO:  dao.DefaultLootDAO,
		partyDAO: dao.DefaultPartyDAO,
	}
}

var (
	DefaultDispatcher            = NewDispatcher()
	_                 Dispatcher = (*DispatcherImpl)(nil)
)

func (d *DispatcherImpl) SendRecommendedPlayersByLootItem(ctx context.Context, partyID string, itemID int64) {
	pt := d.partyDAO.GetParty(ctx, partyID)
	if pt == nil {
		return
	}
	lootInfos := d.lootDAO.GetNeedInfoByItemID(ctx, partyID, itemID)
	if len(lootInfos) == 0 || len(lootInfos[0].Loots) == 0 {
		return
	}
	category := lootInfos[0].Loots[0].Category
	needInfos := lo.Map(lootInfos, func(lt *model.PlayerLootInfo, _ int) *model.PlayerNeedInfo {
		return lt.GetPlayerNeedInfo(category)
	})
	sort.Slice(needInfos, func(i, j int) bool {
		if almostEqual(needInfos[i].NeedScore(), needInfos[j].NeedScore()) {
			return needInfos[i].Player.Role < needInfos[j].Player.Role
		} else if needInfos[i].NeedScore() > needInfos[j].NeedScore() {
			return true
		} else {
			return false
		}
	})
	msg := FormatNeedCardMessage(ctx, itemID, needInfos)
	resp, err := CreateCardMessage(ctx, pt.ChannelID, msg)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("send message error: %s", err.Error()))
	} else {
		log.Info().Msg(fmt.Sprintf("send message success, msg_id: %s", resp.MsgID))
	}
}

func (d *DispatcherImpl) ObtainLootItem(ctx context.Context, partyID, playerName, playerServer string, itemID int64) {
	pt := d.partyDAO.GetParty(ctx, partyID)
	if pt == nil {
		return
	}
	lootInfo := d.lootDAO.GetPlayerLoot(ctx, partyID, playerName, playerServer)
	if lootInfo == nil {
		return
	}
	lt, err := d.lootDAO.ObtainLoot(ctx, partyID, playerName, playerServer, itemID)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("obtain loot error: %s", err.Error()))
	}
	if lt == nil {
		return
	}
	msg := FormatObtainMessage(ctx, lootInfo, itemID)
	resp, err := CreateCardMessage(ctx, pt.ChannelID, msg)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("send message error: %s", err.Error()))
	} else {
		log.Info().Msg(fmt.Sprintf("send message success, msg_id: %s", resp.MsgID))
	}
}
