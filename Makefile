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
	@echo "25 custom sleep"
	go run ./task25/main.go

t26:
	@echo "26 check string uniqueness"
	go run ./task26/main.go

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
	@echo "  tidy                 - Format and tidy up the Go code using gofumpt and gci"
	@echo "  lint                 - Lint the project code and ensure it adheres to best practices"
	@echo "  tools                - Install or update necessary Go tools"
	@echo "  t1                   - Run the task related to structure embedding and inheritance simulation"
	@echo "  t2-1 to t2-8         - Run different solutions for concurrent calculation of squares from an array"
	@echo "  t3-1, t3-2           - Run solutions for concurrent calculation of the sum of squares"
	@echo "  t4-1 to t4-3         - Run different approaches for implementing concurrent workers reading from a channel"
	@echo "  t5-1 to t5-4         - Run solutions to implement program termination after N seconds using various techniques"
	@echo "  t6                   - Demonstrate different methods to stop a goroutine"
	@echo "  t7-1 to t7-4         - Implement and compare methods for concurrent map writing"
	@echo "  t8                   - Run the task to modify a specific bit in an int64 variable"
	@echo "  t9                   - Execute the number conveyor task"
	@echo "  t10                  - Group temperature fluctuations into buckets of 10 degrees"
	@echo "  t11                  - Find the intersection of two sets"
	@echo "  t12                  - Create a set from a sequence of strings"
	@echo "  t13                  - Swap two numbers without using a temporary variable"
	@echo "  t14                  - Determine the type of a variable at runtime from an interface{}"
	@echo "  t15                  - Discuss the drawbacks of slicing in Go and provide a safe alternative"
	@echo "  t16                  - Implement quicksort"
	@echo "  t17                  - Implement binary search"
	@echo "  t18-1 to t18-3       - Implement and compare methods for a concurrent counter"
	@echo "  t19                  - Reverse a unicode string"
	@echo "  t20                  - Reverse the words in a sentence"
	@echo "  t21                  - Implement the adapter pattern"
	@echo "  t22                  - Perform arithmetic operations on large numbers"
	@echo "  t23-1 to t23-3       - Implement different methods to remove an element from a slice"
	@would_be "  t24                  - Calculate the distance between two points represented as structures"
	@echo "  t25                  - Implement a custom sleep function"
	@echo "  t26                  - Check if all characters in a string are unique and case-insensitive"

