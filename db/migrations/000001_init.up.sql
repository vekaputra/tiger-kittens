CREATE TABLE IF NOT EXISTS users
(
    id UUID PRIMARY KEY,
    username VARCHAR(200) NOT NULL,
    email VARCHAR(200) NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tigers
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(200)             NOT NULL,
    date_of_birth DATE                     NOT NULL,
    last_seen     TIMESTAMP WITH TIME ZONE NOT NULL,
    last_lat      NUMERIC(12, 8)           NOT NULL,
    last_long     NUMERIC(12, 8)           NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tiger_sightings
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    UUID                     NOT NULL
        CONSTRAINT fk_tiger_sightings_users REFERENCES users,
    tiger_id   INTEGER                  NOT NULL
        CONSTRAINT fk_tiger_sightings_tigers REFERENCES tigers,
    image      TIMESTAMP WITH TIME ZONE NOT NULL,
    lat        NUMERIC(12, 8)           NOT NULL,
    long       NUMERIC(12, 8)           NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);