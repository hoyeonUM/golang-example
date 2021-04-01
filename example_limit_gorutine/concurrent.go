package example_limit_gorutine

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func SelectTable(offset int) {
	time.Sleep(300 * time.Millisecond)
	log.Println(fmt.Sprintf("select * from table limit %d,10", offset))
}

func Run(concurrentCount int, totalCount int) {
	concurrentGoroutines := make(chan struct{}, concurrentCount)
	var wg sync.WaitGroup
	for offset := 0; offset < totalCount; offset += 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			concurrentGoroutines <- struct{}{}
			SelectTable(i)
			<-concurrentGoroutines

		}(offset)

	}
	wg.Wait()
}
