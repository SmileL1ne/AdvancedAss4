package usecase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
)

type IContactUsecase interface {
	Create(contact.Contact) (int, error)  // returns id of created contact
	GetByID(int) (contact.Contact, error) // accepts contact id
	Update(contact.Contact) error
	Delete(int) error // accepts contact id

	CreateGroup(group.Group) (int, error) // returns id of created group
	DeleteGroup(int) error                // accepts group id
	UpdateGroup(group.Group) error
	GetGroupByID(int) (group.Group, error) // accepts group id
	InsertContactToGroup(int, int) error   // accepts contact id and group id
	DeleteContactFromGroup(int, int) error // accepts contact id and group id
}
