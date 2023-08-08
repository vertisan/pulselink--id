build:
	cd app && GOOS=linux GOARCH=amd64 go build -o ../bin/ && cd ..

proto.update:
	rm -f app/pkg/common/proto/* \
		&& curl -L -O -H 'Cache-Control: no-cache, no-store' https://github.com/vrsf-playground/grpc-proto/archive/refs/heads/master.zip \
		&& unzip -o master.zip && rm -rf master.zip \
		&& mv grpc-proto-master/proto/id.proto app/pkg/common/proto/ \
		&& rm -rf grpc-proto-master

proto.compile.id:
	protoc --go_out=./app/internal/id/pb/ --go-grpc_out=./app/internal/id/pb/ app/pkg/common/proto/id.proto
