-- +goose Up
CREATE TABLE IF NOT EXISTS session(
    id TEXT NOT NULL PRIMARY KEY,
    refresh_token TEXT NOT NULL,
    access_token TEXT NOT NULL,
    user_id TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS user(
    id TEXT NOT NULL PRIMARY KEY,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    alias TEXT DEFAULT '',
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL DEFAULT '',
    birth_date TEXT NOT NULL DEFAULT '1000-01-01 00:00:00',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS chat(
    id TEXT NOT NULL PRIMARY KEY,
    chat_name TEXT NOT NULL,
    chat_icon TEXT NOT NULL DEFAULT '',
    chat_type INT NOT NULL
);

CREATE TABLE IF NOT EXISTS chat_messages(
    id TEXT NOT NULL PRIMARY KEY,
    message_type INT NOT NULL,
    message_text TEXT NOT NULL DEFAULT '',
    time TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    chat_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    FOREIGN KEY(chat_id) REFERENCES chat(id),
    FOREIGN KEY(user_id) REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS chat_members(
    chat_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    role INT NOT NULL DEFAULT '0',
    FOREIGN KEY(chat_id) REFERENCES chat(id),
    FOREIGN KEY(user_id) REFERENCES user(id)
);

-- +goose Down
DROP TABLE session;
DROP TABLE user;
DROP TABLE chat;
DROP TABLE chat_messages;
DROP TABLE chat_members;

