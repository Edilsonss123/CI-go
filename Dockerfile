FROM docker:latest

EXPOSE 3000

WORKDIR /app
ENV HOST=localhost DBPORT=5432

ENV USER=root PASSWORD=root DBNAME=root

RUN echo "\n$STAGE_DEPLOY\n\" && if [ "$STAGE_DEPLOY" = "true" ]; then \
       COPY ./main main; \
    fi


ENTRYPOINT [ "/main" ]
