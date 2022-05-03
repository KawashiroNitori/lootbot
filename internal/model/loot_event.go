package model

import "github.com/KawashiroNitori/lootbot/internal/macro"

type LootEvent struct {
	Timestamp         int64               `json:"timestamp,omitempty"`
	LootMessage       *LootMessage        `json:"lootMessage,omitempty"`
	LootEventType     macro.LootEventType `json:"lootEventType,omitempty"`
	LootEventTypeName string              `json:"lootEventTypeName,omitempty"`
	IsLocalPlayer     bool                `json:"isLocalPlayer,omitempty"`
	PlayerName        string              `json:"playerName,omitempty"`
	World             string              `json:"world,omitempty"`
	Roll              int64               `json:"roll,omitempty"`
	TerritoryTypeID   int64               `json:"territoryTypeId,omitempty"`
	ContentID         int64               `json:"contentId,omitempty"`
	LootEventID       string              `json:"lootEventId,omitempty"`
	ItemName          string              `json:"itemName,omitempty"`
}

type LootMessage struct {
	XIVChatType         int64    `json:"xivChatType,omitempty"`
	LogKind             int64    `json:"logKind,omitempty"`
	LogKindName         string   `json:"logKindName,omitempty"`
	LootMessageType     int64    `json:"lootMessageType,omitempty"`
	LootMessageTypeName string   `json:"lootMessageTypeName,omitempty"`
	Message             string   `json:"message,omitempty"`
	MessageParts        []string `json:"messageParts,omitempty"`
	ItemID              int64    `json:"itemId,omitempty"`
	ItemName            string   `json:"itemName,omitempty"`
	IsHQ                bool     `json:"isHq,omitempty"`
}
