CREATE TYPE  status AS ENUM ('PENDING', 'ACTIVE', 'INACTIVE');

CREATE TABLE "services" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "status" status NOT NULL DEFAULT 'PENDING',
    "name" varchar NOT NULL,
    "password" varchar NOT NULL,
    "deleted_at" timestamptz,
    "created_at" timestamptz NOT NULL default now(),
    "updated_at" timestamptz NOT NULL default now()
);
CREATE UNIQUE INDEX services_name_key on services (name) where deleted_at IS NULL;
