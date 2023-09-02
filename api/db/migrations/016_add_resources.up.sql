CREATE TABLE resources (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       user_id uuid NOT NULL,
       t varchar(50) NOT NULL,
       value int default 0 CHECK (value >= 0 and value <= max),
       max int default 0,
       writer uuid NOT NULL,
       UNIQUE(user_id, writer, t)
);

CREATE INDEX idx_resources_user_id_type ON resources(user_id, writer, t);
