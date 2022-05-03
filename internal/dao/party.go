package dao

import (
	"context"
	"github.com/KawashiroNitori/lootbot/ent"
	"github.com/KawashiroNitori/lootbot/ent/party"
	"github.com/KawashiroNitori/lootbot/internal/resource"
)

type PartyDAO interface {
	CreateParty(ctx context.Context, channelID string) (*ent.Party, error)
	GetParty(ctx context.Context, partyID string) *ent.Party
	GetPartyByChannelID(ctx context.Context, channelID string) *ent.Party
}

type PartyDAOImpl struct {
	client *ent.Client
}

func NewPartyDAO() *PartyDAOImpl {
	return &PartyDAOImpl{
		client: resource.DBClient,
	}
}

var (
	DefaultPartyDAO          = NewPartyDAO()
	_               PartyDAO = (*PartyDAOImpl)(nil)
)

func (p *PartyDAOImpl) CreateParty(ctx context.Context, channelID string) (*ent.Party, error) {
	pt, err := p.client.Party.
		Create().
		SetChannelID(channelID).
		Save(ctx)
	return pt, err
}

func (p *PartyDAOImpl) GetParty(ctx context.Context, partyID string) *ent.Party {
	pt := p.client.Party.
		Query().
		Where(party.ID(partyID)).
		FirstX(ctx)
	return pt
}

func (p *PartyDAOImpl) GetPartyByChannelID(ctx context.Context, channelID string) *ent.Party {
	pt := p.client.Party.
		Query().
		Where(party.ChannelID(channelID)).
		FirstX(ctx)
	return pt
}
