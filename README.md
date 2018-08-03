```
 ____                       ____ _     ___ 
|  _ \ _ __ ___  _ __ ___  / ___| |   |_ _|
| |_) | '__/ _ \| '_   _ \| |   | |    | | 
|  __/| | | (_) | | | | | | |___| |___ | | 
|_|   |_|  \___/|_| |_| |_|\____|_____|___|
```

CLI shell for [Prometheus](https://prometheus.io) for testing Prometheus PromQL expressions.
This allows users to manually add metrics and evaluate PromQL expressions against the metrics for testing purposes.

Demo
====

![Demo](demo/demo.svg)
[![asciicast](https://asciinema.org/a/WSsYo9Yo5UP3RubRyLqJyjV0Y.png)](https://asciinema.org/a/WSsYo9Yo5UP3RubRyLqJyjV0Y)


Install
=======

    go get github.com/kevinjqiu/promcli

Build
=====

    git clone https://github.com/kevinjqiu/promcli.git
    deps ensure
    make build
