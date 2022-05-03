package model

type SearchResp struct {
	Results []*Item
}

type Item struct {
	ID               int64     `json:"ID,omitempty"`
	Name             string    `json:"Name,omitempty"`
	Icon             string    `json:"Icon,omitempty"`
	IconHD           string    `json:"IconHD,omitempty"`
	Description      string    `json:"Description,omitempty"`
	ItemKind         *ItemKind `json:"ItemKind,omitempty"`
	MateriaSlotCount int64     `json:"MateriaSlotCount,omitempty"`
}

type ItemKind struct {
	ID   int64  `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
}
