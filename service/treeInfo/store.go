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

func (s *Store) GetTreeInfo(PLimit int) ([]*types.TreeInfoResponse, error) {
	rows, err := s.db.Query("CALL GetTreeData(?)", PLimit)
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

	return trees, nil
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

func (s *Store) GetTreeInfoLocation(PLimit int) ([]*types.TreeInfoLocation, error) {
	rows, err := s.db.Query("CALL GetTreeDataLocation(?)", PLimit)
	if err != nil {
		return nil, err
	}

	u := new(types.TreeInfoLocation)
	var trees []*types.TreeInfoLocation
	for rows.Next() {
		u, err = scanRowTreeInfoLocation(rows)
		if err != nil {
			return nil, err
		}
		trees = append(trees, u)
	}

	return trees, nil
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
