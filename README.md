#Ciao Project

[![Go Report Card](https://goreportcard.com/badge/github.com/01org/ciao)](https://goreportcard.com/report/github.com/01org/ciao)
[![Build Status](https://travis-ci.org/01org/ciao.svg?branch=master)](https://travis-ci.org/01org/ciao)
[![Coverage Status](https://coveralls.io/repos/github/01org/ciao/badge.svg?branch=master)](https://coveralls.io/github/01org/ciao?branch=master)

Ciao is the "Cloud Integrated Advanced Orchestrator".  Its goal is
to provide an easy to deploy, secure, scalable cloud orchestration
system which handles virtual machines, containers, and bare metal apps
agnostically as generic workloads.  Implemented in the Go language, it
separates logic into "controller", "scheduler" and "launcher" components
which communicate over the "Simple and Secure Node Transfer Protocol
(SSNTP)".

[Controller](https://github.com/01org/ciao/blob/master/ciao-controller)
is responsible for policy choices around tenant workloads.

[Scheduler](https://github.com/01org/ciao/blob/master/ciao-scheduler)
implements a push scheduling, finding a first fit on cluster compute
nodes for a controller approved workload instance.

[Launcher](https://github.com/01org/ciao/blob/master/ciao-launcher)
abstracts the specific launching details for the different workload
types (eg: virtual machine, container, bare metal).  Launcher reports
compute node statistics to the scheduler and controller.  It also reports
per-instance statistics up to controller.

An additional set of componentry provides [ciao
network](https://github.com/01org/ciao/blob/master/networking)
connectivity for workload instances and insures tenant isolation.
Workloads (whether container or VM) are automatically placed in a unified
L2 network, one such network per tenant.

A [cli](https://github.com/01org/ciao/tree/master/ciao-cli) and
[webui](https://github.com/01org/ciao-webui) are available.

All ciao components communicate with each other via
[SSNTP](https://github.com/01org/ciao/blob/master/ssntp/README.md) using a
set of [payloads](https://github.com/01org/ciao/blob/master/payloads).

This GitHub repository contains documentation on the
various sub-components of ciao in their respective
subdirectories.  A comprehensive [ciao cluster setup
document](https://clearlinux.org/documentation/ciao-cluster-setup.html)
is also available.

The ciao development team can be reached via our [mailing
list](https://lists.clearlinux.org/mailman/listinfo/ciao-devel) and on IRC
in channel #ciao-project on [Freenode](https://freenode.net/kb/answer/chat).
