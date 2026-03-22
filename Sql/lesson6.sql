--1
CREATE FUNCTION lessons_count_safe(p_course_id INT)
RETURNS INT
LANGUAGE sql
AS $$
    SELECT COALESCE(
        (SELECT COUNT(*) FROM lessons WHERE course_id = p_course_id),
        0
    );
$$;

--2
CREATE FUNCTION student_started_count(p_student INT, p_course INT)
RETURNS INT
LANGUAGE sql
AS $$
    SELECT COUNT(*)
    FROM progress pr
    JOIN lessons l ON l.id = pr.lesson_id
    WHERE pr.student_id = p_student
      AND l.course_id = p_course
      AND pr.status = 'in_progress';
$$;

--3
CREATE FUNCTION student_avg_score_done(p_student INT, p_course INT)
RETURNS NUMERIC(5,2)
LANGUAGE sql
AS $$
    SELECT COALESCE(ROUND(AVG(pr.score)::numeric, 2), 0)
    FROM progress pr
    JOIN lessons l ON l.id = pr.lesson_id
    WHERE pr.student_id = p_student
      AND l.course_id = p_course
      AND pr.status = 'done';
$$;

--4
CREATE VIEW v_student_course_progress AS
SELECT 
    s.id AS student_id,
    s.full_name,
    c.id AS course_id,
    c.title AS course_title,
    (SELECT COUNT(*) FROM lessons l WHERE l.course_id = c.id) AS total_lessons,
    (SELECT COUNT(*) 
     FROM progress pr
     JOIN lessons l ON l.id = pr.lesson_id
     WHERE pr.student_id = s.id AND l.course_id = c.id AND pr.status = 'done') AS done_lessons,
    (SELECT COUNT(*) 
     FROM progress pr
     JOIN lessons l ON l.id = pr.lesson_id
     WHERE pr.student_id = s.id AND l.course_id = c.id AND pr.status = 'in_progress') AS in_progress_lessons,
    CASE 
        WHEN (SELECT COUNT(*) FROM lessons l WHERE l.course_id = c.id) = 0 THEN 0
        ELSE ROUND(
            100.0 * 
            (SELECT COUNT(*) 
             FROM progress pr
             JOIN lessons l ON l.id = pr.lesson_id
             WHERE pr.student_id = s.id AND l.course_id = c.id AND pr.status = 'done') 
            / (SELECT COUNT(*) FROM lessons l WHERE l.course_id = c.id), 2)
    END AS completion_percent
FROM enrollments e
JOIN students s ON s.id = e.student_id
JOIN courses c ON c.id = e.course_id;

--5 
CREATE FUNCTION student_course_lessons(p_student INT, p_course INT)
RETURNS TABLE(position INT, lesson_title TEXT, status TEXT, score INT)
LANGUAGE sql
AS $$
    SELECT 
        l.position,
        l.title,
        COALESCE(pr.status, 'not_started') AS status,
        COALESCE(pr.score, 0) AS score
    FROM lessons l
    LEFT JOIN progress pr 
        ON pr.lesson_id = l.id AND pr.student_id = p_student
    WHERE l.course_id = p_course
    ORDER BY l.position;
$$;

--6
SELECT 
    s.full_name AS student,
    COUNT(*) AS done_count
FROM progress pr
JOIN students s ON s.id = pr.student_id
WHERE pr.status = 'done'
GROUP BY s.id, s.full_name
ORDER BY done_count DESC
LIMIT 3;

--7
SELECT 
    c.title AS course_title,
    l.title AS lesson_title,
    COUNT(*) AS in_progress_cnt
FROM progress pr
JOIN lessons l ON l.id = pr.lesson_id
JOIN courses c ON c.id = l.course_id
WHERE pr.status = 'in_progress'
GROUP BY c.title, l.title
HAVING COUNT(*) >= 2;

--8
CREATE FUNCTION best_course_for_student(p_student INT)
RETURNS TABLE(course_id INT, course_title TEXT, completion_percent NUMERIC(5,2))
LANGUAGE sql
AS $$
    SELECT course_id, course_title, completion_percent
    FROM v_student_course_progress
    WHERE student_id = p_student
    ORDER BY completion_percent DESC
    LIMIT 1;
$$;

--9
SELECT 
    s.email AS student_email,
    c.title AS course_title,
    l.title AS lesson_title,
    pr.status
FROM progress pr
JOIN lessons l ON l.id = pr.lesson_id
JOIN courses c ON c.id = l.course_id
JOIN students s ON s.id = pr.student_id
LEFT JOIN enrollments e 
    ON e.student_id = pr.student_id AND e.course_id = l.course_id
WHERE e.student_id IS NULL;

--10
CREATE FUNCTION student_done_count(p_student INT, p_course INT)
RETURNS INT
LANGUAGE sql
AS $$
    SELECT COUNT(*)
    FROM progress pr
    WHERE pr.student_id = p_student
      AND pr.status = 'done'
      AND pr.lesson_id IN (SELECT id FROM lessons WHERE course_id = p_course);
$$;
