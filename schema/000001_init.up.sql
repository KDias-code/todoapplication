CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS todo_lists
(
    id SERIAL PRIMARY KEY,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE IF NOT EXISTS users_lists
(
    id SERIAL PRIMARY KEY,
    user_id integer not null,
    list_id integer not null,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS todo_items
(
    id SERIAL PRIMARY KEY,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);

CREATE TABLE IF NOT EXISTS lists_items
(
    id SERIAL PRIMARY KEY,
    item_id integer not null,
    list_id integer not null,
    FOREIGN KEY (item_id) REFERENCES todo_items(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(id) ON DELETE CASCADE
);
