.PHONY: postgres adminer migrate
# Run postgresdb in the background and when it finishes it cleans everything. The --network make the db available to the host, 
# so no need to specify port (localhost) and the last flag is to set the environment variable for the password.
postgres:
	docker run --rm -ti -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti -p 8080:8080 adminer

migrate:
	migrate -source file://migrations -database postgres://postgres:secret@192.168.1.38/postgres?sslmode=disable up