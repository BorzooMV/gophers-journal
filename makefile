SEED_EXAMPLE_SCRIPT_SRC=cmd/scripts/seed-db-example/seed-db-example.go
CLEAN_DB_SCRIPT_SRC=cmd/scripts/clean-db/clean-db.go

clean-db:
	go run ${CLEAN_DB_SCRIPT_SRC}

seed-db-example: clean-db
	go run ${SEED_EXAMPLE_SCRIPT_SRC}
