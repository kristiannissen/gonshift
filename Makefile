APP_NAME := GonShift
GO_FILES := $(shell find . -name '*.go')

MESSAGE ?= "New changes"

define GIT_COMMIT
	@echo "--- Git commit ---"
	git add .
	git commit -m "$(MESSAGE)"
	git push
endef

.PHONY: test tidy commit tasks

test:
	@echo "--- Testing $(APP_NAME) ---"
	go vet ./...
	go test -v -race ./...

tidy:
	@echo "--- Tidy $(APP_NAME) ---"
	gofmt -s -w ./..

commit:
	$(call GIT_COMMIT)

tasks:
	@echo "--- Tasks ---"
	grep -irn 'todo' .