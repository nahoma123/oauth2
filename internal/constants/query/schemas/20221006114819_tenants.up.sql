CREATE TABLE "tenants" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "status" status NOT NULL DEFAULT 'ACTIVE',
    "tenant_name" varchar NOT NULL,
    "service_id" UUID NOT NULL,
    "deleted_at" timestamptz,
    "created_at" timestamptz NOT NULL default now(),
    "updated_at" timestamptz NOT NULL default now(),
    UNIQUE("tenant_name","service_id")
);

ALTER TABLE "tenants" add FOREIGN KEY ("service_id") REFERENCES "services"  ("id") ON DELETE CASCADE;