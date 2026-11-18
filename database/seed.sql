
-- =========================================
-- LIBRARY MANAGEMENT SYSTEM - DEMO BOOKS
-- =========================================

INSERT INTO books
(title, author, category, isbn, total_copies, available_copies, description)
VALUES
('The Alchemist', 'Paulo Coelho', 'Fiction', '9780061122415', 4, 3, 'A novel about following dreams and discovering your destiny.'),
('Atomic Habits', 'James Clear', 'Self Help', '9780735211292', 5, 4, 'A practical guide to building good habits.'),
('Rich Dad Poor Dad', 'Robert Kiyosaki', 'Finance', '9781612680194', 4, 3, 'A book about financial education and wealth building.'),
('The Psychology of Money', 'Morgan Housel', 'Finance', '9780857197689', 3, 2, 'Lessons about money, behaviour and investing.'),
('Ikigai', 'Hector Garcia', 'Self Help', '9780143130727', 4, 3, 'Explores the Japanese concept of finding purpose.'),
('The Great Gatsby', 'F. Scott Fitzgerald', 'Fiction', '9780743273565', 3, 2, 'A classic novel set in the Jazz Age.'),
('Wings of Fire', 'A. P. J. Abdul Kalam', 'Biography', '9788173711466', 5, 4, 'The autobiography of Dr. A. P. J. Abdul Kalam.'),
('Think and Grow Rich', 'Napoleon Hill', 'Self Help', '9781585424337', 3, 2, 'A classic guide to personal success.'),
('The Intelligent Investor', 'Benjamin Graham', 'Finance', '9780060555665', 4, 3, 'A classic introduction to value investing.'),
('The Hobbit', 'J. R. R. Tolkien', 'Fantasy', '9780547928227', 4, 3, 'A fantasy adventure in Middle-earth.'),
('Pride and Prejudice', 'Jane Austen', 'Fiction', '9780141439518', 3, 3, 'A classic novel about love and society.'),
('Deep Work', 'Cal Newport', 'Productivity', '9781455586691', 4, 4, 'A guide to focused and meaningful work.'),
('The Lean Startup', 'Eric Ries', 'Business', '9780307887894', 3, 3, 'A method for building successful startups.'),
('Sapiens', 'Yuval Noah Harari', 'History', '9780062316097', 4, 4, 'A history of humankind.'),
('The Power of Now', 'Eckhart Tolle', 'Self Help', '9781577314806', 3, 3, 'A guide to mindfulness and living in the present.'),
('To Kill a Mockingbird', 'Harper Lee', 'Fiction', '9780060935467', 4, 4, 'A novel exploring justice and compassion.'),
('The 7 Habits of Highly Effective People', 'Stephen Covey', 'Productivity', '9781982137274', 3, 3, 'Principles for personal and professional effectiveness.'),
('The Subtle Art of Not Giving a F*ck', 'Mark Manson', 'Self Help', '9780062457714', 4, 4, 'A different approach to happiness and priorities.'),
('The Little Prince', 'Antoine de Saint-Exupery', 'Fiction', '9780156012195', 3, 3, 'A philosophical tale about life and friendship.'),
('A Brief History of Time', 'Stephen Hawking', 'Science', '9780553380163', 4, 4, 'An introduction to cosmology and the universe.'),
('Zero to One', 'Peter Thiel', 'Business', '9780804139298', 3, 3, 'Ideas about innovation and building businesses.'),
('The Monk Who Sold His Ferrari', 'Robin Sharma', 'Self Help', '9780062515674', 4, 4, 'A story about personal growth and fulfilment.'),
('The Da Vinci Code', 'Dan Brown', 'Thriller', '9780307474278', 3, 3, 'A mystery thriller involving art and history.'),
('Educated', 'Tara Westover', 'Biography', '9780399590504', 3, 3, 'A memoir about education and self-discovery.'),
('The Psychology of Selling', 'Brian Tracy', 'Business', '9780785288069', 3, 3, 'Principles of sales and communication.'),
('Richest Man in Babylon', 'George S. Clason', 'Finance', '9780451205360', 4, 4, 'Financial lessons told through parables.'),
('Dune', 'Frank Herbert', 'Science Fiction', '9780441172719', 3, 3, 'A science fiction epic set on a desert planet.'),
('Start With Why', 'Simon Sinek', 'Business', '9781591846444', 3, 3, 'Explores how purpose inspires leadership.'),
('Man''s Search for Meaning', 'Viktor E. Frankl', 'Psychology', '9780807014271', 4, 4, 'Reflections on meaning and resilience.'),
('The Secret', 'Rhonda Byrne', 'Self Help', '9781582701707', 3, 3, 'A popular book about positive thinking.');



