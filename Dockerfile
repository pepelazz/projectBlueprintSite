FROM alpine
# Update package index
RUN apk add --no-cache tzdata
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apk update && apk add ca-certificates && apk add --update curl && apk add zip && rm -rf /var/cache/apk/*

COPY ./src/app /app
COPY ./src/sql /sql
COPY ./src/webClient/dist /webClient/dist

COPY ./src/webServer/site /webServer/site
COPY ./src/webServer/templates /webServer/templates
RUN chmod -Rf 777 /app
ENTRYPOINT ["/app"]

