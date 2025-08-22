-- Seed plans
INSERT INTO plans (plantype, price) VALUES
('Basic', 0.00),
('Gold', 19.99),
('Platnium', 29.99);

-- Seed users
-- Note: Passwords should be hashed before inserting into the database. 
-- user1: Alice
-- user2: Bob
-- user3: Carol
INSERT INTO users (email, password, FirstName, LastName, phone_number, address, city, state, zip_code, is_active) VALUES
('AdamSteinberg16@gmail.com', '$2y$12$TUbq9tTZrsV8z..g7yyNmu7hKNgDT0qbWk3fAwZzGbm6PdfDrcTPG', 'Alice', 'Smith', '9204721284', '123 Maple St', 'Springfield', 'IL', '62701', true),
('bob@example.com', '$2y$12$pLNvbSltpQFkXkQaiLgm1Oiwgvm0XjM.BMa8EYQVe8NQy6k8t.vmS', 'Bob', 'Jones', '555-2222', '456 Oak St', 'Springfield', 'IL', '62702', true),
('carol@example.com', '$2y$12$DEO.IwQTXHa/3kfxl.XGHeg5epQqaZgAXtzxuKMZZT6PzaOejBY5G', 'Carol', 'White', '555-3333', '789 Pine St', 'Springfield', 'IL', '62703', false);

-- Seed gifts
INSERT INTO gifts (name, description, image_url, price, is_active) VALUES
('Chocolate Box', 'Assorted fine chocolates', 'https://example.com/chocolate.jpg', 25.00, true),
('Flower Bouquet', 'Fresh seasonal flowers', 'https://example.com/flowers.jpg', 45.00, true),
('Wine Bottle', 'Premium red wine', 'https://example.com/wine.jpg', 55.00, true),
('Gift Card', 'Gift card for any occasion', 'https://example.com/giftcard.jpg', 30.00, false),
('Spa Voucher', 'Relaxing spa day voucher', 'https://example.com/spa.jpg', 100.00, true),
('Dinner Reservation', 'Reservation at a top restaurant', 'https://example.com/dinner.jpg', 75.00, false);

-- Seed events
INSERT INTO events (name, description, DateOfTheEvent, is_active, user_id) VALUES
('Spring Gala', 'Annual spring fundraising gala', '2025-04-25', true, 1),
('Summer Festival', 'Outdoor music and food festival', '2025-07-10', true, 1),
('Winter Ball', 'Formal winter dance event', '2025-12-15', true, 2);

-- Seed venueslist
INSERT INTO venueslist (name, address, city, state, zip_code, phone_number, email, website, image_url, is_active) VALUES
('Grand Hall', '100 Main St', 'Springfield', 'IL', '62701', '555-4444', 'info@grandhall.com', 'http://grandhall.com', 'https://example.com/venue1.jpg', true),
('City Conference Center', '200 Center Ave', 'Springfield', 'IL', '62702', '555-5555', 'contact@ccc.com', 'http://ccc.com', 'https://example.com/venue2.jpg', true);

-- Seed experienceslist
INSERT INTO experienceslist (name, description, address, city, state, zip_code, phone_number, price, image_url, is_active) VALUES
('Hot Air Balloon Ride', 'Experience the thrill of flight', '500 Balloon Rd', 'Springfield', 'IL', '62704', '555-6666', 150.00, 'https://example.com/balloon.jpg', true),
('Cooking Class', 'Learn gourmet recipes', '300 Chef St', 'Springfield', 'IL', '62705', '555-7777', 80.00, 'https://example.com/cooking.jpg', true);

-- Seed resteruatelist (restaurants)
INSERT INTO resteruatelist (name, description, address, city, state, zip_code, image_url, website, is_active) VALUES
('Mama’s Kitchen', 'Authentic Italian cuisine', '123 Pasta Blvd', 'Springfield', 'IL', '62706', 'https://example.com/mamas.jpg', 'http://mamaskitchen.com', true),
('Sushi World', 'Fresh sushi and sashimi', '456 Fish Rd', 'Springfield', 'IL', '62707', 'https://example.com/sushi.jpg', 'http://sushiworld.com', true);

-- Seed partners
INSERT INTO partners (firstname, lastname, nickname, birthdate, phone_number, user_id, anniversary) VALUES
('John', 'Doe', 'JD', '1980-08-19', '555-8888', 1, '2005-08-19'),
('Jane', 'Smith', NULL, '1990-08-19', '555-9999', 2, '2015-08-19');

-- Seed messages
INSERT INTO messages (partner_id, content, status, sent_at, created_at, updated_at) VALUES
(1, 'Welcome to our service!', 'sent', '2025-01-01 10:00:00+00', NOW(), NOW()),
(2, 'Your reservation is confirmed.', 'pending', NULL, NOW(), NOW());

-- Seed subscriptions
INSERT INTO subscriptions (user_id, plan_id, status, start_date, next_payment, cancelled_at, payed) VALUES
(1, 2, 'active', '2025-01-01 00:00:00+00', '2025-02-01 00:00:00+00', NULL, true),
(2, 3, 'cancelled', null, null, '2024-07-15 00:00:00+00', false),
(2, 1, 'past_due', '2024-06-15 00:00:00+00', '2024-07-15 00:00:00+00', '2024-07-01 00:00:00+00', true);


-- Seed attachedgifts
INSERT INTO attachedgifts (event_id, gift_id, quantity, is_active) VALUES
(1, 1, 10, true),
(1, 2, 5, true),
(2, 3, 3, true),
(1, 3, 2, false),
(2, 2, 5, true),
(3, 3, 2, true);

-- Seed reservations
INSERT INTO reservations (
    user_id, partner_id, event_id, restaurant_id, experience_id, venue_id, reservation_time, party_size, special_requests
) VALUES
(1, 1, 1, 1, NULL, 1, '2025-04-25 19:00:00', 4, 'Window seat please'),
(2, 2, 2, NULL, 1, 2, '2025-07-10 12:00:00', 2, NULL);

-- Seed reminders
INSERT INTO reminders (
    partner_id, reservation_id, "ScheduleDate", "ScheduleTime", "ReminderTimeBefore"
) VALUES
(1, 1, '2025-04-24', '17:00:00', 120),
(2, 2, '2025-07-09', '10:00:00', 60);

