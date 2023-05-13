package utils

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func StartCron(pwd string) {
	cronSchedule, err := GetEnv("CRON_SCHEDULE")
	if err != nil {
		panic(err)
	}

	if _, err := cron.ParseStandard(cronSchedule); err != nil {
		panic("Invalid cron job string:" + cronSchedule)
	}

	fmt.Println("Starting cron job with schedule:", cronSchedule)

	c := cron.New()
	c.AddFunc(cronSchedule, func() {
		RunAllScript(pwd)
	})
	c.Start()
}
