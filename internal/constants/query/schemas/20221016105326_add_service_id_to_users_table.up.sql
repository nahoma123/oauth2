ALTER TABLE users ADD "service_id" UUID NOT NULL;
ALTER TABLE users ADD FOREIGN KEY("service_id") REFERENCES services("id") ON DELETE CASCADE;
