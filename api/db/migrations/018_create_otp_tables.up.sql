CREATE TABLE sp_otps (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       phone varchar(20) NOT NULL UNIQUE,
       code varchar(10) NOT NULL,
       expiry TIMESTAMPTZ NOT NULL,
       writer uuid NOT NULL
);
