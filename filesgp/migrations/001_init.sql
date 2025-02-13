CREATE TABLE files (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    filename TEXT NOT NULL,
    filetype TEXT NOT NULL,
    filesize BIGINT NOT NULL,
    storage_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
