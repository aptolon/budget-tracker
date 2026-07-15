CREATE TABLE users (
    id              UUID        PRIMARY KEY,
    version         INTEGER     NOT NULL    DEFAULT 1,
    role            TEXT        NOT NULL    CHECK (role IN ('admin', 'user')),
    login           TEXT        NOT NULL    UNIQUE CHECK (char_length(login) BETWEEN 3 AND 32),
    password_hash   TEXT        NOT NULL 
);

CREATE TABLE categories (
    id      UUID    PRIMARY KEY,
    version INTEGER NOT NULL DEFAULT 1,
    user_id UUID    NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title   TEXT    NOT NULL CHECK (char_length(title) BETWEEN 3 AND 32),
UNIQUE (user_id, title)
);

CREATE TABLE subcategories (
    id          UUID    PRIMARY KEY,
    version     INTEGER NOT NULL DEFAULT 1,
    category_id UUID    NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    title       TEXT    NOT NULL CHECK (char_length(title) BETWEEN 3 AND 32),
    UNIQUE (category_id, title)
);

CREATE TABLE expenses (
    id              UUID            PRIMARY KEY,
    version         INTEGER         NOT NULL DEFAULT 1,
    subcategory_id  UUID            NOT NULL REFERENCES subcategories(id) ON DELETE CASCADE,
    title           TEXT            NOT NULL CHECK (char_length(title) BETWEEN 3 AND 32),
    description     TEXT            CHECK (char_length(description) <= 200),
    amount          NUMERIC(12,2)   NOT NULL CHECK (amount > 0),
    spent_at        TIMESTAMPTZ     NOT NULL
);