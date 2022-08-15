FROM golang:1.16-alpine

RUN apk add wget && apk update
RUN apk add --no-cache bash
WORKDIR /app 
COPY . /app/

RUN go mod download

# ARGUMENT
ARG LICENSEY_KEY
ARG API_FGIP
ARG PORT
ARG ROLLBARTOKEN

## ENVIRONMENT
ENV API_KEY=${API_FGIP}
ENV PORT=${PORT}
ENV ROLLBAR_TOKEN=${ROLLBARTOKEN}

RUN go build -o /app/gopi
RUN echo "export licensekey=${LICENSEY_KEY}" >> /app/config/maxmind.config
EXPOSE 3045

CMD ["sh", "start.sh"]
