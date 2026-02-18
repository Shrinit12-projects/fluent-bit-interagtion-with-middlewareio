```pgsql
Go App → writes structured logs → log file
Fluent Bit → reads log file → sends to Middleware Agent
Middleware Agent → sends to Middleware.io
```

---

## Folder Structure
```pgsql
go-fluentbit-middleware/
│
├── app/
│   ├── main.go
│   └── go.mod
│
├── fluent-bit/
│   └── fluent-bit.conf
│
├── middleware-agent/
│   └── (no config needed)
│
├── logs/
│   └── app.log
│
└── docker-compose.yml
```
