moduleCompile:
	cd ./protos && protoc --go_out=./module --go_opt=paths=source_relative --go-grpc_out=./module --go-grpc_opt=paths=source_relative ./Module.proto

monitorCompile:
	cd ./protos && protoc --go_out=./monitor --go_opt=paths=source_relative --go-grpc_out=./monitor --go-grpc_opt=paths=source_relative ./Monitor.proto

dbCompile:
	cd ./sonic/database && go generate .

compileEnt:
	set REDIS_URL=rediss://default:n6luoc78ac44pgs0@test-redis-do-user-9223163-0.b.db.ondigitalocean.com:25061 && cd ./sonic/database && go generate ./ent

monitorControllerCompile:
	cd ./protos && protoc --go_out=./monitorController --go_opt=paths=source_relative --go-grpc_out=./monitorController --go-grpc_opt=paths=source_relative ./MonitorController.proto

shapeCompile:
	cd sonic/antibots/shape && protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative ./Shape.proto

taskControllerCompile:
	cd ./protos && protoc --go_out=./taskController --go_opt=paths=source_relative --go-grpc_out=./taskController --go-grpc_opt=paths=source_relative ./Module.proto ./TasksController.proto