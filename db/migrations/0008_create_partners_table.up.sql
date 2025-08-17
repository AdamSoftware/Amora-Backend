CREATE TABLE IF NOT EXISTS partners (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    nickname TEXT,
    -- this is for the birthdate of the partner so that we can setup this event automaticlly for flowers and things
    birthdate DATE NOT NULL,
    -- this is for the anniversary of the partner wedding or relationship will be used to setup of the text messages
    anniversary DATE NOT NULL,
    phone_number TEXT NOT NULL,
    user_id INT UNIQUE REFERENCES users(id)
);

