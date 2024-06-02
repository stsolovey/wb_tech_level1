t1:
	go run ./task1/main.go

t2-1:
	@echo "Конкурентное вычисление с использованием канала"
	go run ./task2/solution1/main.go

t2-2:
	@echo "Конкурентное вычисление с сохранением результатов в срез"
	go run ./task2/solution2/main.go

t2-3:
	@echo "Работа с каналами и структурами для заданий"
	go run ./task2/solution3/main.go

t2-4:
	@echo "Использование 'sync.Map' для хранения результатов"
	go run ./task2/solution4/main.go

t2-5:
	@echo "Паттерн Fan-In для сбора результатов в один канал"
	go run ./task2/solution4/main.go

t2-6:
	@echo "Использование пула горутин с ограниченным размером"
	go run ./task2/solution4/main.go

t2-7:
	@echo "Использование пакета 'reflect.Select'"
	go run ./task2/solution4/main.go

t2-8:
	@echo "Использование атомарных операций"
	go run ./task2/solution4/main.go

t3-1:
	@echo ""
	go run ./task3/solution1/main.go

t3-2:
	@echo ""
	go run ./task3/solution2/main.go

t4-1:
	@echo ""
	go run ./task4/solution1/main.go 3

t4-2:
	@echo ""
	go run ./task4/solution2/main.go 3

t4-3:
	@echo ""
	go run ./task4/solution3/main.go 3

tidy:
	gofumpt -w .
	gci write . --skip-generated -s standard -s default
	go mod tidy

lint: tidy
	golangci-lint run ./...

tools:
	go install mvdan.cc/gofumpt@latest
	go install github.com/daixiang0/gci@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest


help:
	@echo "Available commands:"
	@echo "  tidy                 - Format and tidy up the Go code"
	@echo "  lint                 - Lint and format the project code"
	@echo "  tools                - Install necessary tools"
	@echo "  t1                   - Execute task 1"
	@echo "  t2-1                 - Execute task 2 solution 1"
	@echo "  t2-2                 - Execute task 2 solution 2"
	@echo "  t2-3                 - Execute task 2 solution 3"
	@echo "  t2-4                 - Execute task 2 solution 4"
	@echo "  t2-5                 - Execute task 2 solution 5"
	@echo "  t2-5                 - Execute task 2 solution 6"
	@echo "  t2-5                 - Execute task 2 solution 7"
	@echo "  t2-5                 - Execute task 2 solution 8"
