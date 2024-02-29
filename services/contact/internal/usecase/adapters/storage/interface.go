package storage

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
)

type IContactAdapter interface {
	Insert(contact.Contact) (int, error)  // returns id of created contact
	GetByID(int) (contact.Contact, error) // accepts contact id
	Update(contact.Contact) error
	Delete(int) error // accepts contact id

	InsertGroup(group.Group) (int, error)  // returns id of created group
	GetGroupByID(int) (group.Group, error) // accepts group id
	InsertContactToGroup(int, int) error   // accepts contact id and group id
}
