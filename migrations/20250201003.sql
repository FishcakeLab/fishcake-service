ALTER TABLE wallet_addresses ADD COLUMN device varchar COLLATE "pg_catalog"."default" NOT NULL;
ALTER TABLE activity_participants_addresses ADD COLUMN "status" int2 DEFAULT 0;