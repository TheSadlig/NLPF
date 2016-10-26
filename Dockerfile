FROM golang

WORKDIR "/home"

ENV GOPATH=/home
ENV PATH=$PATH:/usr/local/go/bin
RUN go get github.com/gin-gonic/gin

ADD install_couchdb_jessie.sh /srv
WORKDIR "/srv"
RUN chmod +x install_couchdb_jessie.sh
RUN ./install_couchdb_jessie.sh
RUN go get github.com/rhinoman/couchdb-go

ADD . /srv/
RUN go build main.go
RUN chmod +x launch.sh
CMD ./launch.sh


EXPOSE 80 5984