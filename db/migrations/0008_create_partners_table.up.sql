CREATE TABLE IF NOT EXISTS partners (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    nickname TEXT,
    -- this is for the birthdate of the partner so that we can setup this event automaticlly for flowers and things
    birthdate DATE NOT NULL,
    phone_number TEXT NOT NULL
);

