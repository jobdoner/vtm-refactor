regen:
	rm -rf  restapi/operations/  && swagger generate server -f swagger.yaml --exclude-main -A vtm-refactor
#TODO:
# addhandler:
# 	make go ast script with handlers as targer
# 	go run tooling/main.go $(ARGS)