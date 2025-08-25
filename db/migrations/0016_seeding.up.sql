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
('Spa Voucher', 'Relaxing spa day voucher', 'https://example.com/spa.jpg', 100.00, true);

-- Seed events
INSERT INTO events (name, description, DateOfTheEvent, is_active, user_id) VALUES
('Spring Gala', 'Annual spring fundraising gala', '2025-04-25', true, 1),
('Summer Festival', 'Outdoor music and food festival', '2025-07-10', true, 1),
('Winter Ball', 'Formal winter dance event', '2025-12-15', true, 2);

-- Seed venueslist
INSERT INTO venueslist (name, address, city, state, zip_code, phone_number, email, website, image_url, is_active, price, description) VALUES
('Grand Hall', '100 Main St', 'Springfield', 'IL', '62701', '555-1234', 'FakeEmail@gmail.com', 'http://grandhall.com', 'https://example.com/grandhall.jpg', true, 500.00, 'Spacious venue for large events'),
('City Conference Center', '200 Center St', 'Springfield', 'IL', '62702', '555-5678', 'FakeFakeEmail@gmail.com', 'http://cityconference.com', 'https://example.com/conference.jpg', true, 300.00, 'Modern facilities for meetings and conferences'),
('Movie Theater', '150 Cinema Rd', 'Springfield', 'IL', '62703', '555-8765', 'Movie@gmail.com', 'http://movietheater.com', 'https://example.com/movietheater.jpg', true, 700.99, 'Perfect for private screenings and events'),
('Lakeside Pavilion', '300 Lake St', 'Springfield', 'IL', '62703', '555-9012', 'FakeFake21@gmail.com', 'http://lakesidepavilion.com', 'https://example.com/lakeside.jpg', false, 400.00, 'Beautiful lakeside venue for weddings and parties');

-- Seed experienceslist
INSERT INTO experienceslist (name, description, address, city, state, zip_code, phone_number, price, image_url, is_active, email, website) VALUES
('Hot Air Balloon Ride', 'Experience the thrill of flight', '500 Balloon Rd', 'Springfield', 'IL', '62704', '555-6666', 150.00, 'https://example.com/balloon.jpg', true, 'TestEmail@gmail.com', 'http://hotairballoon.com'),
('Cooking Class', 'Learn gourmet recipes', '300 Chef St', 'Springfield', 'IL', '62705', '555-7777', 80.00, 'https://example.com/cooking.jpg', true, 'TestEmail@gmail.com', null),
('Wine Tasting', 'Sample fine wines', '600 Vineyard Ln', 'Springfield', 'IL', '62706', '555-8888', 60.00, 'https://example.com/wine.jpg', false,'TestEmail@gmail.com', null);

-- Seed resteruatelist (restaurants)
INSERT INTO resteruatelist (name, description, address, city, state, zip_code, image_url, website, is_active, phone_number, email) VALUES
('The Fancy Fork', 'Upscale dining experience', '400 Gourmet Ln', 'Springfield', 'IL', '62706', 'https://example.com/fancyfork.jpg', 'http://fancyfork.com', true, '555-1010', 'FakeEmail@gmail.com'),
('Pizza Palace', 'Family-friendly pizza place', '600 Slice St', 'Springfield', 'IL', '62707', 'https://example.com/pizzapalace.jpg', 'http://pizzapalace.com', true, '555-2020', 'FakeEmail22@gmail.com'),
('Burger Barn', 'Casual burger joint', '700 Grill Ave', 'Springfield', 'IL', '62708', 'https://example.com/burgerbarn.jpg', 'http://burgerbarn.com', false, '555-3030', 'FakeEmail23@gmail.com');

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
INSERT INTO reservations (user_id, partner_id, event_id, restaurant_id, experience_id, venue_id, reservation_time, party_size, special_requests) VALUES
(1, 1, 1, 1, NULL, 1, '2025-04-25 19:00:00', 4, 'Window seat please'),
(2, 2, 2, NULL, 1, 2, '2025-07-10 12:00:00', 2, NULL);

-- Seed reminders
INSERT INTO reminders (partner_id, reservation_id, "ScheduleDate", "ScheduleTime", "ReminderTimeBefore") VALUES
(1, 1, '2025-04-24', '17:00:00', 120),
(2, 2, '2025-07-09', '10:00:00', 60);

