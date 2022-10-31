TARGETS=login_api login_service

docker:
	$(foreach target, ${TARGETS},\
		docker build -t "${target}:latest";\
	)
