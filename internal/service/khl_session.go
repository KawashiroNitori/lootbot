package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/KawashiroNitori/lootbot/internal/macro"
	"github.com/KawashiroNitori/lootbot/internal/model"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/lonelyevil/khl"
	"github.com/samber/lo"
	"math"
	"strings"
)

var (
	DefaultKHLSession *khl.Session
	DefaultHTTPClient = httpclient.NewClient()
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func FormatNeedKMarkdown(ctx context.Context, needInfos []*model.PlayerNeedInfo) string {
	if len(needInfos) == 0 {
		return "> 😅 啊咧？好像没有人需求这件装备呢～"
	}
	var md strings.Builder
	md.WriteString("> ")

	for _, needInfo := range needInfos {
		md.WriteString(fmt.Sprintf("%s `%s` `%s` %s: %s (%d/%d)\n",
			needInfo.NeedEmoji(),
			needInfo.Player.Role.String(),
			needInfo.Player.Job.Name(),
			needInfo.Player.Name,
			needInfo.Category.Name(),
			needInfo.Obtained,
			needInfo.All,
		))
	}
	return md.String()
}

func GetXIVItem(ctx context.Context, itemID int64) *model.Item {
	resp, err := DefaultHTTPClient.Get(fmt.Sprintf("https://cafemaker.wakingsands.com/Item/%d", itemID), nil)
	if err != nil {
		panic(err)
	}
	var item *model.Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		panic(err)
	}
	return item
}

func SearchXIVItemByName(ctx context.Context, name string) *model.Item {
	resp := lo.Must(DefaultHTTPClient.Get(fmt.Sprintf("https://cafemaker.wakingsands.com/search?indexes=Item&string=%s", name), nil))
	var data *model.SearchResp
	lo.Must0(json.NewDecoder(resp.Body).Decode(&data))
	if len(data.Results) == 0 {
		return nil
	}
	return GetXIVItem(ctx, data.Results[0].ID)
}

func FormatNeedCardMessage(ctx context.Context, itemID int64, needInfos []*model.PlayerNeedInfo) *khl.CardMessage {
	item := GetXIVItem(ctx, itemID)
	card := &khl.CardMessageCard{
		Size: "lg",
		Modules: []interface{}{
			khl.CardMessageSection{
				Mode: khl.CardMessageSectionModeRight,
				Text: &khl.CardMessageElementKMarkdown{
					Content: fmt.Sprintf("%s\n%s",
						fmt.Sprintf("掉落了「**%s**」", item.Name),
						FormatNeedKMarkdown(ctx, needInfos),
					),
				},
				Accessory: &khl.CardMessageElementImage{
					Size: "sm",
					Src:  fmt.Sprintf("https://cafemaker.wakingsands.com/%s", item.IconHD),
				},
			},
		},
	}
	msg := &khl.CardMessage{card}
	return msg
}

func FormatObtainMessage(ctx context.Context, playerLootInfo *model.PlayerLootInfo, itemID int64) *khl.CardMessage {
	item := GetXIVItem(ctx, itemID)
	card := &khl.CardMessageCard{
		Size: "lg",
		Modules: []interface{}{
			khl.CardMessageSection{
				Mode: khl.CardMessageSectionModeRight,
				Text: &khl.CardMessageElementKMarkdown{
					Content: fmt.Sprintf("`%s` `%s` %s 🌸 %s 获得了「**%s**」\n*%s*",
						playerLootInfo.Player.Role.String(),
						playerLootInfo.Player.Job.Name(),
						playerLootInfo.Player.Name,
						playerLootInfo.Player.Server,
						item.Name,
						item.Description,
					),
				},
				Accessory: &khl.CardMessageElementImage{
					Size: "sm",
					Src:  fmt.Sprintf("https://cafemaker.wakingsands.com/%s", item.IconHD),
				},
			},
		},
	}
	return &khl.CardMessage{card}
}

func FormatPlayerLootInfoCard(ctx context.Context, lootInfo *model.PlayerLootInfo) *khl.CardMessageCard {
	player := lootInfo.Player
	var lootBuilder strings.Builder
	for _, category := range macro.CategoryValues() {
		if len(lootInfo.CategoryLoot[category]) == 0 {
			continue
		}
		lootBuilder.WriteString(fmt.Sprintf("**%s**\n", category.Name()))
		for _, loot := range lootInfo.CategoryLoot[category] {
			if loot.IsObtained {
				lootBuilder.WriteString("✅")
			} else {
				lootBuilder.WriteString("✋")
			}
			lootBuilder.WriteString(fmt.Sprintf(" %s\n", loot.ItemName))
		}
		lootBuilder.WriteString("\n")
	}
	card := &khl.CardMessageCard{
		Modules: []interface{}{
			&khl.CardMessageSection{
				Text: &khl.CardMessageElementKMarkdown{
					Content: fmt.Sprintf("`%s` `%s` %s 🌸 %s", player.Role.String(), player.Job.Name(), player.Name, player.Server),
				},
			},
			&khl.CardMessageSection{
				Text: &khl.CardMessageParagraph{
					Cols: 1,
					Fields: []any{
						&khl.CardMessageElementKMarkdown{
							Content: lootBuilder.String()[:lootBuilder.Len()-1],
						},
					},
				},
			},
		},
	}
	return card
}

func FormatLootInfosMessage(ctx context.Context, lootInfos []*model.PlayerLootInfo) *khl.CardMessage {
	var modules []any
	for _, lootInfo := range lootInfos {
		card := FormatPlayerLootInfoCard(ctx, lootInfo)
		modules = append(modules, card.Modules...)
		modules = append(modules, &khl.CardMessageDivider{})
	}
	msg := &khl.CardMessage{&khl.CardMessageCard{
		Modules: modules[:len(modules)-1],
	}}
	return msg
}

func CreateCardMessage(ctx context.Context, channelID string, message *khl.CardMessage) (*khl.MessageResp, error) {
	msg := message.MustBuildMessage()
	fmt.Println(msg)
	return DefaultKHLSession.MessageCreate(&khl.MessageCreate{
		MessageCreateBase: khl.MessageCreateBase{
			Type:     khl.MessageTypeCard,
			TargetID: channelID,
			Content:  msg,
		},
	})
}
