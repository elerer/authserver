FROM scratch
MAINTAINER j@gmail.com

EXPOSE 8081

COPY ./authserver ./authserver
ADD ./files ./files

ENTRYPOINT ["./authserver"]
