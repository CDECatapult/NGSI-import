# NGSI-import
Script to import NGSI data sources as Data Source Specification on the IoT Data Marketplace from a list.

[![License badge](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

+ [Introduction](#def-introduction)
+ [How to Build](#def-build)
+ [Configuration](#def-conf)
+ [How to Use](#def-use)
+ [License](#def-license)

---

<br>

<a name="def-introduction"></a>
## Introduction

This project is part of the EU H2020 [SynchroniCity](https://synchronicity-iot.eu) project and it is an add on for the [SynchroniCity IoT Data Marketplace](https://github.com/caposseleDigicat/SynchroniCityDataMarketplace).

- You will find the source code of this project in GitHub [here](https://github.com/caposseleDigicat/NGSI-import)

Thanks to this component you will be able to import data sources specification directly on the IoT Data Marketplace. This project reads from a file a list of NGSI pairs [Entity Type] 
[Entity ID] and creates the respective data source specification on the marketplace. 

<a name="def-build"></a>
## How to Build

Inside the `bin` folder you will find 3 different version already compiled for `Mac OSX`, `Windows`, and `Linux`. If you wish to recompile it, just read the following instructions. 

Requirements: [Go Programming Language](https://golang.org/doc/install)

To build the binary for your `OS` and `architecrure` just run:

```
GOOS=[OS] GOARCH=[ARCH] go build main.go utility.go
```

where [OS] and [ARCH] are your Operating System and Architecture respectively. Some examples for OSX, Windows and Linux (with 32bit architecture) are:

```
GOOS=darwin GOARCH=386 go build main.go utility.go

GOOS=windows GOARCH=386 go build main.go utility.go

GOOS=linux GOARCH=386 go build main.go utility.go
```

These commads will create a `main` (`main.exe` for windows) executable file. For your reference, you can find a similar set of commands inside the script `build.sh`.

## License

The MIT License
 
Copyright (C) 2018 Digital Catapult.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
