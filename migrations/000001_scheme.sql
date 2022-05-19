-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE links (
    short_url   varchar(255) not null,
    long_url varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE links
-- +goose StatementEnd
