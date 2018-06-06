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

This tool is part of the EU H2020 [SynchroniCity](https://synchronicity-iot.eu) project and it is an add on for the [SynchroniCity IoT Data Marketplace](https://github.com/caposseleDigicat/SynchroniCityDataMarketplace).

- You will find the source code of this tool in GitHub [here](https://github.com/caposseleDigicat/NGSI-import)

Thanks to this tool you will be able to import data sources specification directly on the IoT Data Marketplace. This tool reads from a file a list of NGSI pairs [Entity Type] 
[Entity ID] and creates the respective data source specification on the marketplace. 

<a name="def-build"></a>
## How to Build

Inside the `bin` folder you will find 3 different version already compiled for `Mac OSX`, `Windows`, and `Linux`. If you wish to recompile it, just read the following instructions. 

Requirements: [Go Programming Language](https://golang.org/doc/install)

To build the binary for your `OS` and `architecture` just run:

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

<a name="def-conf"></a>
## Configuration

This tool requires a configuration file `config.json` to be filled as:

```
"marketplace_url": "",
"proxy_url": "",
"appplication_id": "",
"marketplace_username": "",
"marketplace_oauth2_token": "",
"brand":"",
"fiware_service":""
```

- `marketplace_url`: the url of the marketplace (e.g, http://marketplace.eu:8004).
- `proxy_url`: the url of the PEP proxy used to protect the data sources (e.g, http://wilmaPepProxy.eu:7000).
- `application_id`: the OAuth2 client ID of the Orion Context Broker registered on the Identity Management (e.g, 53626045d3bd4f8c84487f77944fa586).
- `marketplace_username`: the Username of the user that will be considerd the data provider for the imported data sources on the marketplace (e.g, mario). Please note that this user must have the role `data_provider` for the Orion Context Broker application on the IdM.
- `marketplace_auth2_token`: the Access Token of the user above after authenticating on the IoT Data Marketplace portal (e.g, oykjWSK32a3zQils7et9cD4FPeNpsI). You can find this token under the section `Settings` on the IoT Data Marketplace portal.
- `marketplace_url`: the brand used to fill the data source specifications in the marketplace (e.g, Manchester).
- `fiware_service`: the Fiware-Service used on the Orion Context Broker while registering the NGSI entities (e.g, manchester). It can be empty if not used.

As an example:

```
"marketplace_url": "http://proxy.docker:8004",
"proxy_url": "http://wilma.docker:7000",
"appplication_id": "53626045d3bd4f8c84487f77944fa586",
"marketplace_username": "mario",
"marketplace_oauth2_token": "oykjWSK32a3zQils7et9cD4FPeNpsI",
"brand":"My Brand",
"fiware_service":""
```

Please note that the configuration file must be saved in the same directory of the executable.

<a name="def-use"></a>
## How to Use

Before using the tool, you need to create a file containing the list of data sources that you wish to upload as data source specification on the IoT Data Marketplace.
This file should be created as list of pairs [Entity Type] [Entity ID], where [Entity Type] is mandatory while [Entity ID] is optional. When using only [Entity Type], a data source specification for all the entities of type [Entity Type] will be created. For your reference, the following is an example of the content of such file:

```
type1
type2   entity2
type2   entity3
type3   entity4
```

This list will allow the creation of 4 different data source specifications with the following names and urls:  

1. name: `type1` and url: `http://wilma.docker:7000/v2/entities?type=type1`
2. name: `type2 : entity2` and url: `http://wilma.docker:7000/v2/entities?type=type2&id=entity2`
2. name: `type2 : entity3` and url: `http://wilma.docker:7000/v2/entities?type=type2&id=entity3`
2. name: `type3 : entity4` and url: `http://wilma.docker:7000/v2/entities?type=type3&id=entity4`

After creating this file, you can run the tool and pass the path of the file (e.g., ./datasourcelist.dat) as argument:

```
./bin/import ./datasourcelist.dat
```

You will see an output similar to:

```
 ----------------------------------
|           NGSI Import            |
 ----------------------------------

The following data sources will be imported:
type1
type2 : entity2
type2 : entity3
type3 : entity4
WARNING: Are you sure? (yes/no)
```

By typing `no` the program will terminate while by typing `yes` the tool will import the respective data sources specification on the IoT Data Marketplace.
Once done, you will find the imported data source specifications under `My stock - Data source specifications` on the IoT Data Marketplace portal, with status `Active`. If you wish to create offerings with those, just set their status to `Launched`. 
## License

The MIT License
 
Copyright (C) 2018 Digital Catapult.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
