drop trigger if exists trg_set_reminder_fields on reminders;


CREATE TRIGGER trg_set_reminder_fields
BEFORE INSERT OR UPDATE ON reminders
FOR EACH ROW
EXECUTE FUNCTION set_reminder_fields();
