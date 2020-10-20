cd pbfiles
protoc --go_out=plugins=grpc:../services models.proto
protoc --go_out=plugins=grpc:../services prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services prod.proto
protoc --go_out=plugins=grpc:../services order.proto
protoc --grpc-gateway_out=logtostderr=true:../services order.proto
protoc --go_out=plugins=grpc:../services --validate_out=lang=go:../services models.proto
protoc --go_out=plugins=grpc:../services user.proto
protoc --grpc-gateway_out=logtostderr=true:../services user.proto