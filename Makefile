##### Convenience targets ######

REPO:=github.com/tikasan/goa-simple-sample

init: depend bootstrap
gen: clean generate
rerun: clean generate run
rerundb: clean generate rundb

depend:
	@which goagen || go get -v github.com/goadesign/goa/goagen
	@go get -v ./...

bootstrap:
	@goagen bootstrap -d $(REPO)/design

main:
	@goagen main -d $(REPO)/design

js:
	@goagen js -d $(REPO)/design

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger
	@rm -rf schema
	@rm -rf js
	@rm -f build

generate:
	@goagen app     -d $(REPO)/design
	@goagen swagger -d $(REPO)/design
	@goagen client -d $(REPO)/design
	@goagen js -d $(REPO)/design
	@goagen schema -d $(REPO)/design
	@go build -o build

model:
	@ls models | grep -v '_defined.go' | xargs rm -f
	@goagen --design=$(REPO)/design gen --pkg-path=github.com/goadesign/gorma

swaggerUI:
	@open http://localhost:8080/swagger/index.html

run:
	@go run main.go

rundb:
	@go run main.go --dbrun

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

##### Database ######

DBNAME:=celler
ENV:=development

migrate/init:
	mysql -u celler -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/up:
	sql-migrate up -env=$(ENV)

migrate/down:
	sql-migrate down -env=$(ENV)

migrate/status:
	sql-migrate status -env=$(ENV)

migrate/dry:
	sql-migrate up -dryrun -env=$(ENV)

##### Docker ######

docker/build: Dockerfile docker-compose.yml
	docker-compose build

docker/start:
	docker-compose up -d

docker/stop:
	docker-compose down

docker/logs:
	docker-compose logs

docker/clean:
	docker-compose rm

docker/ssh:
	docker exec -it celler /bin/bash

##### Curls ######

BASEURL:=http://localhost:8080/api/v1

## action ##

curl_action_id:
	curl -v '$(BASEURL)/actions/$(ID)'

curl_action_ping:
	curl -v '$(BASEURL)/actions/ping'

curl_action_hello:
	curl -v '$(BASEURL)/actions/hello?name=$(name)'

## validation ##

curl_validation:
	curl -v '$(BASEURL)/validation?ID=1&defaultType=&email=satak%40gmail.com&enumType=A&integerType=10&stringType=foo&reg=12abc'

## response ##

curl_response_users:
	curl -v '$(BASEURL)/response/users'

curl_response_users_id:
	curl -v '$(BASEURL)/response/users/1'

curl_response_users_array:
	curl -v '$(BASEURL)/response/users/array'

curl_response_users_hash:
	curl -v '$(BASEURL)/response/users/hash'

## method ##

curl_method_follow:
	curl -v -X PUT '$(BASEURL)/method/users/follow'

curl_method_unfollow:
	curl -v -X DELETE '$(BASEURL)/method/users/follow'

curl_method_get:
	curl -v '$(BASEURL)/method/get'

curl_method_post:
	curl -v -X POST '$(BASEURL)/method/post'

curl_method_put:
	curl -v -X PUT '$(BASEURL)/method/put'

curl_method_delete:
	curl -v -X DELETE '$(BASEURL)/method/delete'

curl_method_etc:
	curl -v '$(BASEURL)/method/users/1/follow/3'
