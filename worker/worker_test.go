package worker

import (
	"errors"
	"testing"

	"github.com/go-load-test/config"
	"github.com/go-load-test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_WorkerProcess(t *testing.T) {
	task0 := mocks.NewTask(200)
	task1 := mocks.NewTask(404)
	tasks := []Tasker{&task0, &task1}

	w := worker{
		config: config.DefaultConfig,
		tasks:  tasks,
	}

	task0.On("RequestStr").Return("something")
	task0.On("ResponseStr").Return("something")
	task0.On("Request").Return(nil)

	task1.On("RequestStr").Return("something")
	task1.On("ResponseStr").Return("something")
	task1.On("Request").Return(nil)

	err := w.Process()

	task0.AssertExpectations(t)
	task1.AssertExpectations(t)

	assert.Equal(
		t,
		nil,
		err,
		"Worker process should not return an error",
	)

	assert.Equal(
		t,
		`
	Hits: 		 	2
	Failed: 	 	0
	StatusCodes: 	
		200 => 	 	1
		404 => 	 	1
	Timings (avg):
		Connect: 	0
		Request: 	0
		Duration: 	0
	`,
		w.Report(),
		"Worker report should be valid",
	)
}
func Test_WorkerProcessWithError(t *testing.T) {
	task0 := mocks.NewTask(200)
	task1 := mocks.NewTask(404)
	tasks := []Tasker{&task0, &task1}
	customConfig := config.DefaultConfig
	customConfig.IsSerie = true

	w := worker{
		config: customConfig,
		tasks:  tasks,
	}

	task0.On("RequestStr").Return("something")
	task0.On("ErrorStr", errors.New("Boum")).Return("something")
	task0.On("ResponseStr").Return("something")
	task0.On("Request").Return(errors.New("Boum"))

	task1.On("RequestStr").Return("something")
	task1.On("ResponseStr").Return("something")
	task1.On("Request").Return(nil)

	err := w.Process()

	task0.AssertExpectations(t)
	task1.AssertExpectations(t)

	assert.Equal(
		t,
		nil,
		err,
		"Worker process should not return an error",
	)

	assert.Equal(
		t,
		`
	Hits: 		 	2
	Failed: 	 	1
	StatusCodes: 	
		200 => 	 	1
		404 => 	 	1
	Timings (avg):
		Connect: 	0
		Request: 	0
		Duration: 	0
	`,
		w.Report(),
		"Worker report should be valid",
	)
}

func Test_WorkerRanges(t *testing.T) {
	mySlice := []Tasker{
		&mocks.Task{},
		&mocks.Task{},
		&mocks.Task{},
		&mocks.Task{},
		&mocks.Task{},
		&mocks.Task{},
	}

	assert.Equal(
		t,
		[][]Tasker{{
			&mocks.Task{},
			&mocks.Task{},
			&mocks.Task{},
		}, {
			&mocks.Task{},
			&mocks.Task{},
			&mocks.Task{},
		}},
		ranges(mySlice, 3),
		"Ranges with 3 elements from 6",
	)

	assert.Equal(
		t,
		[][]Tasker{{
			&mocks.Task{},
			&mocks.Task{},
		}, {
			&mocks.Task{},
			&mocks.Task{},
		}, {
			&mocks.Task{},
			&mocks.Task{},
		}},
		ranges(mySlice, 2),
		"Ranges with 2 elements from 6",
	)

	mySlice = []Tasker{&mocks.Task{}}

	assert.Equal(
		t,
		[][]Tasker{{&mocks.Task{}}},
		ranges(mySlice, 2),
		"Ranges with less elements than targetted size",
	)

	mySlice = []Tasker{&mocks.Task{}, &mocks.Task{}}

	assert.Equal(
		t,
		[][]Tasker{{&mocks.Task{}, &mocks.Task{}}},
		ranges(mySlice, 2),
		"Ranges with as many elements as targetted size",
	)
}
