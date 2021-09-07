moduleCompile:
	cd ./protos && protoc --go_out=./module --go_opt=paths=source_relative --go-grpc_out=./module --go-grpc_opt=paths=source_relative ./Module.proto

monitorCompile:
	cd ./protos && protoc --go_out=./monitor --go_opt=paths=source_relative --go-grpc_out=./monitor --go-grpc_opt=paths=source_relative ./Monitor.proto ./MonitorController.proto

dbCompile:
	cd ./sonic/database && go generate .

compileEnt:
	set REDIS_URL=rediss://default:n6luoc78ac44pgs0@test-redis-do-user-9223163-0.b.db.ondigitalocean.com:25061 && cd ./sonic/database && go generate ./ent

monitorControllerCompile:
	cd ./protos && protoc --go_out=./monitorController --go_opt=paths=source_relative --go-grpc_out=./monitorController --go-grpc_opt=paths=source_relative ./MonitorController.proto

shapeCompile:
	cd sonic/antibots/shape && protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative ./Shape.proto --experimental_allow_proto3_optional

taskControllerCompile:
	cd ./protos && protoc --go_out=./taskController --go_opt=paths=source_relative --go-grpc_out=./taskController --go-grpc_opt=paths=source_relative ./Module.proto ./TasksController.proto

captchaCompile:
	cd ./protos && protoc --go_out=./captcha --go_opt=paths=source_relative --go-grpc_out=./captcha --go-grpc_opt=paths=source_relative ./Captcha.proto

proxyRaterCompile:
		cd ./protos && protoc --go_out=./proxy-rater --go_opt=paths=source_relative --go-grpc_out=./proxy-rater --go-grpc_opt=paths=source_relative ./ProxyRater.proto

clientProxyCompile:
		cd ./protos && protoc --go_out=./clientProxy --go_opt=paths=source_relative --go-grpc_out=./clientProxy --go-grpc_opt=paths=source_relative ./ClientProxy.proto

ticketCompile:
	cd sonic/antibots/ticket && protoc --go_out=./protos --go_opt=paths=source_relative --go-grpc_out=./protos --go-grpc_opt=paths=source_relative ./Ticket.proto

gqlCompile:
	go get github.com/99designs/gqlgen/cmd@v0.13.0
	cd ./sonic/database && go run github.com/99designs/gqlgen gen