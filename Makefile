.PHONY: all build test lint generate-schema clean install-skill install-cli docs docs-serve

all: build test lint

build:
	go build ./...

# Install dfspec CLI to $GOPATH/bin
install-cli:
	go install ./cmd/dfspec
	@echo "Installed dfspec to $(shell go env GOPATH)/bin/dfspec"

test:
	go test -v ./...

lint:
	golangci-lint run

generate-schema:
	go run cmd/generate-schema/main.go

clean:
	rm -f schemas/*.schema.json

# Install skills globally for Claude Code
install-skill: install-skill-dfspec-guide install-skill-dfspec-exec

install-skill-dfspec-guide:
	@mkdir -p ~/.claude/skills
	@if [ -L ~/.claude/skills/dfspec-guide ]; then \
		rm ~/.claude/skills/dfspec-guide; \
	fi
	ln -s $(PWD)/skills/dfspec-guide ~/.claude/skills/dfspec-guide
	@echo "Installed dfspec-guide skill to ~/.claude/skills/dfspec-guide"

install-skill-dfspec-exec:
	@mkdir -p ~/.claude/skills
	@if [ -L ~/.claude/skills/dfspec-exec ]; then \
		rm ~/.claude/skills/dfspec-exec; \
	fi
	ln -s $(PWD)/skills/dfspec-exec ~/.claude/skills/dfspec-exec
	@echo "Installed dfspec-exec skill to ~/.claude/skills/dfspec-exec"

# Uninstall skills
uninstall-skill:
	rm -f ~/.claude/skills/dfspec-guide
	rm -f ~/.claude/skills/dfspec-exec
	@echo "Uninstalled all skills"

# Regenerate everything
regenerate: clean generate-schema
	@echo "Schemas regenerated"

# Validate examples against schemas (requires ajv-cli)
validate-examples:
	@command -v ajv >/dev/null 2>&1 || { echo "ajv-cli required: npm install -g ajv-cli"; exit 1; }
	ajv validate -s schemas/constitution.schema.json -d examples/example-constitution.json
	ajv validate -s schemas/roadmap.schema.json -d examples/example-roadmap.json
	ajv validate -s schemas/mrd.schema.json -d examples/example-mrd.json
	ajv validate -s schemas/prd.schema.json -d examples/example-prd.json
	ajv validate -s schemas/evaluation.schema.json -d examples/example-evaluation.json

# Install MkDocs dependencies
docs-deps:
	pip install -r requirements.txt

# Build documentation site
docs:
	mkdocs build

# Serve documentation locally
docs-serve:
	mkdocs serve

# Deploy docs to GitHub Pages
docs-deploy:
	mkdocs gh-deploy
