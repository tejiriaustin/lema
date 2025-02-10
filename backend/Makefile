.PHONY: app mocks test
api:
	go run main.go api

rm-mocks:
	rm -rf ./testutils/mocks.*

gen-mocks:
	mockery --all --output=testutils/mocks --case=underscore --keeptree

mocks: rm-mocks gen-mocks

test:
	go test -v -coverprofile=cover.out.tmp -coverpkg=./... ./...

deploy:
    docker build -t lema .
    docker save lema > lema.tar
    scp lema.tar ec2-user@ec2-16-171-143-117.eu-north-1.compute.amazonaws.com:/home/ec2-user/
    ssh ec2-user@ec2-16-171-143-117.eu-north-1.compute.amazonaws.com "docker load < lema.tar && docker run -d -p 8080:8080 lema"
