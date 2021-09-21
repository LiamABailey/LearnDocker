# LearnDocker
#### Implementation of A Simple API Using Go, Gin, MongoDB and Docker

System utilizes the following to store and serve information about Fruits:
- MongoDB: Backend database
- [Gin](https://github.com/gin-gonic/gin): Fast, easy to use Go web framework  
- Docker: Used to separate the API and Mongo backend. The docker-compose file in this project spins up two containers, one for each component.

Code is organized in the following way:
- /cmd: Contains `main.go`, which starts the services
- /pkg: API components
  - `mongointeracts.go`: Functions managing interactions with Mongo
  - `testapi.go`: API specification bridging calls with `mongointeracts`
  - `teststructs.go`: Definitions of relevant data structures
- /mongodb: MongoDB docker file, sample data


API intentionally supports few endpoints; this was written as an exercise in getting Docker working, not in building a meaningful service.
