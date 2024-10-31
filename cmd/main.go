package main

import (
	"fmt"
	"time"
	"worker-pool/internal"
	"worker-pool/internal/workerpool"
)

func main() {
	var pool internal.WorkerPool = workerpool.NewWorkerPool()

	workersNumber := 3 // Задаем кол-во воркеров

	// Добавляем несколько воркеров в пул
	for i := 0; i < workersNumber; i++ {
		pool.AddWorker()
	}

	// Отправляем 20 задач в пул
	for i := 0; i < 20; i++ {
		pool.AddTask(fmt.Sprintf("Task %d", i))
		time.Sleep(120 * time.Millisecond) // Имитиция работы
	}

	pool.RemoveWorker() // Удалили одного воркера
	pool.AddWorker()    // Добавили одного воркера
	pool.AddWorker()    // Добавили одного воркера

	// Отправляем еще 20 задач в пул
	for i := 100; i < 120; i++ {
		pool.AddTask(fmt.Sprintf("Task %d", i))
		time.Sleep(300 * time.Millisecond) // Имитация работы
	}

	pool.RemovePool() // Очищаем весь пул перед завершенеим

}