-- =========================================
-- DEMO MEMBERS
-- =========================================

INSERT INTO members
(name, email, phone, address, membership_type, membership_date, membership_status)
VALUES
('Arun Kumar', 'arun.kumar@example.com', '9876543210', 'Kochi, Kerala', 'Regular', CURRENT_DATE - 180, 'Active'),
('Priya Nair', 'priya.nair@example.com', '9876543211', 'Thrissur, Kerala', 'Premium', CURRENT_DATE - 150, 'Active'),
('Rahul Menon', 'rahul.menon@example.com', '9876543212', 'Kottayam, Kerala', 'Regular', CURRENT_DATE - 120, 'Active'),
('Anjali Thomas', 'anjali.thomas@example.com', '9876543213', 'Alappuzha, Kerala', 'Premium', CURRENT_DATE - 100, 'Active'),
('Adithya Raj', 'adithya.raj@example.com', '9876543214', 'Kozhikode, Kerala', 'Regular', CURRENT_DATE - 90, 'Active'),
('Sneha Mathew', 'sneha.mathew@example.com', '9876543215', 'Ernakulam, Kerala', 'Student', CURRENT_DATE - 75, 'Active'),
('Vishnu Prasad', 'vishnu.prasad@example.com', '9876543216', 'Palakkad, Kerala', 'Regular', CURRENT_DATE - 60, 'Active'),
('Meera Joseph', 'meera.joseph@example.com', '9876543217', 'Kannur, Kerala', 'Premium', CURRENT_DATE - 50, 'Active'),
('Nikhil Das', 'nikhil.das@example.com', '9876543218', 'Idukki, Kerala', 'Student', CURRENT_DATE - 40, 'Active'),
('Aparna Suresh', 'aparna.suresh@example.com', '9876543219', 'Kollam, Kerala', 'Regular', CURRENT_DATE - 30, 'Active'),
('Joel George', 'joel.george@example.com', '9876543220', 'Pathanamthitta, Kerala', 'Student', CURRENT_DATE - 20, 'Active'),
('Kavya Ramesh', 'kavya.ramesh@example.com', '9876543221', 'Malappuram, Kerala', 'Premium', CURRENT_DATE - 10, 'Active');



-- =========================================
-- DEMO BORROWING TRANSACTIONS
-- 10 currently issued + 10 returned
-- =========================================

INSERT INTO loans
(book_id, member_id, issue_date, due_date, returned_at, status)
VALUES

-- OVERDUE LOANS (4)
(
  (SELECT id FROM books WHERE isbn = '9780061122415'),
  (SELECT id FROM members WHERE email = 'arun.kumar@example.com'),
  CURRENT_DATE - 20, CURRENT_DATE - 5, NULL, 'Issued'
),
(
  (SELECT id FROM books WHERE isbn = '9780735211292'),
  (SELECT id FROM members WHERE email = 'priya.nair@example.com'),
  CURRENT_DATE - 18, CURRENT_DATE - 3, NULL, 'Issued'
),
(
  (SELECT id FROM books WHERE isbn = '9781612680194'),
  (SELECT id FROM members WHERE email = 'rahul.menon@example.com'),
  CURRENT_DATE - 16, CURRENT_DATE - 2, NULL, 'Issued'
),
(
  (SELECT id FROM books WHERE isbn = '9780857197689'),
  (SELECT id FROM members WHERE email = 'anjali.thomas@example.com'),
  CURRENT_DATE - 14, CURRENT_DATE - 1, NULL, 'Issued'
),

