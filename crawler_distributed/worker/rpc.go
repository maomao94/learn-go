package worker

import "learn-go/crawler/engine"

type CrawlService struct{}

func (CrawlService) Process(
	req Request,
	result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializedResult(engineResult)
	return nil
}
