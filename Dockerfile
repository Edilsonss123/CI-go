FROM docker:latest

EXPOSE 3000

WORKDIR /app
ENV HOST=localhost DBPORT=5432

ENV USER=root PASSWORD=root DBNAME=root

RUN echo "Checking STAGE_DEPLOY: $STAGE_DEPLOY" && \
    if [ "$STAGE_DEPLOY" = "true" ]; then \
        echo "Copying ./main to main" && \
        COPY ./main main; \
    fi


ENTRYPOINT [ "/main" ]
