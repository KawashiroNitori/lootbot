package bot

import (
	"context"
	"fmt"
	"github.com/KawashiroNitori/lootbot/internal/dao"
	"github.com/KawashiroNitori/lootbot/internal/service"
	"github.com/lonelyevil/khl"
	"github.com/phuslu/log"
)

type LootChatHandler struct {
	lootDAO  dao.LootDAO
	partyDAO dao.PartyDAO
}

func NewLootChatHandler() *LootChatHandler {
	return &LootChatHandler{
		lootDAO:  dao.DefaultLootDAO,
		partyDAO: dao.DefaultPartyDAO,
	}
}

func (h *LootChatHandler) MessageIn(msgCtx *khl.KmarkdownMessageContext) {
	ctx := context.Background()
	if msgCtx.Extra.Author.Bot {
		return
	}
	if msgCtx.Common.Content != "/loot" {
		return
	}
	pt := h.partyDAO.GetPartyByChannelID(ctx, msgCtx.Common.TargetID)
	if pt == nil {
		return
	}
	loots := h.lootDAO.GetPartyLoot(ctx, pt.ID)
	if len(loots) == 0 {
		return
	}
	msg := service.FormatLootInfosMessage(ctx, loots)
	resp, err := msgCtx.Session.MessageCreate(&khl.MessageCreate{
		MessageCreateBase: khl.MessageCreateBase{
			Type:     khl.MessageTypeCard,
			TargetID: msgCtx.Common.TargetID,
			Content:  msg.MustBuildMessage(),
		},
	})
	if err != nil {
		panic(err)
	}
	log.Info().Msg(fmt.Sprintf("msg send success: %s", resp.MsgID))
}
