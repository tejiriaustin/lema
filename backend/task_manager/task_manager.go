package task_manager

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/tejiriaustin/lema/database"
	"github.com/tejiriaustin/lema/env"
)

const (
	defaultMaxWorkers = 3
)

type (
	Run struct {
		maxWorkers    uint
		Config        *env.Environment
		Jobs          []Task
		Executor      map[string]Handler
		DBClient      *database.Client
		runningTasks  map[string]bool
		runningTasksM sync.Mutex
	}

	Task struct {
		ID           int
		TaskName     string
		Interval     time.Duration
		LastExecuted time.Time
	}

	Options func(*Run)

	Handler func(context.Context, *env.Environment)
)

// NewRunner initializes a new Run instance with provided options.
func NewRunner(opts ...Options) *Run {
	r := &Run{
		Jobs:         make([]Task, 0),
		Executor:     make(map[string]Handler),
		maxWorkers:   defaultMaxWorkers,
		runningTasks: make(map[string]bool),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// WithMaxWorkers sets the maximum number of workers.
func WithMaxWorkers(maxWorkers uint) Options {
	return func(r *Run) {
		r.maxWorkers = maxWorkers
	}
}

// WithConfig sets the configuration.
func WithConfig(cfg *env.Environment) Options {
	return func(r *Run) {
		r.Config = cfg
	}
}

// RunTasks begins the execution of registered jobs.
func (r *Run) RunTasks() {
	log.Print("Initializing repeat jobs...")

	workerChannels := make(chan string, r.maxWorkers)
	defer close(workerChannels)

	for i := uint(0); i < r.maxWorkers; i++ {
		go r.worker(context.Background(), workerChannels)
	}

	for {
		for index, job := range r.Jobs {
			now := time.Now()
			if now.After(job.LastExecuted.Add(job.Interval)) {
				r.runningTasksM.Lock()
				if !r.runningTasks[job.TaskName] {
					r.Jobs[index].LastExecuted = now
					r.runningTasks[job.TaskName] = true
					workerChannels <- job.TaskName
				}
				r.runningTasksM.Unlock()
			}
		}
		time.Sleep(5 * time.Second)
	}
}

// RegisterJob registers a new job with the runner.
func (r *Run) RegisterJob(taskName string, interval time.Duration, task Handler) *Run {
	log.Printf("Registering repeat job: %s", taskName)

	if taskName == "" {
		log.Println("Cannot register job with empty name")
		return r
	}
	if _, exists := r.Executor[taskName]; exists {
		log.Printf("Task already registered: %s", taskName)
		return r
	}
	if interval <= 0 {
		log.Printf("Cannot register job with invalid interval: %s", taskName)
		return r
	}
	if task == nil {
		log.Printf("Cannot register job with nil task: %s", taskName)
		return r
	}

	taskID := len(r.Jobs) + 1
	t := Task{
		ID:           taskID,
		TaskName:     taskName,
		Interval:     interval,
		LastExecuted: time.Time{},
	}
	r.Executor[taskName] = task
	r.Jobs = append(r.Jobs, t)
	return r
}

// worker processes tasks from the task channel.
func (r *Run) worker(ctx context.Context, taskChan <-chan string) {
	for taskName := range taskChan {
		if err := r.dispatcher(ctx, taskName); err != nil {
			log.Printf("Error occurred while dispatching task %s: %v", taskName, err)
		}
	}
}

// dispatcher handles the execution of a task.
func (r *Run) dispatcher(ctx context.Context, taskName string) error {
	log.Printf("Dispatching task: %s", taskName)
	defer func() {
		r.runningTasksM.Lock()
		r.runningTasks[taskName] = false
		r.runningTasksM.Unlock()
	}()

	handlerFunc, exists := r.Executor[taskName]
	if !exists || handlerFunc == nil {
		log.Printf("Handler function is nil for task: %s", taskName)
		return errors.New("handler function is nil")
	}

	handlerFunc(ctx, r.Config)
	return nil
}
