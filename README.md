# pool

simple goroutine pool

Usage:

Step 1: define worker function

  func worker(task interface{}) {

    num, ok := task.(int)

    if ok {

      fmt.Printf("Do worker func %d\n", num)
    } else {

      fmt.Println("invalid task!!!")
    }

    time.Sleep(10*time.Second)

  }


Step 2:

  // Create Pool Instance

  var poolObj = pool.NewPool(5, 5, worker)


  
// Adjust Pool Size(Add/Remove)

  poolObj.AdjustPoolSize(10)


  
// Add Task

  poolObj.AddTask(1)

