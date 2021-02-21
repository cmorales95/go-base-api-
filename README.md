# API Rest Base for GO

This is API Rest base with Go based on basic DDD structure

## Getting Started

The following instructions help you to install the app for development purpose

### Prerequisites

#### a. Golang

Go is the principal language of this project, the official website of go https://golang.org/ and click in __download__, then chose your operative system and install, You find the instructions of installation in the page (See *2. Go Install* and then select your O.S and read the instructions)

#### b. Docker

Docker helps us to manage all the environment needed to run the app (include BD and the app) please visit the following links:

- Install Docker:

  - Linux (Ubuntu and others distro): https://docs.docker.com/engine/install/ubuntu/
  - Windows: https://docs.docker.com/docker-for-windows/install/
  - Mac: https://docs.docker.com/docker-for-mac/install/

- Post-Installation of Docker (Only Linux): https://docs.docker.com/engine/install/linux-postinstall/

- Install Docker Compose (Only Linux): https://docs.docker.com/compose/install/

*Note: Windows and Mac already count with docker user and compose included in the installation stand-alone of docker*

*Recommendation:* if your O.S is Windows we recommend turning on the option of __wsl2 (Windows Subsystem for Linux)__, it allows to work with virtualization of Linux into the windows (no dual boot). Sometimes, you could experiment some inconveniences with docker on Windows but with wsl2 that mistakes can be minimized in a high percentage, Please, follow the next link to know how to install wsl2 in Windows and how to run an ubuntu light image over wsl2: https://www.omgubuntu.co.uk/how-to-install-wsl2-on-windows-10


__Important:__ please __DO NOT INSTALL__  databases engine (as Postgres) in your local machines, the project manage automatically the bd from docker images

#### c. Make (Only for Windows Without WSL2)

Make is a utility command to allow the execution of key and short commands from a *Makefile*, go to the website https://stat545.com/make-windows.html and follow the instructions of the installation


### Usage

Once finished the installation you count with command key in the file *Makefile* to run the application and many more:

#### Build the app
To build the containers of the app and bd (only for a production environment or functional local tests)
```
make docker-build-run
```

#### Run only services
Start all the services needed to run the app (only BD by the moment), this command is recommended for the local environment, also we recommend running the app by yourself and avoid using the key command for building the app, this option is created to a faster preparation of the environment for developers
```
make docker-run-services
```

### Stop the services and destroy containers
Seek and destroy all the services started by another key command, this command deletes the containers physically. For that reason, the data saved in BD will be deleted as well
```
make docker-down-all
```


### Build With
- Golang
- Postgres
- Docker
- Git

### Find Us
- cmorales95

### Versioning
We are using *tags* as versioning method of the project


### Authors
- Cristian Morales: crisdavidmm95.dev@gmail.com


### License
Open Source My friend, fork and work!
