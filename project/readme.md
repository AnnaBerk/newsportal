~/go/bin/gentool -db postgres -dsn "host=localhost user=postgres password=mypassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai" -tables "categories,news,statuses,tags"

gentool -c "./gen.tool"