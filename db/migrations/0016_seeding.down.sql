DELETE FROM reminders WHERE partner_id IN (1, 2);

DELETE FROM reservations WHERE user_id IN (1, 2);

DELETE FROM attachedgifts WHERE event_id IN (1, 2);

DELETE FROM subscriptions WHERE user_id IN (1, 2);

DELETE FROM messages WHERE partner_id IN (1, 2);

DELETE FROM partners WHERE id IN (1, 2);

DELETE FROM resteruatelist WHERE name IN ('Mama’s Kitchen', 'Sushi World');

DELETE FROM experienceslist WHERE name IN ('Hot Air Balloon Ride', 'Cooking Class');

DELETE FROM venueslist WHERE name IN ('Grand Hall', 'City Conference Center');

DELETE FROM events WHERE name IN ('Spring Gala', 'Summer Festival', 'Winter Ball');

DELETE FROM gifts WHERE name IN ('Chocolate Box', 'Flower Bouquet', 'Wine Bottle');

DELETE FROM users WHERE email IN ('alice@example.com', 'bob@example.com', 'carol@example.com');

DELETE FROM plans WHERE plantype IN ('Basic', 'Standard', 'Premium');
