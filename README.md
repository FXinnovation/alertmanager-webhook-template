# alertmanager-webhook-template
[![Build Status](https://travis-ci.org/FXinnovation/alertmanager-webhook-template.svg?branch=master)](https://travis-ci.org/FXinnovation/alertmanager-webhook-template)

The alertmanager-webhook-template is a basic Prometheus AlertManager webhook receiver template. 
The goal of this project is to provide a minimal start point from which any webhook receiver may start.
It contains a very minimal set of features common to any webhook receiver

## Getting Started

In summary, copy this project's code to get your webhook receiver started.

### Prerequisites

To run this project, you will need a [working Go environment](https://golang.org/doc/install).

### Installing

```bash
go get -u github.com/FXinnovation/alertmanager-webhook-template
```

## Running the tests

```bash
go test
```

## Deployment

The template listen on port 9876 by default, be sure to change that port in order to suit your new receiver needs.

## Contributing

Refer to [CONTRIBUTING.md](https://github.com/FXinnovation/alertmanager-webhook-template/blob/master/CONTRIBUTING.md).

## License

Apache License 2.0, see [LICENSE](https://github.com/FXinnovation/alertmanager-webhook-template/blob/master/LICENSE).
