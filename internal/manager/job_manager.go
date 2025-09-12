package manager

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/CharlesWhiteSun/gomodx/logger"
)

type JobManager struct {
	jobs         []IJob
	isSequential bool // 是否依序執行
	retryDelay   time.Duration
	wg           sync.WaitGroup
	mu           sync.Mutex
	results      map[string]error
}

func NewJobManager(sequential bool, retryDelay time.Duration) *JobManager {
	return &JobManager{
		isSequential: sequential,
		retryDelay:   retryDelay,
		results:      make(map[string]error),
	}
}

func (jm *JobManager) Register(job IJob) {
	jm.jobs = append(jm.jobs, job)
}

func (jm *JobManager) RunBySequential() (bool, map[string]error) {
	header := "JobManager RunBySequential|"
	if !jm.isSequential {
		emsg := fmt.Errorf("%v you want to use [RunWithThreads()] but sent to the wrong method", header)
		logger.Error(emsg)
		return false, nil
	}
	for _, job := range jm.jobs {
		jm.runWithRetry(job)
	}
	return true, jm.results
}

func (jm *JobManager) RunWithThreads() (bool, map[string]error) {
	header := "JobManager RunWithThreads|"
	if jm.isSequential {
		emsg := fmt.Errorf("%v you want to use [RunBySequential()] but sent to the wrong method", header)
		logger.Error(emsg)
		return false, nil
	}
	for _, job := range jm.jobs {
		jm.wg.Add(1)
		go func(j IJob) {
			defer jm.wg.Done()
			jm.runWithRetry(j)
		}(job)
	}
	jm.wg.Wait()
	return true, jm.results
}

func (jm *JobManager) runWithRetry(job IJob) {
	header := "JobManager runWithRetry|"
	for {
		err := job.Run()
		jm.mu.Lock()
		jm.results[job.Name()] = err
		jm.mu.Unlock()

		if err != nil {
			emsg := fmt.Errorf("%v [%s] job failed, Retrying in %s| error: %v", header, job.Name(), jm.retryDelay, err)
			logger.Error(emsg)
			time.Sleep(jm.retryDelay)
			continue
		}

		log.Printf("%v [%s] job completed successfully", header, job.Name())
		break
	}
}
