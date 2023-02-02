package jlog

import "github.com/sirupsen/logrus"

type LogrusLogger struct {
	Log          *logrus.Logger
	ReportCaller bool
}
