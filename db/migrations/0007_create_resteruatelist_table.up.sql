CREATE TABLE IF NOT EXISTS resteruatelist (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    address TEXT NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    phone_number TEXT,
    email TEXT,
    zip_code TEXT NOT NULL,
    image_url TEXT,
    website TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true
    -- NOTE:: I might use this in the future if I am able to get a price for different restaurants
    -- price NUMERIC(10, 2) NOT NULL
);

