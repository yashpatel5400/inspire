all: inspire.go quick.go random.go regular.go 
	make inspire random regular

inspire: inspire.go quick.go 
	go build inspire.go quick.go

random: random.go quick.go
	go build random.go quick.go

regular: regular.go quick.go
	go build regular.go quick.go

clean:
	rm random regular inspire