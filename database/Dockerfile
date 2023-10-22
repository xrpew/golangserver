FROM postgres:15.3

# Copia el script SQL a la ubicación adecuada
COPY up.sql /docker-entrypoint-initdb.d/

CMD [ "postgres" ]