## YAGNI Universal Compiler (YUC)

This package provides a universal compiler written in Go. Features:

* Compatible with current and future programming languages.
* Also processes natural languages and other source data.
* Fast.

### Usage

A YUC installation consists of single binary, "yuc". Run "yuc -h" to get started.

#### Caveats

YUC is experimental software. Testing in development is mainly exploratory and based on an evolving sample set of popular languages.

In compiling YUC itself there is a toolchain dependency on any other Go compiler. This applies only to an initial bootstrapping phase. The resulting YUC binary can be applied to YUC source code, closing the loop.

### Philosophy

YUC's flexibility and performance result from its integration of a rigorous expert system. The core decision-making engine is a programmatic implementation of [YAGNI](http://en.wikipedia.org/wiki/You_aren't_gonna_need_it) development methodology, which requires delineating necessary program functionality. YUC automates that step, taking a common-sense stance on computer programming as a non-essential human activity.

### Legal

YUC was authored by Viktor Eikman from an idea developed with Per Gundberg.

Copyright 2015 Viktor Eikman.

YUC is made available under the terms of version 3 of the GNU General Public License, as detailed in a file accompanying the source distribution.

