FROM ubuntu:latest

EXPOSE 8000

WORKDIR /app

ENV HOST=localhost DBPORT=5432

ENV USER=root PASSWORD=root DBNAME=root

COPY ./main main
RUN apt-get update && apt-get install -y \
    curl \
    vim
#CMD ["tail", "-f", "/dev/null"]
CMD ls
CMD [ "./main" ]
