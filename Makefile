all:
	@echo "**********************************************************"
	@echo "**                      Makefile                        **"
	@echo "**********************************************************"
	@echo ""
	@echo "Available commands:"
	@echo ""
	@echo "  test-go        - Watch Go files and run tests automatically"
	@echo "  test-go-run    - Run Go tests once"
	@echo "  test-ts        - Watch TypeScript files and run tests automatically" 
	@echo "  test-ts-run    - Run TypeScript tests once"
	@echo ""
	@echo "Usage: make <command>"
	@echo ""

test-go:
	rg --files -t go ./go | entr -c go test -C go ./problems/... -v

test-go-run:
	go test -C go ./problems/... -v

test-ts:
	pnpm -C ts test

test-ts-run:
	pnpm -C ts test:run
