#Host 192.168.182.100
#Port 5433

initdb -D /tmp/primary_db/

#Change in postgresql.conf
listen_addresses = '*'
port = 5433

#Create replication user before create replica db
psql --port=5433 postgres
create user repuser replication;

#Change in pg_hba.conf
# IPv4 local connections:
host    all             repuser             127.0.0.1/32            trust