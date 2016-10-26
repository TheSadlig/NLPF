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
RUN mkdir src
WORKDIR /srv/src

ADD src /srv/src

RUN go build main.go
CMD /srv/launch.sh

EXPOSE 80 5984