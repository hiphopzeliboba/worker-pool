package workerpool

import (
	"fmt"
	"sync"
)

type Pool struct {
	workers   []*Worker   // Список существующих воркеров
	taskQueue chan string // Канал для задач
	mu        sync.Mutex  // Мьютекс для синхронизации доступа к воркерам
}

// NewWorkerPool создает новый экземпляр WorkerPool
func NewWorkerPool() *Pool {
	return &Pool{
		taskQueue: make(chan string),
	}
}

// AddWorker добавляет нового воркера в пул
func (wp *Pool) AddWorker() {
	wp.mu.Lock()         // Блокируем доступ к списку воркеров
	defer wp.mu.Unlock() // Анлочим

	workerId := len(wp.workers) + 1
	worker := NewWorker(workerId, wp)       // Создаем воркер
	wp.workers = append(wp.workers, worker) // Добавляем в список воркеров
	worker.Start()                          // Запуск воркера
}

// NewWorker создает нового воркера
func NewWorker(id int, pool *Pool) *Worker {
	return &Worker{
		id:     id,
		pool:   pool,
		stopCh: make(chan bool),
	}
}

// Start типа инициализирует пул воркеров
func (wp *Pool) Start() {
	wp.taskQueue = make(chan string) // Инициализируем канал для получения задач
}

// RemoveWorker удаляет последнего добавленного воркера из пула
func (wp *Pool) RemoveWorker() {
	wp.mu.Lock()         // Блокируем дочтуп к списку воркеров
	defer wp.mu.Unlock() // Разблокировка

	if len(wp.workers) > 0 { // Проверяем есть ли работающие воркеры
		worker := wp.workers[len(wp.workers)-1]     // Получаем ссылку на последнего добавленного воркера
		worker.stopCh <- true                       // Отправляем сигнал на остановку воркера
		close(worker.stopCh)                        // Закрываем канал для полного освобождения ресурсов
		wp.workers = wp.workers[:len(wp.workers)-1] // Удаляем воркер из списка
	}
}

// RemovePool принудительно завершает работу всех воркеров и очищает пул
func (wp *Pool) RemovePool() {
	wp.mu.Lock()         // Блокируем дочтуп к списку воркеров
	defer wp.mu.Unlock() // Разблокировка

	for _, worker := range wp.workers {
		worker.stopCh <- true // Отправляем сигнал на остановку воркера
		close(worker.stopCh)  // Закрываем канал для полного освобождения ресурсов
	}
	wp.workers = nil // Очищаем список воркеров
	fmt.Println("Pool successful cleared! ")
}

// AddTask отправляет задачу в очередь задач
func (wp *Pool) AddTask(task string) {
	wp.taskQueue <- task // Отправляем задачу в очередь
}
