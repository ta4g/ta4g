# ta4g

[![CircleCI](https://circleci.com/gh/ta4g/ta4g.svg?style=svg)](https://circleci.com/gh/ta4g/ta4g.svg?style=svg)

***Technical Analysis For Go***

Ta4g is an open source Java library for [technical analysis](http://en.wikipedia.org/wiki/Technical_analysis). 

It provides the basic components for creation, evaluation and execution of trading strategies.

This is a port of the popular technical analysis library [ta4j](https://github.com/ta4j/ta4j) from Java -> Golang.

---

**Why convert this library to Golang?**

1. Just for fun, I enjoy learning about new subjects and this is an area I'm not familiar with (yet).
1. I predominately write Golang CLIs, GRPC services, and other applications so using Golang over Java is my preference.
1. To contribute the open source community by converting one of the best libraries for technical analysis to enable more developers to leverage their good work.

---

## Features

1. [x] 100% Pure GO, targeting go 1.16+
1. [ ] More than 130 technical indicators (Aroon, ATR, moving averages, parabolic SAR, RSI, etc.)
1. [ ] A powerful engine for building custom trading strategies
1. [ ] Utilities to run and compare strategies
1. [ ] Minimal 3rd party dependencies
1. [ ] Simple integration
1. [ ] GRPC server/client libraries for easy cross-platform integration.
1. [ ] One more thing: it's Apache License 2.0 licensed

## Roadmap

1. [x] Initial Repo Setup
1. [x] CI/CD configuration
1. [ ] GO Releaser configuration
1. [ ] Document project layout and usage
1. [ ] Implement core structs: Bar, Series, Indicator, Order, Rule, Trade, TradeRecord
1. [ ] Implement backtest framework
1. [ ] Implement feature: charts [go-chart](with https://github.com/wcharczuk/go-chart)
1. [ ] Implement feature: aggregator
1. [ ] Implement feature: analysis
1. [ ] Implement feature: cost
1. [ ] Implement feature: indicators
1. [ ] Implement feature: num
1. [ ] Implement feature: tradereport
1. [ ] Implement feature: tradingrules
1. [ ] Documentation and cleanup 
