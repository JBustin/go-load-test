package worker

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/go-load-test/config"
	"github.com/go-load-test/utils"
)

type worker struct {
	config config.Config
	tasks  []Tasker
	logger utils.Logger
}

func New(urls []string, config config.Config) worker {
	var tasks []Tasker

	for id, url := range urls {
		tasks = append(tasks, newTask(config, id, url, false))
	}

	return worker{
		config: config,
		tasks:  tasks,
		logger: utils.NewLog(config.LogLevel),
	}
}

func (w *worker) Process() error {
	groups := ranges(w.tasks, w.config.Concurrency)

	if w.config.IsSerie {
		for i, group := range groups {
			w.ProcessByGroup(group, i*w.config.Concurrency)
		}
	} else {
		var wg sync.WaitGroup

		for i, group := range groups {
			wg.Add(1)
			time.Sleep(time.Duration(i*w.config.WaitMs) * time.Millisecond)
			go func(group []Tasker, i int) {
				w.ProcessByGroup(group, i*w.config.Concurrency)
				wg.Done()
			}(group, i)
		}

		wg.Wait()
	}

	return nil
}

func (w *worker) ProcessByGroup(tasks []Tasker, startIndex int) {
	var wg sync.WaitGroup

	w.logger.Debug(fmt.Sprintf("Range [%v - %v] (len: %v)\n", startIndex, startIndex+w.config.Concurrency, len(tasks)))

	for i, task := range tasks {
		wg.Add(1)
		go func(i int, task Tasker) {
			defer wg.Done()

			w.logger.Info(task.RequestStr())

			err := task.Request()
			if err != nil {
				task.SetHasFailed(true)
				w.logger.Error(task.ErrorStr(err))
			}

			w.logger.Info(task.ResponseStr())
		}(i, task)
	}
	wg.Wait()
}

func (w *worker) Report() string {
	total := len(w.tasks)
	failures := 0
	connectDuration := 0
	reqDuration := 0
	duration := 0
	codes := make(map[int]int)
	codeKeys := []int{}
	max := 0
	min := int(math.Inf(1))

	for _, t := range w.tasks {
		statusCode := t.StatusCode()
		_, exists := codes[statusCode]
		if exists {
			codes[statusCode] += 1
		} else {
			codeKeys = append(codeKeys, statusCode)
			codes[statusCode] += 1
		}

		if t.HasFailed() {
			failures += 1
		} else {
			connectDuration += int(t.Duration("connect").Milliseconds())
			reqDuration += int(t.Duration("request").Milliseconds())
			duration += int(t.Duration("").Milliseconds())
			if min > int(t.Duration("").Milliseconds()) {
				min = int(t.Duration("").Milliseconds())
			}
			if max < int(t.Duration("").Milliseconds()) {
				max = int(t.Duration("").Milliseconds())
			}
		}
	}

	codesStr := ""
	sort.Ints(codeKeys)
	for _, code := range codeKeys {
		codesStr += fmt.Sprintf(`
		%v => 	 	%v`, code, codes[code])
	}

	return fmt.Sprintf(`
	Hits: 		 	%v
	Failed: 	 	%v
	StatusCodes: 	%v
	Timings (avg):
		Connect: 	%v
		Request: 	%v
		Duration: 	%v
	`,
		total,
		failures,
		codesStr,
		float32(connectDuration/total),
		float32(reqDuration/total),
		float32(duration/total))
}

func ranges(mySlice []Tasker, size int) [][]Tasker {
	if len(mySlice) <= size {
		return [][]Tasker{mySlice}
	}

	result := make([][]Tasker, len(mySlice)/size)

	j := -1

	for i := 0; i < len(mySlice); i++ {
		if i%size == 0 {
			j++
			result[j] = []Tasker{}
		}
		result[j] = append(result[j], mySlice[i])
	}

	return result
}
