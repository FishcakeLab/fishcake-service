DO $$
BEGIN
    -- check wallet_addresses columns device 
    IF NOT EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name='wallet_addresses' 
        AND column_name='device'
    ) THEN
        ALTER TABLE wallet_addresses ADD COLUMN device varchar COLLATE "pg_catalog"."default" NOT NULL;
    END IF;

    -- check activity_participants_addresses  columns status
    IF NOT EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name='activity_participants_addresses' 
        AND column_name='status'
    ) THEN
        ALTER TABLE activity_participants_addresses ADD COLUMN "status" int2 DEFAULT 0;
    END IF;
END
$$;