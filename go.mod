module github.com/untitled-discord-star-project/backend

go 1.22.1

require (
	github.com/a-h/templ v0.2.648
	github.com/gocql/gocql v1.6.0
	github.com/google/uuid v1.6.0
	github.com/scylladb/gocqlx/v2 v2.8.0
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/oklog/ulid/v2 v2.1.0
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.13.0
