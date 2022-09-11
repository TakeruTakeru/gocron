# gocron

utilities for making cron schedule text format.

## Install

```sh
go get github.com/TakeruTakeru/gocron
```

```go
import (
    "github.com/TakeruTakeru/gocron"
)
```

## Example

```go
// from January to April, first day of month, while PM1 to PM2, every 10 minutes...
schedule := gocron.Schedule().Days(1).MonthsRange(time.January, time.April).Hours(13, 14).MinutesInterval(10)

// */10 13,14 1 1-4 *
fmt.Println(schedule)
```

## License
MIT