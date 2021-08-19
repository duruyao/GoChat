package server

import (
	"log"
	"sync"
)

var serverOnce sync.Once
var serverInstance *Server

type Server struct {
	cfg    *Configurator
	logger *Logger
}

func GetServer(configurator *Configurator, logger *Logger) *Server {
	serverOnce.Do(func() {
		serverInstance = &Server{
			cfg:    configurator,
			logger: logger,
		}
	})
	return serverInstance
}

func (s *Server) Start() {
	log.Println("START")
	s.cfg.LoadConf()
	log.Println("Load configuration:\n" + s.cfg.String())
}

func (s *Server) Stop() {
	log.Println("STOP")
	s.logger.CloseFile()
}

func (s *Server) Restart() {
	s.logger.OpenFile()
	log.Println("RESTART")
}

func (s *Server) Pause() {
	log.Println("PAUSE")
}

func (s *Server) Resume() {
	log.Println("RESUME")
}
