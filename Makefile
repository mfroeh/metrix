.PHONY: run
run:
	npm --prefix frontend run dev && go run ./cmd/web -port 8080 -riot-api-key ${RGAPI}
