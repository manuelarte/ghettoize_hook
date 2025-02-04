# Ghettoize Hook

Do you find boring to go through your logs, don't worry

## How To Use It

Ghettoize needs [ollama](https://ollama.com/). By default, is checking ollama locally  http://localhost:11434 with the model [gemma2:2b](https://ollama.com/library/gemma2:2b)

Get ghettoize by:

> go get github.com/manuelarte/ghettoize_hook

And now add the hook to logrus like:

```go
ghettoizeHook, _ := ghuettoize.NewGhettoize()
logrus.AddHook(ghettoizeHook)
```

## Example

You can find an example in the [examples](./examples) folder.

## TODO

+ Being able to configure it better
+ Add roleplays
