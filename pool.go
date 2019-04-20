package pool

type Pool struct {
	PoolSize int
	TaskSize int
	TaskChan chan interface{}
	ctrlChan chan struct{}
	Task func(task interface{})
}

func (poolObj *Pool) init() {
	for i:=0; i<poolObj.PoolSize; i++ {
		go poolObj.worker()
	}
}

func (poolObj *Pool) AddTask(task interface{}) {
	poolObj.TaskChan <- task
}

func (poolObj *Pool) AdjustPoolSize(poolSize int) {
	if poolSize > poolObj.PoolSize {
		for i:= poolObj.PoolSize; i<poolSize; i++ {
			go poolObj.worker()
		}
		poolObj.PoolSize = poolSize
	} else if poolSize < poolObj.PoolSize {
		for i:=poolSize; i<poolObj.PoolSize; i++ {
			poolObj.ctrlChan <- struct{}{}
		}
		poolObj.PoolSize = poolSize
	}
}

func (poolObj *Pool) worker() {
	for {
		stop := false
		select {
			case task := <- poolObj.TaskChan:
				poolObj.Task(task)
			case <- poolObj.ctrlChan:
				stop = true 
		}

		if stop {
			break
		}
	}
}


/* Create Pool Instance
 */
func NewPool(poolSize int, taskSize int, worker func(interface{})) *Pool {
	taskPool := & Pool{
		PoolSize: poolSize,
		TaskSize: taskSize,
		TaskChan: make(chan interface{}, taskSize),
		ctrlChan: make(chan struct{}),
		Task: worker,
	}
	taskPool.init();

	return taskPool
}
