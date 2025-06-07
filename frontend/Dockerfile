# Используем минимальный образ nginx
FROM nginx:alpine

# Удаляем дефолтную заглушку nginx
RUN rm -rf /usr/share/nginx/html/*

# Копируем весь фронтенд в директорию, которую обслуживает nginx
COPY . /usr/share/nginx/html

# Порт 80 — это внутренний порт nginx, он будет проброшен наружу
EXPOSE 80

# Команда по умолчанию (можно не указывать, так как это стандартное поведение nginx)
CMD ["nginx", "-g", "daemon off;"]
