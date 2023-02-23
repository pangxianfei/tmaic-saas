package bootstrap

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

func startSchedule() {
	c := cron.New()

	addCronFunc(c, "@every 30m", func() {

	})

	// Generate sitemap
	addCronFunc(c, "@every 2h", func() {

	})

	c.Start()
}

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		logrus.Error(err)
	}
}
