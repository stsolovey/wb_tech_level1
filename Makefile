t1:
	go run ./task01/main.go

t2-1:
	@echo "Конкурентное вычисление с использованием канала"
	go run ./task02/solution1/main.go

t2-2:
	@echo "Конкурентное вычисление с сохранением результатов в срез"
	go run ./task02/solution2/main.go

t2-3:
	@echo "Работа с каналами и структурами для заданий"
	go run ./task02/solution3/main.go

t2-4:
	@echo "Использование 'sync.Map' для хранения результатов"
	go run ./task02/solution4/main.go

t2-5:
	@echo "Паттерн Fan-In для сбора результатов в один канал"
	go run ./task02/solution4/main.go

t2-6:
	@echo "Использование пула горутин с ограниченным размером"
	go run ./task02/solution4/main.go

t2-7:
	@echo "Использование пакета 'reflect.Select'"
	go run ./task02/solution4/main.go

t2-8:
	@echo "Использование атомарных операций"
	go run ./task02/solution4/main.go

t3-1:
	@echo "Использование каналов для вычисления и суммирования"
	go run ./task03/solution1/main.go

t3-2:
	@echo "Использование атомарных операций для суммирования квадратов"
	go run ./task03/solution2/main.go

t4-1:
	@echo "Использование контекста"
	go run ./task04/solution1/main.go

t4-2:
	@echo "Использование канала 'done'"
	go run ./task04/solution2/main.go

t4-3:
	@echo "Использование sync.WaitGroup"
	go run ./task04/solution3/main.go

t5-1:
	@echo "таймер time.After + select"
	go run ./task05/solution1/main.go

t5-2:
	@echo "таймер time.NewTimer + select"
	go run ./task05/solution2/main.go

t5-3:
	@echo "Использование контекста с таймаутом"
	go run ./task05/solution3/main.go

t5-4:
	@echo "Использование системных сигналов с таймером"
	go run ./task05/solution4/main.go

t6:
	@echo "Различные способы остановки горутины"
	go run ./task06/main.go

t7-1:
	@echo "Конкурентная запись данных в map. sync.Map"
	go run ./task07/solution1/main.go

t7-2:
	@echo "Конкурентная запись данных в map. Каналы"
	go run ./task07/solution2/main.go

t7-3:
	@echo "Конкурентная запись данных в map. sync.Mutex"
	go run ./task07/solution3/main.go

t7-4:
	@echo "Конкурентная запись данных в map. sync.RWMutex"
	go run ./task07/solution4/main.go

t8:
	@echo "8 Изменение i-го бита переменной"
	go run ./task08/main.go

t9:
	@echo "9 Конвеер чисел"
	go run ./task09/main.go

t10:
	@echo "10 Группировка температур"
	go run ./task10/main.go

t11:
	@echo "11 Пересечение множеств"
	go run ./task11/main.go

t12:
	@echo "12 Реализация множества"
	go run ./task12/main.go

t13:
	@echo "13 swap without temp var"
	go run ./task13/main.go

t14:
	@echo "14 определение типа переменной"
	go run ./task14/main.go

t15:
	@echo "15 негативные последствия слайсинга"
	go run ./task15/main.go

t16:
	@echo "16 quicksort"
	go run ./task16/main.go

t17:
	@echo "17 бинарный поиск"
	go run ./task17/main.go

t18-1:
	@echo "18 Счётчик sync/atomic"
	go run ./task18/solution1/main.go

t18-2:
	@echo "18 Счётчик sync мьютексы"
	go run ./task18/solution2/main.go

t18-3:
	@echo "18 Счётчик с каналами"
	go run ./task18/solution3/main.go

t19:
	@echo "19 reverse sting"
	go run ./task19/main.go

t20:
	@echo "20 reverse sentence"
	go run ./task20/main.go

t21:
	@echo "21 паттерн адаптер"
	go run ./task21/main.go

t22:
	@echo "22 большие числа"
	go run ./task22/main.go

t23-1:
	@echo "23-1 удалить элемент слайса"
	go run ./task23/solution1/main.go

t23-2:
	@echo "23-2 удалить элемент слайса"
	go run ./task23/solution2/main.go

t23-3:
	@echo "23-3 удалить элемент слайса"
	go run ./task23/solution3/main.go

t24:
	@echo "24 расстояние между точками"
	go run ./task24/main.go

t25:
	@echo "25 sleep"
	go run ./task25/main.go

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
