##### Convenience targets ######

REPO:=github.com/tikasan/goa-simple-sample

init: depend bootstrap
gen: clean generate
rerun: clean generate run

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

##### Appengine targets ######

#init: depend bootstrap appengine
#gen: clean generate appengine
#rerun: clean generate appengine run

appengine:
	@which gorep || go get -v github.com/novalagung/gorep
	@gorep -path="./vendor/github.com/goadesign/goa" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./models" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./app" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./client" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./tool" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./" \
          -from="../app" \
          -to="$(REPO)/app"
	@gorep -path="./" \
          -from="../client" \
          -to="$(REPO)/client"
	@gorep -path="./" \
          -from="../tool/cli" \
          -to="$(REPO)/tool/cli"

##### Curls ######

BASEURL:=http://localhost:8080/api/v1

curl_action_id:
	curl -v $(BASEURL)/actions/$(id)

curl_action_main:
	curl -v $(BASEURL)/actions/main

curl_action_sub:
	curl -v $(BASEURL)/actions/sub?name=$(name)