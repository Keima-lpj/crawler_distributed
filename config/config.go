package config

const (
	URL = "https://www.zhenai.com/zhenghun"

	//rpc port
	ITEM_SAVE_HOST = "127.0.0.1:1234"
	WORKER_PORT0   = "9000"

	//elasticSearch index
	ES_INDEX = "immoc5"

	//rpc service
	ITEM_SAVE_SERVICE = "ItemSaverService.Save"
	WORKER_SERVICE    = "CrawlService.Process"

	//Parser names
	ParseCityList = "ParserCityList"
	ParseCity     = "ParserCity"
	ParseProfile  = "ParserProfile"
)
