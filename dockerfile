FROM php:8.1-cli

WORKDIR /var/www/html

COPY index.php .

CMD ["php", "index.php"]
