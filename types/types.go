package types

import "database/sql"

type TreeInfoPayload struct {
	Ppage     int `json:"Ppage"`
	PpageSize int `json:"PpageSize"`
}

type TreeInfoByIdPayload struct {
	PtreeID int `json:"PtreeID"`
}

type TreeInfoResponse struct {
	TreeID         int            `json:"tree_id"`
	CommonName     string         `json:"common_name"`
	ScientificName string         `json:"scientific_name"`
	TrunkDiameter  string         `json:"trunk_diameter"`
	CanopyWidth    float64        `json:"canopy_width"`
	Height         float64        `json:"height"`
	Age            int            `json:"age"`
	TreeCondition  string         `json:"tree_condition"`
	Notes          sql.NullString `json:"notes"`
}

type PagedTreeResponse struct {
	Trees      []*TreeInfoResponse `json:"trees"`
	Page       int                 `json:"page"`
	TotalPages int                 `json:"totalPages"`
}

type TreeInfoStore interface {
	GetTreeInfo(Ppage, PpageSize int) ([]*PagedTreeResponse, error)
	GetTreeInfoLocation(Ppage, PpageSize int) ([]*TreeInfoLocationResponse, error)
	GetTreeInfoByID(PtreeID int) (*TreeInfoById, error)
}

type TreeInfoLocation struct {
	TreeID         int     `json:"tree_id"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	CommonName     string  `json:"common_name"`
	ScientificName string  `json:"scientific_name"`
}

type TreeInfoLocationResponse struct {
	Location   []*TreeInfoLocation `json:"location"`
	Page       int                 `json:"page"`
	TotalPages int                 `json:"totalPages"`
}

type TreeInfoById struct {
	TreeID         int     `json:"tree_id"`
	CommonName     string  `json:"common_name"`
	ScientificName string  `json:"scientific_name"`
	TrunkDiameter  string  `json:"trunk_diameter"`
	CanopyWidth    float64 `json:"canopy_width"`
	Height         float64 `json:"height"`
	Age            int     `json:"age"`
	TreeCondition  string  `json:"tree_condition"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}
