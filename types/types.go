package types

import "database/sql"

type TreeInfoPayload struct {
	PLimit int `json:"plimit"`
}

type TreeInfoResponse struct {
	TreeID         int            `json:"tree_id"`
	CommonName     string         `json:"common_name"`
	ScientificName string         `json:"scientific_name"`
	TrunkDiameter  float64        `json:"trunk_diameter"`
	CanopyWidth    float64        `json:"canopy_width"`
	Height         float64        `json:"height"`
	Age            int            `json:"age"`
	TreeCondition  string         `json:"tree_condition"`
	Notes          sql.NullString `json:"notes"`
}

type TreeInfoStore interface {
	GetTreeInfo(PLimit int) ([]*TreeInfoResponse, error)
	GetTreeInfoLocation(PLimit int) ([]*TreeInfoLocation, error)
}

type TreeInfoLocation struct {
	TreeID         int     `json:"tree_id"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	CommonName     string  `json:"common_name"`
	ScientificName string  `json:"scientific_name"`
}
