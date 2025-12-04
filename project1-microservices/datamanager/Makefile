protoc:
	# TAB → neki editori automatski pretvaraju u spaces, proveri!
	protoc --proto_path=proto \
		--go_out=protogen/golang --go_opt=paths=source_relative \
		--go-grpc_out=protogen/golang --go-grpc_opt=paths=source_relative \
		proto/*.proto