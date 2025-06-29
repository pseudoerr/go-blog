# Go-blog -- a lightweight platform for writing with markdown support. 

## Project layout 
 
├── README.md
├── cmd
│   └── go-blog
│       └── main.go
├── config
│   └── local.yml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── db
│   │   └── migrations
│   │       └── 0001_create_tables.up.sql
│   ├── handler
│   ├── models
│   │   └── models.go
│   ├── repository
│   │   ├── post_repository.go
│   │   └── repository.go
│   └── service
├── pkg
│   └── auth
└── web
    ├── static
    └── templates
