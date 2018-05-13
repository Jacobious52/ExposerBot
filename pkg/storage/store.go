package storage

import (
	"encoding/json"
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// Write writes the model to disk
func (s *DataStore) Write(w io.Writer) error {
	err := json.NewEncoder(w).Encode(&s.model)
	if err != nil {
		return err
	}
	return nil
}

// Load reads the model to disk
func (s *DataStore) Load() error {
	file, err := os.Open(s.path)
	if err != nil {
		return err
	}
	defer file.Close()

	s.Lock()
	defer s.Unlock()

	err = json.NewDecoder(file).Decode(&s.model)
	if err != nil {
		return err
	}
	return nil
}

// StartSync starts a routine to periodically save the db if there are changes
func (s *DataStore) StartSync(stop <-chan struct{}) {
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for {
			select {
			case <-ticker.C:
				s.Lock()
				if s.dirty {
					file, err := os.Create(s.path)
					if err != nil {
						log.Errorf("data not saved. %v", err)
						s.Unlock()
						break
					}
					err = s.Write(file)

					if err != nil {
						log.Errorf("data not saved. %v", err)
					} else {
						log.Infoln("saved data")
						s.dirty = false
					}
					file.Close()
				}
				s.Unlock()
			case <-stop:
				log.Infoln("stopping data sync")
				ticker.Stop()
				break
			}
		}
	}()
}
