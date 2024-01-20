-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Actor (
    idActor BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    creationDate DATETIME NOT NULL,
    lastUpdateDate DATETIME DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Actor;
-- +goose StatementEnd
