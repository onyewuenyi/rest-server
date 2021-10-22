# rest-server

# in powershell download scoope https://scoop.sh/ w/ iwr -useb get.scoop.sh | iex
scoop install migrate
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

migrate create -ext sql -dir db/migration -seq init_schema


# contains running
docker ps 
# containers 
docker ps -a

docker start <container name>
docker stop <container name>

docker run \
        --name <container name> \
        -e POSTGRES_USER=${DB_USER} \
        -e POSTGRES_PASSWORD=${DB_PASSWORD} \
        -p "${DB_PORT}":5432 \
        -d <image name e.g. postgres> \
        postgres -N 1000


# windows dependency to run docker exec cmd 
# Linux docker: exec -it taskstore psql -U ${DB_USER} ${DB_NAME}
winpty docker exec -it taskstore <cmd>
e.g. winpty docker exec -it <db name> psql -U ${DB_USER} ${DB_NAME}

winpty docker exec -it postgres12 psql -U charlesonyewuenyi postgres12
 
dont have bash shell in ubuntu so use bin/sh shell 

docker exec -it <container name> bash



# default owner is charlesonyewuenyi
# defautl username val is root
createdb --username=charlesonyewuenyi taskstore
dropdb --username=charlesonyewuenyi taskstore


make <CMD in makefile>