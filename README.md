# go-bin
go binary template

This Project Contain
1) Migration Script
    spawn postgres db using :
    -- docker
    **https://www.datacamp.com/tutorial/postgresql-docker
    -- == network
    docker network create --subnet=172.18.0.0/16 mynetwork
    -- == postgres
    -- ==== create
    docker pull postgres
    docker run --name pg-db -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres

    -- ==== oprational imge
    docker start pg-db
    docker stop pg-db
    docker exec -it pg-db sh

    --net mynetwork --ip 172.18.0.3
    -- ==== check
    docker ps
    docker inspect 7aef7a6d5e26
    docker inspect 1e3e2dad4766 | findstr IPAddress

2) Binary Apps