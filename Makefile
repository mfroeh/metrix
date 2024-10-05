.PHONY: run
run:
	npm --prefix frontend run dev && go run ./cmd/web
