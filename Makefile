test:
	docker compose run --rm boilerplate go test -v ./... | grep -e "--- FAIL" -e "ok" -B 2
