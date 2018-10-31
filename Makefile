build: 
	go build -o webServer main.go router.go
run:
	nohup ./webServer > log.out 2>&1 &
clean:
	rm webServer log.out
