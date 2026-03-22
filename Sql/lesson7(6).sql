--1
EXPLAIN ANALYZE
SELECT *
FROM users
WHERE email = 'student5000@mail.com';

CREATE INDEX idx_users_email ON users(email);

--2
EXPLAIN ANALYZE
SELECT *
FROM attempts
WHERE user_id = 25000;

CREATE INDEX idx_attempts_user_id ON attempts(user_id);

--3
EXPLAIN ANALYZE
SELECT *
FROM attempts
WHERE user_id = 25000
ORDER BY created_at DESC;

CREATE INDEX idx_attempts_user_id_created_at 
ON attempts(user_id, created_at DESC);

--4
EXPLAIN ANALYZE
SELECT *
FROM attempts
WHERE status = 'in_progress';

CREATE INDEX idx_attempts_in_progress
ON attempts(status)
WHERE status = 'in_progress';

--5
EXPLAIN ANALYZE
SELECT 
    u.full_name,
    s.title AS subject_title,
    a.score,
    a.created_at
FROM attempts a
JOIN users u ON u.id = a.user_id
JOIN tests t ON t.id = a.test_id
JOIN subjects s ON s.id = t.subject_id
WHERE s.title = 'Python'
ORDER BY a.score DESC;

CREATE INDEX idx_subjects_title ON subjects(title);
CREATE INDEX idx_tests_subject_id ON tests(subject_id);
CREATE INDEX idx_attempts_test_id_score ON attempts(test_id, score DESC);
CREATE INDEX idx_users_id ON users(id);
