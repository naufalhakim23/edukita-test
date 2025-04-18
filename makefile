.PHONY: start
start:
	air -c ./.air.toml
build-air:
	go build -o ./tmp/main ./


create-migration:
	@read -p "Enter Migration Name: " migration; \
	dir=./migrations; \
	mkdir -p $$dir; \
	timestamp=$$(date +%s); \
	touch $$dir/$$timestamp\_$$migration.up.sql; \
	touch $$dir/$$timestamp\_$$migration.down.sql; \
	chmod -R a+rwX $$dir; \
	echo "Created migration files in $$dir"

db-version:
	go run migrations/scripts/main.go -version

migrate-up:
	go run migrations/scripts/main.go -up

migrate-down:
	go run migrations/scripts/main.go -down