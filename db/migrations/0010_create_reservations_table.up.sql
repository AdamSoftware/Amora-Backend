create table if not exists reservations (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    partner_id INTEGER NOT NULL REFERENCES partners(id) ON DELETE CASCADE,
    event_id INTEGER NOT NULL REFERENCES events(id) ON DELETE CASCADE, -- not optional, the reservation will always be for an event,
    -- an event is anything that is happening is anything that is happeningis anything that is happening, an event is anything that is happening, an event is anything that is happening, an event is anything that is happening, an event is anything that is happening, an event is anything that is happening, an event is anything that is happening 
    restaurant_id INTEGER REFERENCES resteruatelist(id) ON DELETE CASCADE, -- optional, if the reservation is for a restaurant
    experience_id INTEGER REFERENCES experienceslist(id) ON DELETE CASCADE, -- optional, if the reservation is for an experience
    venue_id INTEGER REFERENCES venueslist(id) ON DELETE CASCADE, -- optional, if the reservation is for a venue
    reservation_time TIMESTAMP NOT NULL,
    party_size INTEGER NOT NULL CHECK (party_size > 0),
    special_requests TEXT
);
