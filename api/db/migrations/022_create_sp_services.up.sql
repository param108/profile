CREATE TABLE IF NOT EXISTS sp_services (
       id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
       code varchar(100) UNIQUE NOT NULL,
       name TEXT NOT NULL,
       category varchar(100) NOT NULL,
       unit varchar(50) NOT NULL,
       description TEXT NOT NULL,
       short_description TEXT NOT NULL,
       question TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_sp_services_category ON sp_services(category);

CREATE INDEX IF NOT EXISTS idx_sp_services_code ON sp_services(code);
