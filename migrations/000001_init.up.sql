CREATE SCHEMA todoapp;

CREATE TABLE todoapp.users (
    id serial PRIMARY KEY,
    version bigint NOT NULL DEFAULT 1,
    full_name varchar(100) NOT NULL,
    phone_number varchar(15) NOT NULL UNIQUE
);

ALTER TABLE
    todoapp.users
ADD
    CONSTRAINT full_name_check CHECK (
        length(full_name) > 3
        AND length(full_name) <= 100
    );

ALTER TABLE
    todoapp.users
ADD
    CONSTRAINT phone_number_check CHECK (phone_number ~ '^\+?[1-9]\d{1,14}$');

CREATE TABLE todoapp.todos (
    id serial PRIMARY KEY,
    version bigint NOT NULL DEFAULT 1,
    title varchar(100) NOT NULL,
    description varchar(1000),
    completed boolean NOT NULL DEFAULT false,
    started_at timestamptz NOT NULL DEFAULT NOW(),
    completed_at timestamptz,
    user_id int NOT NULL REFERENCES todoapp.users(id) ON DELETE CASCADE
);

ALTER TABLE
    todoapp.todos
ADD
    CONSTRAINT title_check CHECK (
        length(title) > 0
        AND length(title) <= 100
    );

ALTER TABLE
    todoapp.todos
ADD
    CONSTRAINT completed_check CHECK (
        (
            completed = false
            AND completed_at IS NULL
        )
        OR (
            completed = TRUE
            AND completed_at IS NOT NULL
            AND completed_at >= started_at
        )
    );