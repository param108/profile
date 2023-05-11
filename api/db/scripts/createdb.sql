CREATE USER profile WITH password 'profile';
ALTER USER profile WITH login;
CREATE DATABASE profile;
GRANT ALL ON DATABASE profile to profile;
ALTER DATABASE profile owner to profile;
\c profile
CREATE EXTENSION "uuid-ossp";
