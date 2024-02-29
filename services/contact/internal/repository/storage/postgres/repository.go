package postgres

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/repository"
	"database/sql"
)

type PgRepository struct {
	db *sql.DB
}

func New(db *sql.DB) repository.IContactRepository {
	return &PgRepository{
		db: db,
	}
}

func (r *PgRepository) Insert(contact contact.Contact) (int, error) {
	query := `
		INSERT INTO contacts (name, surname, patronymic, phone)
		VALUES ($1, $2, $3, $4)
	`

	res, err := r.db.Exec(query, contact.Name, contact.Surname, contact.Patronymic, contact.Phone)
	if err != nil {
		return 0, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (r *PgRepository) GetByID(id int) (contact.Contact, error) {
	query := `
		SELECT *
		FROM contacts
		WHERE id = $1
	`

	var c contact.Contact

	if err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Surname, &c.Patronymic, &c.Phone); err != nil {
		return contact.Contact{}, err
	}

	return c, nil
}

func (r *PgRepository) Update(contact contact.Contact) error {
	query := `
		UPDATE contacts
		SET name = $1, surname = $2, patronymic = $3, phone = $4
		WHRERE id = $5
	`

	_, err := r.db.Exec(query, contact.Name, contact.Surname, contact.Patronymic, contact.Phone, contact.ID)

	return err
}

func (r *PgRepository) Delete(id int) error {
	query := `
		DELETE 
		FROM contacts
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)

	return err
}

func (r *PgRepository) InsertGroup(group group.Group) (int, error) {
	query := `
		INSERT INTO groups (name)
		VALUES ($1)
	`

	res, err := r.db.Exec(query, group.Name)
	if err != nil {
		return 0, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (r *PgRepository) DeleteGroup(id int) error {
	query := `
		DELETE
		FROM groups
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id)

	return err
}

func (r *PgRepository) UpdateGroup(group group.Group) error {
	query := `
		UPDATE groups
		SET name = $1
		WHRERE id = $2
	`

	_, err := r.db.Exec(query, group.Name, group.ID)

	return err
}

func (r *PgRepository) GetGroupByID(id int) (group.Group, error) {
	query := `
		SELECT *
		FROM groups
		WHERE id = $1
	`
	var g group.Group

	if err := r.db.QueryRow(query, id).Scan(&g.ID, &g.Name); err != nil {
		return group.Group{}, err
	}

	return g, nil
}

func (r *PgRepository) InsertContactToGroup(contactID int, groupID int) error {
	query := `
		INSERT INTO contact_group (contact_id, group_id)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(query, contactID, groupID)

	return err
}

func (r *PgRepository) DeleteContactFromGroup(contactID int, groupID int) error {
	query := `
		DELETE
		FROM contact_group
		WHERE contact_id = $1 AND group_id = $2
	`

	_, err := r.db.Exec(query, contactID, groupID)

	return err
}
