
CREATE TABLE IF NOT EXISTS snippets (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    content text NOT NULL,
    created timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    expired timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1
);

CREATE INDEX idx_snippets_created ON snippets(created);