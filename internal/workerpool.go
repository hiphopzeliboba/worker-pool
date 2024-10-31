package internal

// WorkerPool определяет методы для пула воркеров и сокрывает реализации
type WorkerPool interface {
	AddWorker()          // Добавляем воркер в пул
	RemoveWorker()       // Удаляем врокер из пула
	AddTask(task string) // Добавляем задачу в очередь
	RemovePool()         // Очищаем весь пул
}
