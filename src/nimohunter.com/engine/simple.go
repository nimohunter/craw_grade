package engine

import (
	"log"
	"nimohunter.com/model"
)

type SimpleEngine struct {
	WorkCount int
}

func (e *SimpleEngine) Run(seeds ...model.Request) {

	inChannel := make(chan model.Request, 1e4)
	outChannel := make(chan model.ParseResult, 1e4)
	for i := 0; i < e.WorkCount; i++ {
		go createWorker(inChannel, outChannel)
	}

	var requests []model.Request
	for _, r := range seeds {
		inChannel <- r
	}


	for  {
		select {
		case result, ok := <- outChannel:
			if ok == false {
				break
			}

			requests = append(requests, result.Requests...)

			for _, item := range result.Items {
				log.Printf("Got item %v", item)
			}
		}

	}
}


