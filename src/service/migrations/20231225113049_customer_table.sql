-- +goose Up
-- +goose StatementBegin
CREATE TABLE "customer" (
    "id"            BIGINT GENERATED ALWAYS AS IDENTITY,
    "username"      VARCHAR(255),
    "email"         VARCHAR(360) NOT NULL UNIQUE,
    "password_hash" VARCHAR(512) NOT NULL,
    "is_confirmed"  BOOLEAN NOT NULL DEFAULT false,
    "created_at"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at"    TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY ("id")
);

CREATE INDEX "customer_email_idx" ON "customer" ("email");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "customer_email_idx";
DROP TABLE "customer";
-- +goose StatementEnd
