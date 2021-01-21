TRUNCATE TABLE progress;

-- 外部キー制約を一時的にはずす
-- これをしないとTRUNCATE できない
SET foreign_key_checks = 0;
TRUNCATE TABLE routines;
-- TRUNCATEした後は戻す
SET foreign_key_checks = 1;