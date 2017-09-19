# banktime 
A golang library for working with US Federal Reserve Banks days of payment processing

- All times will be converted to Eastern Time (America/New_York)
- If a holiday falls on a sunday the following monday is an observed holiday 
- Add banking days skips holidays, weekends, and observed days

Project is based on the excellent calendar library by @rickar

```text
go get "github.com/rickar/cal"
```

Some example use cases 

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
