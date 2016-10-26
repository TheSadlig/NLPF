#!/bin/bash
# Inspired by http://verbally.flimzy.com/install-couchdb-1-6-1-debian-8-2-jessie/

# Erlang
echo -e "deb http://packages.erlang-solutions.com/debian jessie contrib" |  tee /etc/apt/sources.list.d/erlang-solutions.list
wget -qO - http://packages.erlang-solutions.com/debian/erlang_solutions.asc |  apt-key add -

# Update packages
 apt-get update

# Install dependencies
 apt-get install build-essential curl libmozjs185-1.0 libmozjs185-dev libcurl4-openssl-dev libicu-dev wget curl -y

# Install older version of erlang for couch
 apt-get install erlang-dev=1:17.5.3 erlang-base=1:17.5.3 erlang-crypto=1:17.5.3 \
                      erlang-nox=1:17.5.3 erlang-inviso=1:17.5.3 erlang-runtime-tools=1:17.5.3 \
                      erlang-inets=1:17.5.3 erlang-edoc=1:17.5.3 erlang-syntax-tools=1:17.5.3 \
                      erlang-xmerl=1:17.5.3 erlang-corba=1:17.5.3 erlang-mnesia=1:17.5.3 \
                      erlang-os-mon=1:17.5.3 erlang-snmp=1:17.5.3 erlang-ssl=1:17.5.3 \
                      erlang-public-key=1:17.5.3 erlang-asn1=1:17.5.3 erlang-ssh=1:17.5.3 \
                      erlang-erl-docgen=1:17.5.3 erlang-percept=1:17.5.3 erlang-diameter=1:17.5.3 \
                      erlang-webtool=1:17.5.3 erlang-eldap=1:17.5.3 erlang-tools=1:17.5.3 \
                      erlang-eunit=1:17.5.3 erlang-ic=1:17.5.3 erlang-odbc=1:17.5.3 \
                      erlang-parsetools=1:17.5.3 -y

# Hold packages to avoid updates
 apt-mark hold erlang-dev erlang-base erlang-crypto erlang-nox erlang-inviso erlang-runtime-tools \
                      erlang-inets erlang-edoc erlang-syntax-tools erlang-xmerl erlang-corba \
                      erlang-mnesia erlang-os-mon erlang-snmp erlang-ssl erlang-public-key \
                      erlang-asn1 erlang-ssh erlang-erl-docgen erlang-percept erlang-diameter \
                      erlang-webtool erlang-eldap erlang-tools erlang-eunit erlang-ic erlang-odbc \
                      erlang-parsetools -y

# Set up couch environment
 useradd -d /var/lib/couchdb couchdb
 mkdir -p /usr/local/{lib,etc}/couchdb /usr/local/var/{lib,log,run}/couchdb /var/lib/couchdb
 chown -R couchdb:couchdb /usr/local/{lib,etc}/couchdb /usr/local/var/{lib,log,run}/couchdb
 chmod -R g+rw /usr/local/{lib,etc}/couchdb /usr/local/var/{lib,log,run}/couchdb

# Download & build sources
cd /tmp
wget http://apache.panu.it/couchdb/source/1.6.1/apache-couchdb-1.6.1.tar.gz
tar xzf apache-couchdb-1.6.1.tar.gz
cd apache-couchdb-1.6.1
./configure --prefix=/usr/local --with-js-lib=/usr/lib --with-js-include=/usr/include/js --enable-init
make &&  make install

# Finish setting up the environment
 chown couchdb:couchdb /usr/local/etc/couchdb/local.ini
 ln -s /usr/local/etc/init.d/couchdb /etc/init.d/couchdb
 ln -s /usr/local/etc/couchdb /etc
 update-rc.d couchdb defaults
 /etc/init.d/couchdb start

# Verify
echo "Waiting..."
sleep 5
echo "Test couchdb"
curl http://127.0.0.1:5984/
echo ""
echo ""
cp /etc/couchdb/local.ini /etc/couchdb/local_back.ini
sed -i 's/;bind_address.*/bind_address=0.0.0.0/' /etc/couchdb/local.ini
sed -i 's/;admin.*/admin = admin/' /etc/couchdb/local.ini
echo "************ DONE *********"
echo "You now need to edit /etc/couchdb/local.ini and add the following line to the httpd section"
echo "bind_adress 0.0.0.0"
echo "Then restart service with"
echo "etc/init.d/couchdb restart"
echo "Finally create an admin user in Futon"
