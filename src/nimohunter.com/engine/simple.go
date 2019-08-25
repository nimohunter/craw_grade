package engine

import (
	"nimohunter.com/model"
)

type SimpleEngine struct {
	WorkCount int
}

func (e *SimpleEngine) Run(seeds ...model.Request) {

	inChannel := make(chan model.Request, 1e4)
	outChannel := make(chan model.ParseResult, 1e4)
	startSignal := make(chan int)

	for _, r := range seeds {
		inChannel <- r
	}

	for i := 0; i < e.WorkCount; i++ {
		go createWorker(startSignal, inChannel, outChannel)
	}

	close(startSignal)

	for {
		select {
		case result, ok := <-outChannel:
			if ok == false {
				break
			}

			for _, request := range result.Requests {
				inChannel <- request
			}
		}

	}
}
