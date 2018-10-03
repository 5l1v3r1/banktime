# moov-io/banktime

[![GoDoc](https://godoc.org/github.com/moov-io/banktime?status.svg)](https://godoc.org/github.com/moov-io/banktime)
[![Build Status](https://travis-ci.com/moov-io/banktime.svg?branch=master)](https://travis-ci.com/moov-io/banktime)
[![Coverage Status](https://codecov.io/gh/moov-io/banktime/branch/master/graphs/badge.svg?branch=master)](https://codecov.io/gh/moov-io/banktime/)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/banktime)](https://goreportcard.com/report/github.com/moov-io/banktime)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/banktime/master/LICENSE)

A golang library for working with US Federal Reserve Banks days of payment processing

- All times will be converted to Eastern Time (America/New_York)
- If a holiday falls on a sunday the following monday is an observed holiday
- Add banking days skips holidays, weekends, and observed days

Project is based on the excellent calendar library by [@rickar](https://github.com/rickar).

```text
$ go get "github.com/rickar/cal"
```

Some example use cases:

```go
package main

import (
    "time"
    "github.com/moov-io/banktime"
)

func main() {
    t := time.Date(2018, time.January, 11, 1, 0, 0, 0, time.Local)

    // Post a payment today for same-day ACH?
    sameDay := NewBankTime(T).IsBankDay()
    println(sameDay) // true

    // Need to post an normal ACH payment in two banking days.
    // This date has a weekend and monday holiday
    postingDate := NewBankTime(T).AddBankingDay(2)
    println(postingDate) // 2018-01-16 01:00:00 -0500 EST
}
```

## Getting Help

 channel | info
 ------- | -------
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce an problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](http://moov-io.slack.com/) | Join our slack channel to have an interactive discussion about the development of the project. [Request an invite to the slack channel](https://join.slack.com/t/moov-io/shared_invite/enQtNDE5NzIwNTYxODEwLTRkYTcyZDI5ZTlkZWRjMzlhMWVhMGZlOTZiOTk4MmM3MmRhZDY4OTJiMDVjOTE2MGEyNWYzYzY1MGMyMThiZjg)

## Contributing

Yes please! Please start by reviewing our [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) then start creating or responding issues and fixing bugs.

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
