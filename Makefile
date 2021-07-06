protoCompile:
	protoc --go_out=./protos --go_opt=paths=source_relative --go-grpc_out=./protos --go-grpc_opt=paths=source_relative ./Module.proto

monitorCompile:
	protoc --go_out=./monitors --go_opt=paths=source_relative --go-grpc_out=./monitors --go-grpc_opt=paths=source_relative ./Monitor.proto

compileEnt:
	set REDIS_URL=rediss://default:n6luoc78ac44pgs0@test-redis-do-user-9223163-0.b.db.ondigitalocean.com:25061 && cd ./sonic/database && go generate ./ent