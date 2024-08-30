CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    genre VARCHAR(255) NOT NULL,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    location_id INTEGER
);

CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    case_id INTEGER NOT NULL,
    shelf_name VARCHAR(255) NOT NULL,
    book_id INTEGER REFERENCES books(id) ON DELETE CASCADE
);


