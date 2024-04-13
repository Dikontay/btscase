run:
	go run ./cmd/api/

postgres_docker_download:
	docker pull postgres:alpine

postgres_run:
	docker run -itd -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -v /data:/var/lib/postgresql/data --name postgresql postgres:alpine
	
# postgres_create_db:
# 	docker exec -it postgresql bash
# 	psql -U postgres
# 	create database ozinse;