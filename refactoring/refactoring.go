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

// init  define flags by run program
func init() {
	flag.IntVar(&dur, "d", 10, "set working time -d 10")
	flag.IntVar(&freq, "f", 3, "set frequency time -f 3")
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

// Task structure for filling
type Task struct {
	id         int
	createdAt  string
	updatedAt  string
	taskResult []byte
}

// String method for printing in fmt package
func (t Task) String() string {
	return fmt.Sprint(t.id)
}

// TaskCreator create tasks in a loop for a specified time
func TaskCreator(wg *sync.WaitGroup, ch chan Task, dur int) {
	timeAfter := time.After(time.Second * time.Duration(dur))
	defer close(ch)
	for {
		select {
		case <-timeAfter:
			return
		default:
			wg.Add(1)
			fTime := time.Now().Format(time.RFC3339)
			if time.Now().Nanosecond()%2 > 0 {
				fTime = "Some error occurred"
			}
			ch <- Task{createdAt: fTime, id: int(time.Now().Unix())}
		}
	}
}

// TaskFiller fill taskResult and updateAt field in Task instance
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

// TaskSorter sorting tasks on done and error result
func TaskSorter(wg *sync.WaitGroup, mu *sync.RWMutex, task Task, doneTasks *[]Task, undoneTasks *[]string) {
	defer wg.Done()
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

func Printer(mu *sync.RWMutex, doneTasks *[]Task, undoneTasks *[]string) {
	fmt.Println("Error:")
	mu.Lock()
	for _, res := range *undoneTasks {
		fmt.Println(res)
	}
	mu.Unlock()
	fmt.Println("Done:")
	mu.Lock()
	for _, res := range *doneTasks {
		fmt.Println(res)
	}
	mu.Unlock()
}

func main() {
	flag.Parse()
	var mu sync.RWMutex
	var wg sync.WaitGroup
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	ticker := time.NewTicker(time.Second * time.Duration(freq))
	ctrl := make(chan struct{})

	taskChan := make(chan Task, 10)
	go TaskCreator(&wg, taskChan, dur)

	doneTasks := make([]Task, 0)
	undoneTasks := make([]string, 0)

	go func() {
	endWork:
		for {
			select {
			case c, ok := <-taskChan:
				if !ok {
					break endWork
				}
				go TaskSorter(&wg, &mu, TaskFiller(c), &doneTasks, &undoneTasks)
			case <-ticker.C:
				go Printer(&mu, &doneTasks, &undoneTasks)
			case <-ctx.Done():
				return
			}
		}
		go func() {
			wg.Wait()
			ctrl <- struct{}{}
		}()
	}()

	<-ctrl
	Printer(&mu, &doneTasks, &undoneTasks)
}
