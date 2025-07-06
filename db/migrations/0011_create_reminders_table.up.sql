create table if not exists reminders (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    partner_id INT NOT NULL REFERENCES partners(id) ON DELETE CASCADE,
    reservation_id INT NOT NULL REFERENCES reservations(id) ON DELETE CASCADE,
    ScheduleDate DATE NOT NULL,
    ScheduleTime TIME NOT NULL, 
    ReminderDate DATE,
    ReminderTime TIME,
    ReminderTimeBefore INT NOT NULL DEFAULT 120 -- in minutes
);



