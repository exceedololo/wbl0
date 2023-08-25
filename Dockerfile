# Используем официальный образ PostgreSQL
FROM postgres:latest

# Определяем переменные окружения для создания пользователя
ENV POSTGRES_USER wbadmin
ENV POSTGRES_PASSWORD 19az%&ty56

# Копируем файлы инициализации базы данных в контейнер
COPY wborderfile.sql /docker-entrypoint-initdb.d/
