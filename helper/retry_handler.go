package helper

import "github.com/lancer2672/DandelionServer_Go/constants"

func RetryHandler(realService func() error) error {
	var err error
	for i := 0; i < constants.CALL_RETRY_TIMES; i++ {
		err = realService()
		if err == nil {
			return nil
		}
	}
	return err
}
