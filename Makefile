.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/coupon coupon/main.go
	# env GOOS=linux go build -ldflags="-s -w" -o bin/listCoupons listCoupons/main.go
	# env GOOS=linux go build -ldflags="-s -w" -o bin/createItem createItem/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
