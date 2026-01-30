insert into users(username, password, role, created_at) values
('admin', 'admin', 2, datetime('now')),
('test', 'test', 0, datetime('now'));

insert into submissions(user_id, submission_date, created_at) values
(1, datetime('now', '-1day'), datetime('now')),
(2, datetime('now'), datetime('now'));

insert into work_items(submission_id, name, description, created_at) values
(1, 'haircut', 'make a haircut?', datetime('now')),
(2, 'item', 'item desc', datetime('now'));