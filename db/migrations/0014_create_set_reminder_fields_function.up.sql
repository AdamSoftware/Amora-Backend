drop function if exists set_reminder_fields() cascade;

CREATE OR REPLACE FUNCTION set_reminder_fields()
RETURNS TRIGGER AS $$
DECLARE
    reservations_ts TIMESTAMP;
    reminder_ts TIMESTAMP;
BEGIN
    reservations_ts := NEW."ScheduleDate"::timestamp + NEW."ScheduleTime";
    reminder_ts := reservations_ts - (interval '1 minute' * NEW."ReminderTimeBefore");
    NEW."ReminderDate" := reminder_ts::date;
    NEW."ReminderTime" := reminder_ts::time;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


