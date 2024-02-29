CREATE TABLE IF NOT EXISTS contacts (
    id INT SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255) NOT NULL,
    phone VARCHAR(30) NOT NULL UNIQUE
)

CREATE TABLE IF NOT EXISTS groups (
    id INT SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
)

CREATE TABLE IF NOT EXISTS contact_group (
    contact_id INT  REFERENCES contacts(id),
    group_id INT REFERENCES groups(id),
    PRIMARY KEY(contact_id, group_id)
)