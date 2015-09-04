To download latest postgres image and start it:

docker run --name pg  -d postgres

Connect to this server with: localhost:5432 db=postgres user=postgres pwd=postgres

To stop container: docker stop pg

To start container: docker start pg

Docker postgres page: https://registry.hub.docker.com/_/postgres/
