package service

import (
	"flag"

	"github.com/gogf/gf/os/glog"
	"github.com/kardianos/service"
	"github.com/sanrentai/service/log"
)

var logger service.Logger

type Program struct {
	Name        string // Required name of the service. No spaces suggested.
	DisplayName string // Display name, spaces allowed.
	Description string // Long description of service.
	exit        chan struct{}
	Runfunc     func()
}

func (p *Program) Start(s service.Service) error {
	if service.Interactive() {
	} else {
		log.SetStdoutPrint(false)
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *Program) run() error {
	p.Runfunc()

	return nil
}
func (p *Program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}

func Run(prg *Program) {
	svcFlag := flag.String("s", "", "Control the system service.")
	flag.Parse()
	svcConfig := &service.Config{
		Name:        prg.Name,
		DisplayName: prg.DisplayName,
		Description: prg.Description,
	}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		glog.Fatal(err)
	}
	errs := make(chan error, 5)
	l, err := s.SystemLogger(errs)
	if err != nil {
		glog.Fatal(err)
	}
	logger = log.New(l)
	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
