package joblog_test

import (
	"errors"
	"os"
	"testing"

	"github.com/elandgroup/joblog"

	"github.com/pangpanglabs/goutils/test"
)

var url = os.Getenv("JOB_LOG")

func TestLog(t *testing.T) {
	t.Run("test normal logic", func(t *testing.T) {
		jobLog := joblog.New(url,
			"test", map[string]interface{}{"log": "this is test"})
		err := jobLog.Info("good")
		test.Ok(t, err)

		err = jobLog.Warning(struct{ Name string }{"xiaoxinmiao"})
		test.Ok(t, err)

		err = jobLog.Error(errors.New("this is bug."))
		test.Ok(t, err)

		err = jobLog.Finish()
		test.Ok(t, err)
	})

	t.Run("test options param", func(t *testing.T) {
		//test log.Disable:this content will not be logged
		jobLog := joblog.New(url,
			"test", map[string]interface{}{"log": "this is test 2"}, func(log *joblog.JobLog) {
				log.JobName = "JobName"
				log.ActionName = "ActionName"
				log.Disable = true
			})
		test.Ok(t, jobLog.Err)
		err := jobLog.Info("good 2")
		test.Ok(t, err)

		err = jobLog.Warning(struct{ Name string }{"xiaoxinmiao 2"})
		test.Ok(t, err)

		err = jobLog.Error(errors.New("this is bug. 2"))
		test.Ok(t, err)
	})
	t.Run("test firstMessage param type", func(t *testing.T) {
		type Dto struct{ Message string }
		test.Ok(t, joblog.New(url, "test", Dto{Message: "how are you 1"}).Err)
		test.Ok(t, joblog.New(url, "test", &Dto{Message: "how are you 2"}).Err)
		test.Ok(t, joblog.New(url, "test", map[string]interface{}{"message": "how are you 3"}).Err)
	})

	t.Run("test http response", func(t *testing.T) {
		type Dto struct{ Message string }
		test.Equals(t, "request err,status is 404", joblog.New("http://localhost:8080/v1/xxx", "test", Dto{Message: "how are you 1"}).Err.Error())
	})

}
