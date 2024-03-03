FROM ubuntu:latest

EXPOSE 8000

WORKDIR /app

ENV DB_HOST=localhost DB_PORT=5432
ENV DB_USER=root DB_PASSWORD=root DB_NAME=root

COPY ./main main

RUN apt-get update && apt-get install -y \
    curl \
    vim

CMD ["tail", "-f", "/dev/null"]
#CMD [ "./main" ]
