package thread_pool_utils

import (
	"fmt"
	"forum/logger"
	"sync"
)

type task struct {
	function interface{}
	args     []interface{}
}

type ThreadPool struct {
	workerCount int
	taskQueue   chan task
	wg          sync.WaitGroup
	done        chan struct{}
	isStarted   bool
	mutex       sync.Mutex
}

var (
	defaultPool     *ThreadPool
	defaultPoolOnce sync.Once
)

func GetDefaultPool() *ThreadPool {
	defaultPoolOnce.Do(func() {
		defaultPool = New(10)
		defaultPool.Start()
	})
	return defaultPool
}

func New(workerCount int) *ThreadPool {
	if workerCount <= 0 {
		workerCount = 1
	}

	return &ThreadPool{
		workerCount: workerCount,
		taskQueue:   make(chan task, workerCount*2),
		done:        make(chan struct{}),
		isStarted:   false,
	}
}

func (tp *ThreadPool) Start() {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()

	if tp.isStarted {
		return
	}

	tp.isStarted = true
	tp.wg.Add(tp.workerCount)
	for i := 0; i < tp.workerCount; i++ {
		go tp.worker()
	}
}

func Execute(function interface{}, args ...interface{}) {
	GetDefaultPool().Execute(function, args...)
}

func (tp *ThreadPool) Execute(function interface{}, args ...interface{}) {
	if !tp.isStarted {
		tp.Start()
	}

	t := task{
		function: function,
		args:     args,
	}

	// Send the task to the queue
	select {
	case <-tp.done:
		return
	case tp.taskQueue <- t:
	}
}

func (tp *ThreadPool) worker() {
	defer tp.wg.Done()

	for {
		select {
		case <-tp.done:
			return
		case t, ok := <-tp.taskQueue:
			if !ok {
				return
			}
			tp.executeTask(t)
		}
	}
}

func (tp *ThreadPool) executeTask(t task) {
	defer func() {
		if r := recover(); r != nil {
			logger.GetLogInstance().Error("Error while executing thread_pool_utils:\n\t" + fmt.Sprint(r))
		}
	}()

	// Use type assertions to call the function with its arguments
	switch f := t.function.(type) {
	case func():
		if len(t.args) == 0 {
			f()
		}
	case func(interface{}):
		if len(t.args) == 1 {
			f(t.args[0])
		}
	case func(interface{}, interface{}):
		if len(t.args) == 2 {
			f(t.args[0], t.args[1])
		}
	case func(interface{}, interface{}, interface{}):
		if len(t.args) == 3 {
			f(t.args[0], t.args[1], t.args[2])
		}
	case func(interface{}, interface{}, interface{}, interface{}):
		if len(t.args) == 4 {
			f(t.args[0], t.args[1], t.args[2], t.args[3])
		}
	case func(interface{}, interface{}, interface{}, interface{}, interface{}):
		if len(t.args) == 5 {
			f(t.args[0], t.args[1], t.args[2], t.args[3], t.args[4])
		}
	case func(...interface{}):
		f(t.args...)
	}
}

func (tp *ThreadPool) Shutdown() {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()

	if !tp.isStarted {
		return
	}

	close(tp.done)
	tp.wg.Wait()
	close(tp.taskQueue)
	tp.isStarted = false
}

func ShutdownDefaultPool() {
	if defaultPool != nil {
		defaultPool.Shutdown()
	}
}

func SetDefaultPoolSize(size int) {
	defaultPoolOnce.Do(func() {
		defaultPool = New(size)
		defaultPool.Start()
	})
}
