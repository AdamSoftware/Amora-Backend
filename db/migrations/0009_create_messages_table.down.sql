drop trigger if exists trg_update_messages_updated_at on messages;
drop function if exists update_updated_at_column();
drop table if exists messages;
drop type if exists message_statuses;
