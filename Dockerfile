FROM ubuntu
ADD ./httpserver /httpserver
ENTRYPOINT ["/httpserver"]