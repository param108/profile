-- Insert the one user
INSERT INTO users (id, handle, profile, role, writer) values ('a592e6ab-91d1-49a7-9435-ab3c04f77ab9', 'param108', 'twitter', 'user', 'd603b6aa-8fa9-40cd-9c45-8b793f110741') ON CONFLICT DO NOTHING;
