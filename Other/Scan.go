package Other

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

func Scan() {
	target, err := ReadConf(TargetFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	ScanCh := make(chan string, len(target))
	Result := make(chan interface{}, 6*len(target))
	for _, val := range target {
		ScanCh <- val
		Wg.Add(1)
	}
	for n := 0; n < Thread; n++ {
		go func() {
			for url := range ScanCh {
				RequestHead(url, Result)
				Wg.Done()
			}

		}()
		go func() {
			for val := range Result {
				fmt.Println(val)
				Wg.Done()
			}
		}()
	}
	Wg.Wait()
}
