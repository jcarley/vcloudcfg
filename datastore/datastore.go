package datastore

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jcarley/vcloudcfg/model"
)

type Datastore struct {
	Boxes    *BoxList
	Filename string
}

type BoxList struct {
	Boxes []model.Box `json:boxes`
}

func NewDatastore() *Datastore {
	return &Datastore{Filename: "Boxfile.json"}
}

func (ds *Datastore) Load() error {
	file, err := os.OpenFile(ds.Filename, os.O_RDONLY, 0666)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln("Boxfile.json does not exist", err)
		}
		return err
	}
	defer file.Close()

	var boxes *BoxList
	err = json.NewDecoder(file).Decode(&boxes)
	if err != nil {
		return err
	}
	ds.Boxes = boxes
	return nil
}

func (ds *Datastore) Save() error {
	file, err := os.OpenFile(ds.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(ds.Boxes)
	return err
}

func (ds *Datastore) GetBox(boxname string) (*model.Box, error) {
	return nil, nil
}
