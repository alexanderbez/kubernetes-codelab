FROM golang:1.7
ENV APP_HOME /app
RUN mkdir $APP_HOME
ADD pkg/hello-world $APP_HOME

CMD /app/hello-world
