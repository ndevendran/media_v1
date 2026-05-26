-- +goose Up
CREATE TABLE videos(
	id UUID Primary Key,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    duration_seconds INTEGER NOT NULL
);

-- +goose Down
DROP TABLE videos;
