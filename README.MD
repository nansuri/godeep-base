# osp-be-gp Service
This is OSP Base

## Prequisites
#### Local Run
- go (of course)
- redis (use staging env)
- mariadb/mysql (use staging env)
- phpmyadmin (use on the staging env)

## Usage
#### Local Use
1. Clone this repository
2. Copy `.env.template` and create your own `.env` file and fill the config.
3. run the app using `go run main.go`

#### Production Use
1. Prepare all requisites
2. run `./build.sh`

## Structure

This is the current structure of the project

```
.
├── Dockerfile
├── README.MD
├── application
│   └── ..
├── build.sh
├── domain
│   ├── entity
│   │   └── ..
│   ├── enum
│   │   ├── ..
│   ├── repository
│   │   └── ..
│   └── service
├── go.mod
├── go.sum
├── infrastructure
│   ├── auth
│   │   ├── ..
│   ├── persistence
│   │   ├── ..
│   └── security
│       └── ..
├── interfaces
│   ├── fileupload
│   │   ├── ..
│   ├── middleware
│   │   └── ..
│   ├── routers
│   │   ├── ..
│   └── ..
├── main.go
└── utils
    ├── ..
```

## License
DANA Indonesia Copyright@2023