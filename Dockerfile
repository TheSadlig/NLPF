FROM golang

WORKDIR "/home"

ENV GOPATH=/home
ENV PATH=$PATH:/usr/local/go/bin

ADD install_couchdb_jessie.sh /srv
WORKDIR "/srv"
RUN chmod +x install_couchdb_jessie.sh

RUN ./install_couchdb_jessie.sh

ADD launch.sh /srv
RUN chmod +x launch.sh

RUN go get github.com/rhinoman/couchdb-go
RUN go get github.com/gin-gonic/gin
RUN go get github.com/nu7hatch/gouuid

RUN mkdir go-server
WORKDIR /srv/go-server

ADD go-server /srv/go-server

RUN go build .
CMD /srv/launch.sh

EXPOSE 80 5984