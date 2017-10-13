ARG NGINX_VERSION
FROM nginx:${NGINX_VERSION}	

COPY nginx/index.html /var/www/polarfalcon.com/public_html/
COPY nginx/polarfalcon.conf /etc/nginx/conf.d/
