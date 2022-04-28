package dao

import (
	"context"
)

type Loot struct {
	ID int64 `json:"id"`
}

type LootDAO interface {
	CreateLoot(ctx context.Context)
}
