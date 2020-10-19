# grpc-demo
## protoc生成.go文件
create go file
```bash
cd pbfiles
protoc --go_out=../services prod.proto
```
update go file
```bash
protoc --go_out=plugins=grpc:../services prod.proto
```
## 生成自签证书
### 单向认证
用如下方式生成证书然后将server.crt和server_no_passwd.key放入server/keys下，server.crt放入client/keys下，用于访问验证
```bash
genrsa -des3 out server.key 2048 // 会生成server.key私钥文件
req -new -key server.key -out server.csr // 会生成server.csr
// 其中common name也就死域名：grpcserver.com
rsa -in server.key -out server_no_passwd.key // 删除密码
x509 -req -days 365 -in server.csr -signkey server_no_passwd.key -out server.crt // 生成server.crt
```
### 双向认证
生成根证书
```bash
genrsa -out ca.key 2048
req -new -x509 -days 3650 -key ca.key -out ca.pem
```
生成服务端证书
```bash
genrsa -out server.key 2048
req -new -key server.key -out server.csr // 本地使用localhost作为域名
x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem
```
服务端代码
```go
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	cred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},        // 服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert, // 双向验证
		ClientCAs:    certPool,
	})

	rpcServer := grpc.NewServer(grpc.Creds(cred))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	rpcServer.Serve(lis)
```
生成客户端
```bash
ecparam -genkey -name secp384r1 -out client.key
req -new -key client.key -out client.csr
x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem
```
客户端代码
```go
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	cred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, // 客户端证书
		ServerName:   "localhost",             // 域名
		RootCAs:      certPool,
	})

```


```bash
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
```