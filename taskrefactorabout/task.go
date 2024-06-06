package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	dur  int
	freq int
)

func init() {
	flag.IntVar(&dur, "duration", 10, "set working time --duration=10")
	flag.IntVar(&freq, "frequency", 3, "set frequency time --frequency=3")
}

// Приложение эмулирует получение и обработку неких тасков.
// Пытается и получать, и обрабатывать в многопоточном режиме.
// Приложение должно генерировать таски 10 сек.
// Каждые 3 секунды должно выводить в консоль результат
// всех обработанных к этому моменту тасков (отдельно успешные и отдельно с ошибками).

// ЗАДАНИЕ: сделать из плохого кода хороший и рабочий - as best as you can.
// Важно сохранить логику появления ошибочных тасков.
// Важно оставить асинхронные генерацию и обработку тасков.
// Сделать правильную мультипоточность обработки заданий.
// Обновленный код отправить через pull-request в github
// Как видите, никаких привязок к внешним сервисам нет - полный карт-бланш на модификацию кода.

type Task struct {
	id         int
	createdAt  string
	updatedAt  string
	taskResult []byte
}

func (t Task) String() string {
	return fmt.Sprint(t.id)
}

func TaskCreator(ch chan Task, dur int) {
	timeAfter := time.After(time.Second * time.Duration(dur))
	defer close(ch)
	for {
		select {
		case <-timeAfter:
			return
		default:
			fTime := time.Now().Format(time.RFC3339)
			if time.Now().Nanosecond()%2 > 0 {
				fTime = "Some error occurred"
			}
			ch <- Task{createdAt: fTime, id: int(time.Now().Unix())}
		}
	}
}

func TaskFiller(task Task) Task {
	t, err := time.Parse(time.RFC3339, task.createdAt)
	if err != nil || !t.After(time.Now().Add(-20*time.Second)) {
		task.taskResult = []byte("something went wrong")
	} else {
		task.taskResult = []byte("task has been successed")
	}
	task.updatedAt = time.Now().Format(time.RFC3339Nano)
	time.Sleep(time.Millisecond * 150)
	return task
}

func TaskSorter(mu *sync.RWMutex, task Task, doneTasks *[]Task, undoneTasks *[]string) {
	if string(task.taskResult[14:]) == "successed" {
		mu.Lock()
		*doneTasks = append(*doneTasks, task)
		mu.Unlock()
	} else {
		err := fmt.Sprint(task.id)
		mu.Lock()
		*undoneTasks = append(*undoneTasks, err)
		mu.Unlock()
	}
}

func main() {
	flag.Parse()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	ch := make(chan Task)
	go TaskCreator(ch, dur)

	var mu sync.RWMutex
	doneTasks := make([]Task, 0)
	undoneTasks := make([]string, 0)
	ticker := time.NewTicker(time.Second * time.Duration(freq))
	for {
		select {
		case c, ok := <-ch:
			if !ok {
				return
			}
			go TaskSorter(&mu, TaskFiller(c), &doneTasks, &undoneTasks)
		case <-ticker.C:
			go func() {
				fmt.Println("Error:")
				mu.RLock()
				for _, res := range undoneTasks {
					fmt.Println(res)
				}
				mu.RUnlock()
				fmt.Println("Done:")
				mu.RLock()
				for _, res := range doneTasks {
					fmt.Println(res)
				}
				mu.RUnlock()
			}()
		case <-ctx.Done():
			return
		}
	}
}
