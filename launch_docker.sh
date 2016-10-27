docker build -t nlpf .
sudo docker run -p 5984:5984 -p 9090:9090 -i -t nlpf
