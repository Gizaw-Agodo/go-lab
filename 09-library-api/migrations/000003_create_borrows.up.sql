CREATE TABLE borrows (
    id BIGSERIAL PRIMARY KEY,

    user_id BIGINT NOT NULL REFERENCES users(id),

    book_id BIGINT NOT NULL REFERENCES books(id),

    borrowed_at TIMESTAMP NOT NULL DEFAULT NOW(),

    returned_at TIMESTAMP
);