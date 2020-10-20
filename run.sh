cd pbfiles
protoc --go_out=plugins=grpc:../services prod.proto
protoc --go_out=plugins=grpc:../services models.proto
protoc --go_out=plugins=grpc:../services order.proto
protoc --grpc-gateway_out=logtostderr=true:../services prod.proto