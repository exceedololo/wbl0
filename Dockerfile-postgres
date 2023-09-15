# Используем официальный образ PostgreSQL
FROM postgres:latest

# Указываем переменные окружения для настройки PostgreSQL
ENV POSTGRES_USER wbadmin
ENV POSTGRES_PASSWORD 19azty56
ENV POSTGRES_DB wborderdb

# Копируем SQL-файл внутрь контейнера
COPY wborderfile.sql /docker-entrypoint-initdb.d/

# Открываем порт для подключения к базе данных
EXPOSE 5432