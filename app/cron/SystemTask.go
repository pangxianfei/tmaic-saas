package cron

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

func StartSchedule() {
	c := cron.New()
	addCronFunc(c, "@every 30m", func() {})
	addCronFunc(c, "@every 2h", func() {})
	c.Start()
}

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		logrus.Error(err)
	}
}
