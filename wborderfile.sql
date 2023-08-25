-- Создание базы данных, если она не существует
DO $$BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'wborderdb') THEN
        CREATE DATABASE wborderdb;
END IF;
END$$;

-- Подключение к базе данных wborderdb
\c wborderdb;

-- Создание схемы wborderscheme, если она не существует
CREATE SCHEMA IF NOT EXISTS wborderscheme;

-- Создание пользователя wbadmin, если он не существует
DO $$BEGIN
    IF NOT EXISTS (SELECT FROM pg_user WHERE usename = 'wbadmin') THEN
        CREATE USER wbadmin WITH PASSWORD '19az%&ty56';
END IF;
END$$;

-- Предоставление прав на использование схемы пользователю wbadmin
GRANT USAGE ON SCHEMA wborderscheme TO wbadmin;

-- Предоставление всех привилегий на таблицы схемы wborderscheme пользователю wbadmin
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA wborderscheme TO wbadmin;

-- Создание таблицы orders в схеме wborderscheme
CREATE TABLE IF NOT EXISTS wborderscheme.orders(
                                                   order_uid TEXT PRIMARY KEY,
                                                   date_created TIMESTAMP,
                                                   data JSONB
);
