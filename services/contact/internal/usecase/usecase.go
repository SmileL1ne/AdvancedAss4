package usecase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/repository"
)

type ContactUsecase struct {
	contactRepo repository.IContactRepository
}

func New(repo repository.IContactRepository) IContactUsecase {
	return &ContactUsecase{
		contactRepo: repo,
	}
}

func (cs *ContactUsecase) Create(contact contact.Contact) (int, error) {
	return cs.contactRepo.Insert(contact)
}

func (cs *ContactUsecase) GetByID(id int) (contact.Contact, error) {
	return cs.contactRepo.GetByID(id)
}

func (cs *ContactUsecase) Update(contact contact.Contact) error {
	return cs.contactRepo.Update(contact)
}

func (cs *ContactUsecase) Delete(id int) error {
	return cs.contactRepo.Delete(id)
}

func (cs *ContactUsecase) CreateGroup(group group.Group) (int, error) {
	return cs.contactRepo.InsertGroup(group)
}

func (cs *ContactUsecase) UpdateGroup(group group.Group) error {
	return cs.contactRepo.UpdateGroup(group)
}

func (cs *ContactUsecase) DeleteGroup(groupID int) error {
	return cs.contactRepo.DeleteGroup(groupID)
}

func (cs *ContactUsecase) GetGroupByID(id int) (group.Group, error) {
	return cs.contactRepo.GetGroupByID(id)
}

func (cs *ContactUsecase) InsertContactToGroup(contactID int, groupID int) error {
	return cs.contactRepo.InsertContactToGroup(contactID, groupID)
}

func (cs *ContactUsecase) DeleteContactFromGroup(contactID int, groupID int) error {
	return cs.contactRepo.DeleteContactFromGroup(contactID, groupID)
}
