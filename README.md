# tmitter

### installation + usage
You can start server in two ways:

#### run docker containers with predefined environment
1. [install docker engine](https://docs.docker.com/engine/install/)
2. [install docker-compose](https://docs.docker.com/compose/install/)
3. run in terminal ```sudo docker-compose up```
#### set up all environment on your machine
1. [install golang](https://go.dev/doc/install)
2. [set up mySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)
3. provide .env file with ``MYSQL_URI`` and ``APP_PORT`` variables
4. run in terminal ```go run .```
