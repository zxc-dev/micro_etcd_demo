TARGETS=login_api login_service

docker:
	$(foreach target, ${TARGETS},\
		docker build . -t "${target}:latest" -f ./${target}/Dockerfile ;\
	)

build:
	go build -o ./bin/login_service ./login_service