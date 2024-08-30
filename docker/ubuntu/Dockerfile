FROM ubuntu:latest


COPY ./build .
RUN apt update
CMD apt install nodejs && apt install npm && npm i -g http-server && http-server