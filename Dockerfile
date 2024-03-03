FROM docker:latest

EXPOSE 3000

WORKDIR /app
ENV HOST=localhost DBPORT=5432

ENV USER=root PASSWORD=root DBNAME=root

ARG STAGE_DEPLOY=true

RUN COPY --from=intermediate ./main main

CMD [ "./main" ]
