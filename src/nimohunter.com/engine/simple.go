package engine

import (
	"nimohunter.com/model"
)

type SimpleEngine struct {
	WorkCount        int
	ItemCollectCount int
	ItemChain        chan model.Item
}

func (e *SimpleEngine) Run(seeds ...model.Request) {

	inChannel := make(chan model.Request, 1e4)
	outChannel := make(chan model.ParseResult, 1e4)
	startSignal := make(chan int)
	e.ItemChain = make(chan model.Item, 1e4)

	for _, r := range seeds {
		inChannel <- r
	}

	for i := 0; i < e.WorkCount; i++ {
		go createWorker(startSignal, inChannel, outChannel)
	}

	for i := 0; i < e.ItemCollectCount; i++ {
		go createItemCollectWorker(startSignal, e.ItemChain)
	}

	close(startSignal)

	for {
		select {
		case result, ok := <-outChannel:
			if ok == false {
				break
			}

			go func() {
				for _, value := range result.Items {
					e.ItemChain <- value
				}
			}()

			for _, request := range result.Requests {
				inChannel <- request
			}
		}

	}
}
