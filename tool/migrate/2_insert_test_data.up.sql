INSERT INTO
  users(name, firebase_uid)
VALUES
  ('test1', 'uid1'),
  ('test2', 'uid2'),
  ('test3', 'uid3');

INSERT INTO
  routines(name, user_id, started_at)
VALUES
  ("rotine 1", 1, "2021-01-21 22:30:00"),
  ("rotine 2", 2, "2021-01-21 22:30:00"),
  ("rotine 3", 3, "2021-01-21 22:30:00");

INSERT INTO
  progress(routine_id, date)
VALUES
  (1, "2021-01-21"),
  (2, "2021-01-21"),
  (3, "2021-01-21");