<<<<<<< HEAD
=======

>>>>>>> 56c422e032c578a4b71d7cb35be680dbc95ce8b5

## mysql 环境配置

```sql
ubuntu@main:~$ sudo docker pull mysql:latest

ubuntu@main:~$ sudo docker run -itd --name master -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql

ubuntu@main:~$ sudo docker run -itd --name slave -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql


ubuntu@main:~$ sudo docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED              STATUS              PORTS                                                  NAMES
fc0df1f056eb   mysql     "docker-entrypoint.s…"   37 seconds ago       Up 35 seconds       33060/tcp, 0.0.0.0:3307->3306/tcp, :::3307->3306/tcp   slave
24ea9465c249   mysql     "docker-entrypoint.s…"   About a minute ago   Up About a minute   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   master


ubuntu@main:~$ sudo netstat -antup | grep docker
tcp        0      0 0.0.0.0:3306            0.0.0.0:*               LISTEN      13031/docker-proxy
tcp        0      0 0.0.0.0:3307            0.0.0.0:*               LISTEN      13291/docker-proxy
tcp6       0      0 :::3306                 :::*                    LISTEN      13037/docker-proxy
tcp6       0      0 :::3307                 :::*                    LISTEN      13296/docker-proxy


ubuntu@main:~$ sudo apt install mysql-client-core-8.0


ubuntu@main:~$ mysql -uroot -h127.0.0.1 -P3306 -p123456

ubuntu@main:~$ mysql -uroot -h127.0.0.1 -P3307 -p123456


```
## 3306 主

mysql> show master status;
+-------------------+----------+--------------+------------------+-------------------+
| File              | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+-------------------+----------+--------------+------------------+-------------------+
| master-bin.000001 |      156 |              |                  |                   |
+-------------------+----------+--------------+------------------+-------------------+
1 row in set (0.05 sec)




ubuntu@main:~$ sudo docker inspect --format='{{.NetworkSettings.IPAddress}}' master
172.17.0.2
ubuntu@main:~$ sudo docker inspect --format='{{.NetworkSettings.IPAddress}}' slave
172.17.0.3



CHANGE MASTER TO
  MASTER_HOST='192.168.64.2',
  MASTER_USER='root',
  MASTER_PASSWORD='123456',
  MASTER_PORT=3307,
  MASTER_LOG_FILE='binlog.000002',
  MASTER_LOG_POS=675;







mysql> start slave;

mysql> show slave status\G;
*************************** 1. row ***************************
               Slave_IO_State: Waiting for source to send event
                  Master_Host: 172.17.0.2
                  Master_User: root
                  Master_Port: 3307
                Connect_Retry: 60
              Master_Log_File: master-bin.000002
          Read_Master_Log_Pos: 157
               Relay_Log_File: mysql-relay-bin.000002
                Relay_Log_Pos: 325
        Relay_Master_Log_File: master-bin.000001
             Slave_IO_Running: Yes
            Slave_SQL_Running: Yes






## 测试 主

mysql> create database test1 default character set utf8;
Query OK, 1 row affected, 1 warning (0.11 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| test1              |
+--------------------+
5 rows in set (0.10 sec)

## 测试 从

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| test1              |
+--------------------+
5 rows in set (0.18 sec)



```

## V4.0 基于Hash计算实现秒传 （大幅提升拥有海量文件的云盘性能）


```go
handler/handler.go 实现：

FileQueryHandler : 查询批量的文件元信息
TryFastUploadHandler : 尝试秒传接口
DownloadHandler : 文件下载接口


db/userfile.go 实现：

QueryUserFileMetas : 批量获取用户文件信息

meta/filemeta.go 实现：

GetLastFileMetas : 获取批量的文件元信息列表
GetLastFileMetasDB : 批量从mysql获取文件元信息



```