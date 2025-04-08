#Host 192.168.182.100
#Port 5434

pg_basebackup -h localhost --port=5433 -U repuser --checkpoint=fast -D /tmp/replica_db -R -C --slot=slot_name

#Change in postgresql.conf
listen_addresses = '*'
port = 5434

#Change in pg_hba.conf