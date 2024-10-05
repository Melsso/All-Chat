FROM nginx:alpine

COPY nginx.conf /etc/nginx/nginx.conf

COPY localhost-key.pem /etc/nginx/
COPY localhost.pem /etc/nginx/

COPY static/ /usr/share/nginx/html/static/

EXPOSE 443