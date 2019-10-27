FROM alpine:latest

# Install ruby dependency
RUN apk update && apk add curl ruby ruby-bundler ruby-dev build-base curl-dev libxml2-dev libxslt-dev pkgconfig postgresql-dev
# Clean APK cache
RUN rm -rf /var/cache/apk/*

COPY Gemfile /Gemfile
COPY Rakefile /Rakefile
RUN gem install pkg-config -v "~> 1.1" --no-document && bundle install

COPY ./sewadong-service srv/sewadong-service

# Migration data
RUN mkdir -p /db/migrate
COPY ./db/* /db/
COPY ./db/migrate/* /db/migrate/

# Swagger file
COPY ./contents/swagger-ui /contents/swagger-ui

#RUN find $HOME -type d -exec 'chmod' '555' '{}' ';' && \
#    find $HOME -type f -exec 'chmod' '444' '{}' ';' && \
#    find $HOME -type f -exec 'chown' 'root:root' '{}' ';' && \
#    chmod 555 $HOME/braggart-service && \
#    chmod 555 $HOME/contents
#
#USER nobody

CMD ["/srv/sewadong-service"]