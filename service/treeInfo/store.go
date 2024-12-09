package treeinfo

import (
	"database/sql"

	"github.com/CloudOpsOperation/COO-API-ECOTEC/types"
)

type Store struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetTreeInfo(Ppage, PpageSize int) ([]*types.PagedTreeResponse, error) {
	println("Ppage: ", Ppage)
	rows, err := s.db.Query("CALL GetTreeData(?, ?)", Ppage, PpageSize)
	if err != nil {
		return nil, err
	}

	u := new(types.TreeInfoResponse)
	var trees []*types.TreeInfoResponse
	for rows.Next() {
		u, err = scanRowTreeInfo(rows)
		if err != nil {
			return nil, err
		}
		trees = append(trees, u)
	}

	row := s.db.QueryRow("CALL GetTreeTotalPages(?)", PpageSize)
	var totalPages int
	if err := row.Scan(&totalPages); err != nil {
		return nil, err
	}

	response := &types.PagedTreeResponse{
		Trees:      trees,
		Page:       Ppage,
		TotalPages: totalPages,
	}

	return []*types.PagedTreeResponse{response}, nil
}

func scanRowTreeInfo(rows *sql.Rows) (*types.TreeInfoResponse, error) {
	u := new(types.TreeInfoResponse)

	err := rows.Scan(
		&u.TreeID,
		&u.CommonName,
		&u.ScientificName,
		&u.TrunkDiameter,
		&u.CanopyWidth,
		&u.Height,
		&u.Age,
		&u.TreeCondition,
		&u.Notes,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Store) GetTreeInfoLocation(Ppage, PLimit int) ([]*types.TreeInfoLocationResponse, error) {
	rows, err := s.db.Query("CALL GetTreeDataLocation(?, ?)", Ppage, PLimit)
	if err != nil {
		return nil, err
	}

	u := new(types.TreeInfoLocation)
	var location []*types.TreeInfoLocation
	for rows.Next() {
		u, err = scanRowTreeInfoLocation(rows)
		if err != nil {
			return nil, err
		}
		location = append(location, u)
	}

	row := s.db.QueryRow("CALL GetTreeLocationTotalPages(?)", PLimit)
	var totalPages int
	if err := row.Scan(&totalPages); err != nil {
		return nil, err
	}
	response := &types.TreeInfoLocationResponse{
		Location:   location,
		Page:       Ppage,
		TotalPages: totalPages,
	}

	return []*types.TreeInfoLocationResponse{response}, nil
}

func (s *Store) GetTreeInfoLocationAll() ([]*types.TreeInfoLocation, error) {
	rows, err := s.db.Query("CALL GetTreeDataLocationAll()")
	if err != nil {
		return nil, err
	}

	u := new(types.TreeInfoLocation)
	var location []*types.TreeInfoLocation
	for rows.Next() {
		u, err = scanRowTreeInfoLocation(rows)
		if err != nil {
			return nil, err
		}
		location = append(location, u)
	}

	return location, nil
}

func scanRowTreeInfoLocation(rows *sql.Rows) (*types.TreeInfoLocation, error) {
	u := new(types.TreeInfoLocation)

	err := rows.Scan(
		&u.TreeID,
		&u.Latitude,
		&u.Longitude,
		&u.CommonName,
		&u.ScientificName,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Store) GetTreeInfoByID(treeID int) (*types.TreeInfoById, error) {
	row := s.db.QueryRow("CALL GetTreeDataById(?)", treeID)
	u := new(types.TreeInfoById)
	err := row.Scan(
		&u.TreeID,
		&u.Latitude,
		&u.Longitude,
		&u.CommonName,
		&u.ScientificName,
		&u.TrunkDiameter,
		&u.CanopyWidth,
		&u.Height,
		&u.Age,
		&u.TreeCondition,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}
