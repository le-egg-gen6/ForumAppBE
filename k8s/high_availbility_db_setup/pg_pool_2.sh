#Host 192.168.182.100
#Port 9999

mkdir "/pg_pool"

#Change the pgpool.conf
master_slave_mode = on
master_slave_sub_mode = 'stream'
load_balance_mode = on


backend_hostname0 = 'localhost'
                                   # Host name or IP address to connect to for backend 0
backend_port0 = 5433
                                   # Port number for backend 0
backend_weight0 = 0
                                   # Weight for backend 0 (only in load balancing mode)
backend_data_directory0 = '/tmp/primary_db/'
                                   # Data directory for backend 0
backend_flag0 = 'ALLOW_TO_FAILOVER'
                                   # Controls various backend behavior
                                   # ALLOW_TO_FAILOVER, DISALLOW_TO_FAILOVER
                                   # or ALWAYS_MASTER
backend_application_name0 = 'server0'
                                   # walsender's application_name, used for "show pool_nodes" command
backend_hostname1 = 'localhost'
backend_port1 = 5434
backend_weight1 = 2
backend_data_directory1 = '/tmp/replica_db'
backend_flag1 = 'ALLOW_TO_FAILOVER'
backend_application_name1 = 'server1'

health_check_period = 10
health_check_user = 'repuser'

pgpool -n