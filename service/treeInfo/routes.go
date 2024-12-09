package treeinfo

import (
	"fmt"
	"net/http"

	"github.com/CloudOpsOperation/COO-API-ECOTEC/types"
	"github.com/CloudOpsOperation/COO-API-ECOTEC/untils"
	"github.com/gorilla/mux"
)

type TreeInfo struct {
	store types.TreeInfoStore
}

func NewTreeInfo(store *Store) *TreeInfo {
	return &TreeInfo{
		store: store,
	}
}

func (t *TreeInfo) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/treeinfo", t.getTreeInfo).Methods("POST")
	router.HandleFunc("/treeinfo-location", t.getTreeInfoLocation).Methods("POST")
	router.HandleFunc("/treeinfo-location-all", t.getTreeInfoLocationAll).Methods("GET")
	router.HandleFunc("/treeinfobyid", t.getTreeInfoByID).Methods("POST")
	router.HandleFunc("/genarateqr", t.getQrCode).Methods("POST")
	fmt.Println("TreeInfo routes registered")
}

func (t *TreeInfo) getTreeInfo(w http.ResponseWriter, r *http.Request) {
	var payload types.TreeInfoPayload

	if err := untils.ParseJson(r, &payload); err != nil {
		untils.WriteError(w, http.StatusBadRequest, err)
	}

	data, err := t.store.GetTreeInfo(payload.Ppage, payload.PpageSize)

	if err != nil {
		untils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	untils.WriteJson(w, http.StatusOK, data)
}

func (t *TreeInfo) getTreeInfoLocation(w http.ResponseWriter, r *http.Request) {
	var payload types.TreeInfoPayload

	if err := untils.ParseJson(r, &payload); err != nil {
		untils.WriteError(w, http.StatusBadRequest, err)
	}

	data, err := t.store.GetTreeInfoLocation(payload.Ppage, payload.PpageSize)

	if err != nil {
		untils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	untils.WriteJson(w, http.StatusOK, data)
}

func (t *TreeInfo) getTreeInfoByID(w http.ResponseWriter, r *http.Request) {
	var payload types.TreeInfoByIdPayload

	if err := untils.ParseJson(r, &payload); err != nil {
		untils.WriteError(w, http.StatusBadRequest, err)
	}

	data, err := t.store.GetTreeInfoByID(payload.PtreeID)

	if err != nil {
		untils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	untils.WriteJson(w, http.StatusOK, data)
}

func (t *TreeInfo) getQrCode(w http.ResponseWriter, r *http.Request) {
	var payload types.TreeInfoByIdPayload

	if err := untils.ParseJson(r, &payload); err != nil {
		untils.WriteError(w, http.StatusBadRequest, err)
	}

	data, err := untils.GenerateQRCode(payload.PtreeID)

	if err != nil {
		untils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	untils.WriteImage(w, http.StatusOK, data)
}

func (t *TreeInfo) getTreeInfoLocationAll(w http.ResponseWriter, r *http.Request) {
	data, err := t.store.GetTreeInfoLocationAll()

	if err != nil {
		untils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	untils.WriteJson(w, http.StatusOK, data)
}
