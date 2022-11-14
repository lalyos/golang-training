package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bombsimon/logrusr/v3"
	kitlog "github.com/go-kit/log"
	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	"github.com/sirupsen/logrus"
	"github.com/tonglil/gokitlogr"
)

func stdLog() {
	l := log.New(os.Stderr, "---> ", 0)
	//l := log.New(os.Stderr, "---> ", log.Lshortfile|log.Ldate|log.Ltime)
	l.Println("starting ...")
	l.Println("end")
}

func useStdr() logr.Logger {
	var stdl stdr.StdLogger = nil
	var o stdr.Options = stdr.Options{}

	//o = stdr.Options{LogCaller: stdr.Error}
	//stdl = log.New(os.Stderr, "---> ", log.Lshortfile|log.Ldate|log.Ltime)

	//l := stdr.New(stdl)
	l := stdr.NewWithOptions(stdl, o)
	return l.WithName("myapp").WithValues("env", "qa")
}

func useKitlog() logr.Logger {
	//kl := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	kl := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stderr))
	kl = kitlog.With(kl, "ts", kitlog.DefaultTimestampUTC, "caller", kitlog.Caller(5))

	gokitlogr.NameFieldKey = "logger"
	gokitlogr.NameSeparator = "/"
	var log logr.Logger = gokitlogr.New(&kl)

	log = log.WithName("my app")
	log = log.WithValues("format", "json")

	return log
}

func useLogrus() logr.Logger {
	logrusLog := logrus.New()
	logrusLog.SetLevel(logrus.ErrorLevel)
	if !quiet {
		logrusLog.SetLevel(logrus.AllLevels[logLevel+int(logrus.InfoLevel)])
	}
	log := logrusr.New(logrusLog)

	return log.WithName("MyName").WithValues("user", "you")
}

// START OMIT
var logLevel int
var quiet bool

func init() {
	flag.BoolVar(&quiet, "q", false, "quiet")
	flag.IntVar(&logLevel, "v", 0, "verbosity level")
	flag.Parse()
	// END OMIT
	if quiet {
		stdr.SetVerbosity(-1)
	} else {
		stdr.SetVerbosity(logLevel)
	}
}

func main() {
	// stdLog()
	//l := useStdr()
	//l := useKitlog()
	l := useLogrus()

	l.Info("started ...")
	l.V(1).Info("some details")
	l.V(2).Info("even more details")
	l.V(3).Info("who cares???")
	l.V(5).Info("too much information: faaaaart")
	l.Error(fmt.Errorf("boing"), "oooops...")
	l.Info("end")
}