-- ACTIVE LOANS (6)
(
  (SELECT id FROM books WHERE isbn = '9780143130727'),
  (SELECT id FROM members WHERE email = 'adithya.raj@example.com'),
  CURRENT_DATE - 5, CURRENT_DATE + 9, NULL, 'Issued'
),
(
  (SELECT id FROM books WHERE isbn = '9780743273565'),
  (SELECT id FROM members WHERE email = 'sneha.mathew@example.com'),
  CURRENT_DATE - 4, CURRENT_DATE + 10, NULL, 'Issued'
),
(
  (SELECT id FROM books WHERE isbn = '9788173711466'),
  (SELECT id FROM members WHERE email = 'vishnu.prasad@example.com'),
  CURRENT_DATE - 3, CURRENT_DATE + 11, NULL, 'Issued'
),
(
  (SELECT id FROM books WHERE isbn = '9781585424337'),
  (SELECT id FROM members WHERE email = 'meera.joseph@example.com'),
  CURRENT_DATE - 2, CURRENT_DATE + 12, NULL, 'Issued'
),
(
  (SELECT id FROM books WHERE isbn = '9780060555665'),
  (SELECT id FROM members WHERE email = 'nikhil.das@example.com'),
  CURRENT_DATE - 2, CURRENT_DATE + 13, NULL, 'Issued'
),
(
  (SELECT id FROM books WHERE isbn = '9780547928227'),
  (SELECT id FROM members WHERE email = 'aparna.suresh@example.com'),
  CURRENT_DATE - 1, CURRENT_DATE + 14, NULL, 'Issued'
),

-- RETURNED LOANS (10)
(
  (SELECT id FROM books WHERE isbn = '9780141439518'),
  (SELECT id FROM members WHERE email = 'arun.kumar@example.com'),
  CURRENT_DATE - 30, CURRENT_DATE - 16, CURRENT_DATE - 18, 'Returned'
),
(
  (SELECT id FROM books WHERE isbn = '9781455586691'),
  (SELECT id FROM members WHERE email = 'priya.nair@example.com'),
  CURRENT_DATE - 28, CURRENT_DATE - 14, CURRENT_DATE - 17, 'Returned'
),
(
  (SELECT id FROM books WHERE isbn = '9780307887894'),
  (SELECT id FROM members WHERE email = 'rahul.menon@example.com'),
  CURRENT_DATE - 25, CURRENT_DATE - 11, CURRENT_DATE - 12, 'Returned'
),
(
  (SELECT id FROM books WHERE isbn = '9780062316097'),
  (SELECT id FROM members WHERE email = 'anjali.thomas@example.com'),
  CURRENT_DATE - 23, CURRENT_DATE - 9, CURRENT_DATE - 10, 'Returned'
),
(
  (SELECT id FROM books WHERE isbn = '9781577314806'),
  (SELECT id FROM members WHERE email = 'adithya.raj@example.com'),
  CURRENT_DATE - 21, CURRENT_DATE - 7, CURRENT_DATE - 8, 'Returned'
),
(
  (SELECT id FROM books WHERE isbn = '9780060935467'),
  (SELECT id FROM members WHERE email = 'sneha.mathew@example.com'),
  CURRENT_DATE - 19, CURRENT_DATE - 5, CURRENT_DATE - 6, 'Returned'
),
(
  (SELECT id FROM books WHERE isbn = '9781982137274'),
  (SELECT id FROM members WHERE email = 'vishnu.prasad@example.com'),
  CURRENT_DATE - 17, CURRENT_DATE - 3, CURRENT_DATE - 4, 'Returned'
),
(
  (SELECT id FROM books WHERE isbn = '9780062457714'),
  (SELECT id FROM members WHERE email = 'meera.joseph@example.com'),
  CURRENT_DATE - 15, CURRENT_DATE - 1, CURRENT_DATE - 2, 'Returned'
),
(
  (SELECT id FROM books WHERE isbn = '9780156012195'),
  (SELECT id FROM members WHERE email = 'nikhil.das@example.com'),
  CURRENT_DATE - 13, CURRENT_DATE - 1, CURRENT_DATE - 3, 'Returned'
),
(
  (SELECT id FROM books WHERE isbn = '9780553380163'),
  (SELECT id FROM members WHERE email = 'aparna.suresh@example.com'),
  CURRENT_DATE - 11, CURRENT_DATE + 3, CURRENT_DATE - 2, 'Returned'
);
