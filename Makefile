run:
	docker compose run --rm back

test:
	docker compose run --rm back go test -v ./... | grep -e "--- FAIL" -e "ok" -B 2
