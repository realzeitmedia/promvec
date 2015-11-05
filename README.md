Prometheus metric vectors with fixed label value sets.

This is based on [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang),
but avoids that package's synchronization overhead when fetching metrics
by label.

See example_test.go for usage.

Installation:

    `go get github.com/realzeitmedia/promvec`
