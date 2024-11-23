# TIM:

_Small time input library for GoLang_

TIM aims to provide an handful parser that parses legit [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) time such that `1991-08-06T16:56:20+02:00` but also accepts some other forms such that:

- `1991-08-06T16:56:20`, means _August 6th, 1991 at 16:56:20_, according to the current time zone
- `now`, local time on the current computer (with local timezone)

It also allows parse time specified with simple shift operations, such that: 

- `1991-08-06T16:56:20+02:00-12h`, means _12 hours_ before _August 6th, 1991 at 16:56:20 GMT+2_.
- `now-1w1d`, one week and a day ago
