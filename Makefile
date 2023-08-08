build:
	cd app && GOOS=linux GOARCH=amd64 go build -o ../bin/ && cd ..

proto.update:
	rm -rf app/pkg/common/proto \
		&& curl -L -O -H 'Cache-Control: no-cache, no-store' https://github.com/vrsf-playground/grpc-proto/archive/refs/heads/master.zip \
		&& unzip master.zip && rm -rf master.zip \
		&& mv grpc-proto-master/proto app/pkg/common/ \
		&& rm -rf grpc-proto-master

proto.compile.id:
	protoc --go_out=./app/pkg/domain/id/pb/ app/pkg/common/proto/collect.proto
