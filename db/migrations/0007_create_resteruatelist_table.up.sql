CREATE TABLE IF NOT EXISTS resteruatelist (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    address TEXT NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    zip_code TEXT NOT NULL,
    image_url TEXT,
    website TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true
);

