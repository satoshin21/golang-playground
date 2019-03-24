package lessons

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1)

/*
semaphore.Accuire()
semaphore.Release()
*/
func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("##############")
}

func longProcess(ctx context.Context) {

	// Actuireできない場合はそのまま終了させる場合
	//isAcquire := s.TryAcquire(1)
	//if !isAcquire {
	//	fmt.Println("Could not ge lock")
	//	return
	//}
	//defer s.Release(1)

	// semaphoreを待つ場合
	if error := s.Acquire(ctx, 1); error != nil {
		fmt.Println(error)
		return
	}
	defer s.Release(1)

	fmt.Println("wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}