nohup /etc/init.d/couchdb start&
curl -X PUT http://localhost:5984/_config/admins/user -d '"password"'
./main
