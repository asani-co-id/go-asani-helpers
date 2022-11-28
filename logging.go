package asanihelper

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	LogConfig logrus.Logger
	Active    bool
}

func (m LoggerConfig) LogInfo(message interface{}) {
	m.LogConfig.SetOutput(os.Stdout)
	m.LogConfig.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
		FullTimestamp: true,
	})

	m.LogConfig.Level = logrus.InfoLevel

	if m.Active {
		// m.LogConfig.Infoln(message)
		m.LogConfig.Logln(m.LogConfig.Level, message)
	}

}

func (m LoggerConfig) LogWarning(message interface{}) {
	m.LogConfig.SetOutput(os.Stdout)
	m.LogConfig.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
		FullTimestamp: true,
	})

	m.LogConfig.Level = logrus.WarnLevel

	if m.Active {
		m.LogConfig.Logln(m.LogConfig.Level, message)
	}

}

func (m LoggerConfig) LogError(err error) {
	m.LogConfig.SetOutput(os.Stdout)
	m.LogConfig.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
		FullTimestamp: true,
	})

	m.LogConfig.Level = logrus.ErrorLevel

	if m.Active {
		m.LogConfig.Logln(m.LogConfig.Level, err)
	}

}

func (m LoggerConfig) LogDebug(message interface{}) {
	m.LogConfig.SetOutput(os.Stdout)
	m.LogConfig.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
		FullTimestamp: true,
	})

	m.LogConfig.Level = logrus.DebugLevel

	if m.Active {
		m.LogConfig.Logln(m.LogConfig.Level, message)
	}

}

func (m LoggerConfig) LogFatal(err error) {
	m.LogConfig.SetOutput(os.Stdout)
	m.LogConfig.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
		FullTimestamp: true,
	})

	m.LogConfig.Level = logrus.FatalLevel

	if m.Active {
		m.LogConfig.Logln(m.LogConfig.Level, err)
	}

}
