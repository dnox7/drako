.PHONY: lint
lint:
	buf format -w
	buf lint
	buf breaking

.PHONY: gen
gen:
	buf generate && cargo build --manifest-path ./gen/rust/Cargo.toml
