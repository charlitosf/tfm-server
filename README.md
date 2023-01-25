# tfm-server

Back-end REST server of the Password Manager project, an implementation of the server side of [The Integrally Secure Storage Protocol](http://hdl.handle.net/10045/124732). The client is located at [charlitosf/tfm-client](https://github.com/charlitosf/tfm-client).

The project follows the [MVC](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller) architectural pattern.

## Setup

```bash
make
```

Make will automatically compile the whole server and create the server.exe file (compatible with Linux).

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
