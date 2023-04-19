 protoc --go_out=proto/pb --go_opt=paths=source_relative --go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative user.proto
# fiber-grpc
protoc --go_out=proto/redis --go_opt=paths=source_relative --go-grpc_out=proto/redis --go-grpc_opt=paths=source_relative redis.proto