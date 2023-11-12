CREATE TABLE tasks
(
    id         UUID      NOT NULL UNIQUE PRIMARY KEY,
    value      TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL,
    status     INT CHECK ( 0 < status and status < 5 )
);
---- create above / drop below ----
DROP TABLE tasks;