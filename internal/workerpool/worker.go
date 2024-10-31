package workerpool

import (
	"fmt"
)

// Worker представляет отдельного воркера
type Worker struct {
	id     int       // Воркер айди
	pool   *Pool     // Ссылка на пул воркеров
	stopCh chan bool // Канал для остановки воркера
}

// Start запускает воркера в горутине
func (w *Worker) Start() {
	go func() {
		fmt.Printf("Worker %d added\n", w.id)
		for {
			select {
			case task := <-w.pool.taskQueue: // Получаем задачу из очереди
				fmt.Printf("Worker %d working on task: %s\n", w.id, task)
			case <-w.stopCh: // Получаем сигнал на остановку воркера
				fmt.Printf("Worker %d stopped\n", w.id)
				return
			}
		}
	}()
}
