> [!NOTE]  
> To add this hook you need to use [Logrus](https://github.com/sirupsen/logrus)

# 👊🏾Ghettoize Hook👊🏾

Do you find boring to go through your logs?, don't worry I may have a solution...

## How To Use It?

Ghettoize needs [ollama](https://ollama.com/). By default, is checking ollama locally in http://localhost:11434 with the model [gemma2:2b](https://ollama.com/library/gemma2:2b)

Get ghettoize by:

> go get github.com/manuelarte/ghettoize_hook

And now add the hook to logrus like:

```go
ghettoizeHook, _ := ghuettoize.NewGhettoize()
logrus.AddHook(ghettoizeHook)
```

And then... just produce some logs.

## Example

You can find an example on how to use ghettoize_hook, and what it looks like in the [examples](./examples) folder.

## TODO

+ Being able to configure it better
+ Add roleplays
