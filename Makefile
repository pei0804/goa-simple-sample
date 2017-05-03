##### Convenience targets ######

REPO:=github.com/tikasan/goa-simple-sample

init: depend bootstrap
gen: clean generate

depend:
	go get -v ./...

bootstrap:
	@goagen bootstrap -d $(REPO)/design

main:
	@goagen main -d $(REPO)/design

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger

generate:
	@goagen app     -d $(REPO)/design
	@goagen swagger -d $(REPO)/design
	@goagen client -d $(REPO)/design

swaggerUI:
	@open http://localhost:8080/swagger/index.html

run:
	@go run main.go
