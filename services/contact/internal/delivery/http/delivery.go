package http

import (
	"architecture_go/services/contact/internal/delivery"
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/usecase"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ContactDelivery struct {
	uc usecase.IContactUsecase
}

func NewContactDelivery(u usecase.IContactUsecase) delivery.IContactDelivery {
	return &ContactDelivery{
		uc: u,
	}
}

func (cd *ContactDelivery) CreateContact(w http.ResponseWriter, req *http.Request) {
	var con contact.Contact

	if err := readJSON(req, &con); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := cd.uc.Create(con)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data := map[string]int{
		"id": id,
	}
	writeJSON(w, data, http.StatusCreated)
}

func (cd *ContactDelivery) ViewContact(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	con, err := cd.uc.GetByID(id)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, con, http.StatusOK)
}

func (cd *ContactDelivery) UpdateContact(w http.ResponseWriter, req *http.Request) {
	idStr := strings.TrimPrefix(req.URL.Path, "/contacts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var con contact.Contact
	if err := readJSON(req, &con); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	con.ID = id

	err = cd.uc.Update(con)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cd *ContactDelivery) DeleteContact(w http.ResponseWriter, req *http.Request) {
	idStr := strings.TrimPrefix(req.URL.Path, "/contacts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = cd.uc.Delete(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cd *ContactDelivery) CreateGroup(w http.ResponseWriter, req *http.Request) {
	var grp group.Group
	if err := readJSON(req, &grp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := cd.uc.CreateGroup(grp)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data := map[string]int{
		"id": id,
	}
	writeJSON(w, data, http.StatusCreated)
}

func (cd *ContactDelivery) DeleteGroup(w http.ResponseWriter, req *http.Request) {
	idStr := strings.TrimPrefix(req.URL.Path, "/groups/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = cd.uc.DeleteGroup(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cd *ContactDelivery) UpdateGroup(w http.ResponseWriter, req *http.Request) {
	idStr := strings.TrimPrefix(req.URL.Path, "/groups/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var grp group.Group
	if err := readJSON(req, &grp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	grp.ID = id

	err = cd.uc.UpdateGroup(grp)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cd *ContactDelivery) GetGroupByID(w http.ResponseWriter, req *http.Request) {
	idStr := strings.TrimPrefix(req.URL.Path, "/groups/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	grp, err := cd.uc.GetGroupByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, grp, http.StatusOK)
}

func (cd *ContactDelivery) InsertContactToGroup(w http.ResponseWriter, req *http.Request) {
	var data map[string]int
	if err := readJSON(req, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contactID, ok := data["contact_id"]
	if !ok {
		http.Error(w, "Missing contact_id in request body", http.StatusBadRequest)
		return
	}

	groupID, ok := data["group_id"]
	if !ok {
		http.Error(w, "Missing group_id in request body", http.StatusBadRequest)
		return
	}

	err := cd.uc.InsertContactToGroup(contactID, groupID)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (cd *ContactDelivery) DeleteContactFromGroup(w http.ResponseWriter, req *http.Request) {
	var data map[string]int
	if err := readJSON(req, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contactID, ok := data["contact_id"]
	if !ok {
		http.Error(w, "Missing contact_id in request body", http.StatusBadRequest)
		return
	}

	groupID, ok := data["group_id"]
	if !ok {
		http.Error(w, "Missing group_id in request body", http.StatusBadRequest)
		return
	}

	err := cd.uc.DeleteContactFromGroup(contactID, groupID)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
