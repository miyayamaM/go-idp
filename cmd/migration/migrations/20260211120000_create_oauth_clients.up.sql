CREATE TABLE oauth_clients (
    id              BIGSERIAL PRIMARY KEY,
    client_id       UUID NOT NULL UNIQUE,
    client_secret   TEXT NOT NULL,
    redirect_uri    TEXT NOT NULL,

    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
