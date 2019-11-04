package persist

import (
	"crawler/engine"
	"log"
)

type ItemSaverService struct {
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	_, err := engine.Save(item, s.Index)
	if err == nil {
		*result = "ok"
		log.Printf("item %v save success", item)
	} else {
		log.Printf("item %v saved fail, error:%v", item, err)
	}
	return err
}
