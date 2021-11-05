FROM bms-proto-builder:latest AS builder
WORKDIR /proto

# Copy our repo
COPY . ./

# Setup general includes
ENV PROTO_INC "-I ./ \
  -I ../ \
  -I ../../ \
  -I $GOPATH/src \
  -I $GOPATH/pkg/mod"

ENV PROTOC_CMD "protoc ${PROTO_INC} --go_out=. --plugin protoc-gen-go=${GOPATH}/bin/protoc-gen-go --go-grpc_out=. --plugin protoc-gen-go-grpc=${GOPATH}/bin/protoc-gen-go-grpc --go-vtproto_out=. --plugin protoc-gen-go-vtproto=${GOPATH}/bin/protoc-gen-go-vtproto --go-vtproto_opt=features=marshal+unmarshal+size --go-drpc_out=. --plugin protoc-gen-go-drpc=${GOPATH}/bin/protoc-gen-go-drpc --go-drpc_opt=protolib=github.com/planetscale/vtprotobuf/codec/drpc ./*.proto"
# Generate
# RUN ls -R
RUN cd /proto/ && ${PROTOC_CMD}
# RUN ls -R

CMD ["/bin/sh", "-c", "echo Docker done"]