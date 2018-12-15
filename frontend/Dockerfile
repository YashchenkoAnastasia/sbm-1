FROM node:10 as build
# Stage 1 - the build process
RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN npm install -g ember-cli
RUN npm install
RUN ember build -e production

FROM nginx:1.15.3

ENV STATIC_FOLDERS_REGEX "^/(assets|fonts)/"
COPY --from=build /app/dist /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf \
    && ln -s /usr/share/nginx/html /app
COPY docker/nginx.conf /etc/nginx/conf.d/app.conf
COPY docker/ember-proxy-service.sh /

EXPOSE 80

CMD ["/bin/bash", "/ember-proxy-service.sh"]
