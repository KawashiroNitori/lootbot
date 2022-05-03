package handler

import (
	"encoding/json"
	"github.com/KawashiroNitori/lootbot/internal/macro"
	"github.com/KawashiroNitori/lootbot/internal/model"
	"github.com/KawashiroNitori/lootbot/internal/service"
	"github.com/KawashiroNitori/lootbot/internal/util"
	"net/http"
)

type LootUploadHandler struct {
	service service.Dispatcher
}

func NewLootUploadHandler() *LootUploadHandler {
	return &LootUploadHandler{
		service: service.DefaultDispatcher,
	}
}

func (h *LootUploadHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	partyID := r.URL.Query().Get("party_id")

	var events []*model.LootEvent
	defer util.Close(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&events); err != nil {
		panic(err)
	}
	if len(events) == 0 || events[0] == nil {
		return
	}
	event := events[0]
	switch event.LootEventType {
	case macro.LootEventTypeAdd:
		h.service.SendRecommendedPlayersByLootItem(ctx, partyID, event.LootMessage.ItemID)
	case macro.LootEventTypeObtain:
		h.service.ObtainLootItem(ctx, partyID, event.PlayerName, event.World, event.LootMessage.ItemID)
	}
}
