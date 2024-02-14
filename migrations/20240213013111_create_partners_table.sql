-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS cube;
CREATE EXTENSION IF NOT EXISTS earthdistance;

CREATE TABLE partners (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    flooring_materials TEXT[] NOT NULL,
    address_long FLOAT NOT NULL,
    address_lat FLOAT NOT NULL,
    operating_radius INT NOT NULL,
    rating INT NOT NULL CHECK (rating >= 0 AND rating <= 5),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS partners;
DROP EXTENSION IF EXISTS cube;
DROP EXTENSION IF EXISTS earthdistance;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
