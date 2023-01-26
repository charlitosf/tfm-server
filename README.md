# tfm-server

Back-end REST server of the Password Manager project, an implementation of the server side of [The Integrally Secure Storage Protocol](http://hdl.handle.net/10045/124732). The client is located at [charlitosf/tfm-client](https://github.com/charlitosf/tfm-client).

The project follows the [MVC](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller) architectural pattern.

## Setup

Required steps to put the project up and running.

### Prerequisites

- make
- git
- go 1.18^

### Apache HBase

This project uses the distributed database Apache HBase. Therefore, the first step in the project setup is to download HBase from its [website](https://hbase.apache.org/downloads.html) (the "bin" option).

Then, follow the [documentation](https://hbase.apache.org/book.html#quickstart) to set up the database at least as a standalone server (more complex setups also work). Although, to sum up, the steps are:

```bash
# Decompress the tar file
tar xzvf downloaded-file.tar.gz
cd hbase-folder
```

Modify the `conf/hbase-env.sh` file so that it includes an export of the `JAVA_HOME` variable pointing to your Java 8 installation.

```bash
# Start the database
./bin/start-hbase.sh
```

```bash
# Create the necessary tables
./bin/hbase shell
> create 'users', 'data' # Users table, data column family
> create 'passwords', 'data' # Passwords table, data column family
> create 'files', 'data' # Files table, data column family
```

### Server setup

Once the database is running and has the necessary tables created:

```bash
# Clone the server repository
git clone https://github.com/charlitosf/tfm-server

# Set up the .env file
cp .env.copy .env
```

Then, modify the .env file according to your needs and run the project with make.

```bash
make

# Or or only creating the server.exe file and not executing it
make build
```

Make will automatically compile the whole server, create the server.exe file (compatible with Linux), and run it.

## Project Organization

- [api](https://github.com/charlitosf/tfm-server/tree/master/api): API definition using OpenApi 3.0.2
- [cmd/server](https://github.com/charlitosf/tfm-server/tree/master/cmd/server): Main file of the project
- [internal](https://github.com/charlitosf/tfm-server/tree/master/internal): Code specific to this project
  - [controllers](https://github.com/charlitosf/tfm-server/tree/master/internal/controllers): Functions that handle the calls to the API endpoints.
  - [crypt](https://github.com/charlitosf/tfm-server/tree/master/internal/crypt): Project-specific cryptographic functions and utilities.
  - [dataaccess](https://github.com/charlitosf/tfm-server/tree/master/internal/dataaccess): Go functions that communicate with the database.
  - [jwt](https://github.com/charlitosf/tfm-server/tree/master/internal/jwt): Code specific to the creation, verification, etc. of the JWT tokens used in the server.
  - [middleware](https://github.com/charlitosf/tfm-server/tree/master/internal/middleware): Functions that are executed to verify common requirements of some REST endpoints.
  - [services](https://github.com/charlitosf/tfm-server/tree/master/internal/services): Business-layer code.
- [pkg](https://github.com/charlitosf/tfm-server/tree/master/pkg): Code that can be shared with other projects.
  - [httptypes](https://github.com/charlitosf/tfm-server/tree/master/pkg/httptypes): Go structures that hold the data transferred between the client and the server.
  - [stringutilities](https://github.com/charlitosf/tfm-server/tree/master/pkg/stringutilities): String manipulation functions.
