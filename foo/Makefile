generate:
	protoc-b2bchain1234 -I=. \
		-I=../vendor \
		-I=../third_party/googleapis \
		-I=../  \
		--go-b2bchain_out=plugins=grpc:. \
		--grpc-gateway-b2bchain_out=logtostderr=true:. \
		--swagger-b2bchain_out=logtostderr=true:. \
		./*.proto