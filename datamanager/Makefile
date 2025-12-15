PROTOGEN_LOC=../protogen/golang/sensorpg
GATEWAY_PROTOGEN_LOC=../../gateway/Gateway/Protos/Reading.proto

protoc:
	cd proto && protoc --go_out=${PROTOGEN_LOC} --go_opt=paths=source_relative \
	--go-grpc_out=${PROTOGEN_LOC} --go-grpc_opt=paths=source_relative \
	*.proto && \
	cp *.proto ${GATEWAY_PROTOGEN_LOC}
