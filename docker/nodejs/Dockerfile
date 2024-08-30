# Use the official Nginx image.
FROM node:latest


COPY ./build .
RUN npm i -g http-server
EXPOSE 3000
EXPOSE 8080

CMD http-server
# Expose the port Nginx runs on.
