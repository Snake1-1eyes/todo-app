CREATE TABLE users (
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);
CREATE TABLE todo_lists (
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);
CREATE TABLE users_lists (
    id SERIAL NOT NULL UNIQUE,
    user_id int REFERENCES users(id) on delete cascade NOT NULL,
    list_id int REFERENCES todo_lists(id) on delete cascade NOT NULL
);
CREATE TABLE todo_items (
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    done BOOLEAN NOT NULL default false
);
CREATE TABLE lists_items (
    id SERIAL NOT NULL UNIQUE,
    item_id int REFERENCES todo_items(id) on delete cascade NOT NULL,
    list_id int REFERENCES todo_lists(id) on delete cascade NOT NULL
);