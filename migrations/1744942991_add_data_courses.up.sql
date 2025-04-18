-- Add dummy data for everything
INSERT INTO public.users
(id, email, password_hash, first_name, last_name, "role", last_login, is_active, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
VALUES('2996d88a-cc8b-4161-b193-0f70f8e7b904'::uuid, 'hakimnaufal@gmail.com', '$2a$10$rb/K5O3cklthUVOXJRrOM.UNKV/TG8854AaMlnVTX/RKV3Ehmjg76', 'John', 'Kennedy', 'student', '2025-04-18 10:24:26.740', true, '2996d88a-cc8b-4161-b193-0f70f8e7b904'::uuid, '2025-04-18 10:24:09.324', '2996d88a-cc8b-4161-b193-0f70f8e7b904'::uuid, '2025-04-18 10:24:26.635', NULL, NULL);
INSERT INTO public.users
(id, email, password_hash, first_name, last_name, "role", last_login, is_active, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
VALUES('3e8889a3-9365-4a02-ae94-344c0fcaa303'::uuid, 'hakimnaufal212@gmail.com', '$2a$10$yLTAbgICweehC15FAjNY6eOeByjYB6mAHXFOn9Y2QELQB30BIUXn6', 'John', 'Kennedy', 'teacher', '2025-04-18 10:57:29.330', true, '3e8889a3-9365-4a02-ae94-344c0fcaa303'::uuid, '2025-04-17 15:29:29.663', '3e8889a3-9365-4a02-ae94-344c0fcaa303'::uuid, '2025-04-18 10:57:29.210', NULL, NULL);

INSERT INTO public.teachers
(user_id, department, title)
VALUES('3e8889a3-9365-4a02-ae94-344c0fcaa303'::uuid, '', '');

INSERT INTO public.students
(user_id, student_id, enrollment_year, "program")
VALUES('2996d88a-cc8b-4161-b193-0f70f8e7b904'::uuid, '62df64d4-fcef-4cdd-88b1-ec71a3911c83', 2025, 'deeznuts');


INSERT INTO public.courses
(id, code, "name", description, start_date, end_date, is_active, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
VALUES('a934f2da-e6ef-471a-96e8-dd58d05cf207'::uuid, 'MATH101', 'Introduction to Calculus', 'A foundational course covering limits, derivatives, and integrals.', NULL, NULL, true, NULL, '2025-04-18 09:29:03.851', NULL, '2025-04-18 09:29:03.851', NULL, NULL);
INSERT INTO public.courses
(id, code, "name", description, start_date, end_date, is_active, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
VALUES('cb84ebd1-eaea-4f04-a430-79c4e0a0b816'::uuid, 'ENG205', 'Advanced Composition', 'An upper-level course focusing on rhetoric, argumentation, and stylistic analysis.', NULL, NULL, true, NULL, '2025-04-18 09:29:06.913', NULL, '2025-04-18 09:29:06.913', NULL, NULL);
INSERT INTO public.courses
(id, code, "name", description, start_date, end_date, is_active, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
VALUES('44062635-7450-4583-9232-dd79b37fd31a'::uuid, 'ENG01', 'English for Kids', 'This is an introduction english for kids', '2024-08-26 15:29:41.000', '2025-08-26 15:29:41.000', true, '3e8889a3-9365-4a02-ae94-344c0fcaa303'::uuid, '2025-04-18 09:49:22.956', NULL, NULL, NULL, NULL);

INSERT INTO public.assignments
(id, title, "content", description, due_date, course_id, teacher_id, total_points, is_published, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
VALUES('9ab6e9ec-69ee-401d-94e8-9c7fbe16b4a2'::uuid, 'How to count in English', 'this is the content of assignment', 'this is a assignment how to write count', '2025-04-18 10:13:43.285', '44062635-7450-4583-9232-dd79b37fd31a'::uuid, '3e8889a3-9365-4a02-ae94-344c0fcaa303'::uuid, 100.00, true, '3e8889a3-9365-4a02-ae94-344c0fcaa303'::uuid, '2025-04-18 10:13:43.285', NULL, NULL, NULL, NULL);


INSERT INTO public.submissions
(id, assignment_id, student_id, teacher_id, submitted_at, "content", file_url, grade, feedback, graded_at, graded_by, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
VALUES('9a765862-d75d-4bf7-9481-386e8c4ca6ba'::uuid, '9ab6e9ec-69ee-401d-94e8-9c7fbe16b4a2'::uuid, '2996d88a-cc8b-4161-b193-0f70f8e7b904'::uuid, '3e8889a3-9365-4a02-ae94-344c0fcaa303'::uuid, '2025-04-18 10:30:50.308', 'one, two, three, four, five, six, sveen, eight, nine, then', '', 80.00, 'it should be seven and then should be ten, good job', '2025-04-18 11:00:44.550', '3e8889a3-9365-4a02-ae94-344c0fcaa303'::uuid, '2996d88a-cc8b-4161-b193-0f70f8e7b904'::uuid, '2025-04-18 10:30:50.308', '3e8889a3-9365-4a02-ae94-344c0fcaa303'::uuid, '2025-04-18 11:00:44.530', NULL, NULL);
