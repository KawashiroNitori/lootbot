//go:generate go run github.com/dmarkham/enumer -type LootEventType -trimprefix LootEventType -sql -output loot_event_type_string.go
package macro

type LootEventType int

const (
	LootEventTypeAdd LootEventType = iota
	LootEventTypeCast
	LootEventTypeCraft
	LootEventTypeDesynth
	LootEventTypeDiscard
	LootEventTypeGather
	LootEventTypeGreed
	LootEventTypeLots
	LootEventTypeNeed
	LootEventTypeObtain
	LootEventTypePurchase
	LootEventTypeSearch
	LootEventTypeSell
	LootEventTypeUse
)
