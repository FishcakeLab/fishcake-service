DO $$
BEGIN
IF NOT EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name='drop_info'
        AND column_name='system_drop_type'
    ) THEN
ALTER TABLE drop_info ADD COLUMN system_drop_type int2 DEFAULT 0;
END IF;

END
$$;