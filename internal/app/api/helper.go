package api

import (
	"github.com/sirupsen/logrus"
	"github.com/srm-developer/standardWebServer/storage"
	"net/http"
)

func (api *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(logLevel)
	return nil
}

func (api *API) configureRouterField() {
	api.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is rest api!"))
	})
}

func (api *API) configureStorageField() error {
	storageInstance := storage.New(api.config.Storage)

	if err := storageInstance.Open(); err != nil {
		return err
	}
	api.storage = storageInstance
	return nil
}
