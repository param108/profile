CREATE TABLE onetime (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       data TEXT NOT NULL,
       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
