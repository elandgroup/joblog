# joblog

job log for eland job monitor ,job log api:https://wiki.elandsystems.cn/display/DBA/Batch-job+Monitor

## Getting Started

```golang
jobLog := joblog.New(url,"test", map[string]interface{}{"log": "this is test"})
err := jobLog.Info("good")
test.Ok(t, err)
err = jobLog.Warning(struct{ Name string }{"xiaoxinmiao"})
test.Ok(t, err)
err = jobLog.Error(errors.New("this is bug."))
test.Ok(t, err)
```

Note: the **new** method's param parameter can only pass `objects`, `pointer objects` or `maps`.
