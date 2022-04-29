package dao

import (
	"context"
	"sort"

	"github.com/KawashiroNitori/lootbot/ent"
	"github.com/KawashiroNitori/lootbot/ent/loot"
	"github.com/KawashiroNitori/lootbot/internal/macro"
	"github.com/KawashiroNitori/lootbot/internal/model"
	"github.com/KawashiroNitori/lootbot/internal/resource"
	"github.com/samber/lo"
)

type Loot struct {
	ID int64 `json:"id"`
}

type LootDAO interface {
	CreateLoot(ctx context.Context, partyID int64, playerName, playerServer string, role macro.Role, job macro.Job, category macro.Category, itemID int64, itemName string) (*ent.Loot, error)

	GetPlayerLoot(ctx context.Context, partyID int64, playerName, playerServer string) *model.PlayerLootInfo
	GetPartyLoot(ctx context.Context, partyID int64) []*model.PlayerLootInfo
	GetNeedInfoByItemID(ctx context.Context, partyID, itemID int64) []*model.PlayerLootInfo

	ObtainLoot(ctx context.Context, partyID int64, playerName, playerServer string, itemID int64) (*ent.Loot, error)
}

type LootDAOImpl struct {
	client *ent.Client
}

func NewLootDAO() *LootDAOImpl {
	return &LootDAOImpl{
		client: resource.DBClient,
	}
}

var (
	DefaultLootDAO         = NewLootDAO()
	_              LootDAO = (*LootDAOImpl)(nil)
)

func (d *LootDAOImpl) CreateLoot(ctx context.Context, partyID int64, playerName, playerServer string, role macro.Role, job macro.Job, category macro.Category, itemID int64, itemName string) (*ent.Loot, error) {
	loot, err := d.client.Loot.
		Create().
		SetPlayerName(playerName).
		SetPlayerServer(playerServer).
		SetPartyID(partyID).
		SetRole(role).
		SetJob(job).
		SetCategory(category).
		SetItemID(itemID).
		SetItemName(itemName).
		Save(ctx)
	return loot, err
}

func (d *LootDAOImpl) GetPlayerLoot(ctx context.Context, partyID int64, playerName, playerServer string) *model.PlayerLootInfo {
	loots := d.client.Loot.
		Query().
		Where(
			loot.PartyID(partyID),
			loot.PlayerName(playerName),
			loot.PlayerServer(playerServer),
		).
		AllX(ctx)
	if len(loots) == 0 {
		return nil
	}
	for _, ltInfo := range d.formatPlayerLootInfo(ctx, loots) {
		if ltInfo.Player.Name == playerName && ltInfo.Player.Server == playerServer {
			return ltInfo
		}
	}
	return nil
}

func (d *LootDAOImpl) GetPartyLoot(ctx context.Context, partyID int64) []*model.PlayerLootInfo {
	loots := d.client.Loot.
		Query().
		Where(loot.PartyID(partyID)).
		AllX(ctx)
	if len(loots) == 0 {
		return nil
	}
	lootInfoMap := d.formatPlayerLootInfo(ctx, loots)
	lootInfos := lo.Values(lootInfoMap)
	sort.Slice(lootInfos, func(i, j int) bool {
		return lootInfos[i].Player.Role < lootInfos[j].Player.Role
	})
	return lootInfos
}

func (d *LootDAOImpl) formatPlayerLootInfo(ctx context.Context, loots []*ent.Loot) map[model.Player]*model.PlayerLootInfo {
	m := make(map[model.Player]*model.PlayerLootInfo)
	for _, lt := range loots {
		player := model.Player{
			Name:   lt.PlayerName,
			Server: lt.PlayerServer,
			Role:   lt.Role,
			Job:    lt.Job,
		}
		if _, ok := m[player]; !ok {
			m[player] = &model.PlayerLootInfo{
				Player: &player,
			}
		}
		m[player].Loots = append(m[player].Loots, lt)
		if m[player].CategoryLoot == nil {
			m[player].CategoryLoot = make(map[macro.Category][]*ent.Loot)
		}
		m[player].CategoryLoot[lt.Category] = append(m[player].CategoryLoot[lt.Category], lt)
	}
	return m
}

func (d *LootDAOImpl) GetNeedInfoByItemID(ctx context.Context, partyID, itemID int64) []*model.PlayerLootInfo {
	loots := d.client.Loot.
		Query().
		Where(
			loot.PartyID(partyID),
			loot.ItemID(itemID),
		).
		AllX(ctx)

	lootInfoMap := d.formatPlayerLootInfo(ctx, loots)
	lootInfos := lo.Values(lootInfoMap)
	return lootInfos
}

func (d *LootDAOImpl) ObtainLoot(ctx context.Context, partyID int64, playerName, playerServer string, itemID int64) (*ent.Loot, error) {
	lt := d.client.Loot.
		Query().
		Where(
			loot.PartyID(partyID),
			loot.PlayerName(playerName),
			loot.PlayerServer(playerServer),
			loot.ItemID(itemID),
			loot.IsObtained(false),
		).
		FirstX(ctx)
	if lt == nil {
		return nil, nil
	}
	lt, err := lt.Update().
		SetIsObtained(true).
		Save(ctx)
	return lt, err
}
