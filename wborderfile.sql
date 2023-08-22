CREATE DATABASE IF NOT EXISTS wborderdb;
CREATE SCHEMA IF NOT EXISTS wborderscheme;
CREATE USER wbadmin WITH PASSWORD '19az%&ty56';
GRANT USAGE ON SCHEMA wborderscheme TO wbadmin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA wborderscheme TO wbadmin;

CREATE TABLE IF NOT EXISTS wborderdb.orders(
                                               order_uid TEXT PRIMARY KEY,
                                               date_created TIMESTAMP,
                                               data JSONB
)