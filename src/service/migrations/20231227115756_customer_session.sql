-- +goose Up
-- +goose StatementBegin
CREATE TABLE "customer_session" (
    "id"            BIGINT GENERATED ALWAYS AS IDENTITY,
    "customer_id"   BIGINT NOT NULL,
    "token"         VARCHAR(512) NOT NULL UNIQUE,
    "user_agent"    VARCHAR(1024) NOT NULL,
    "expires_at"    TIMESTAMP NOT NULL DEFAULT NOW() + INTERVAL '1 month',
    "created_at"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at"    TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY ("id"),
    FOREIGN KEY ("customer_id") REFERENCES "customer" ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "customer_session";
-- +goose StatementEnd