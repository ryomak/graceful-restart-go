# graceful-restart-go
process削除

cat server1.pid | xargs kill -HUP
