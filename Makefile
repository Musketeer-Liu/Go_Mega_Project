.PHONY: clean upload-vps

# Put your server ip or nickname here
host = musketeer

clean:
    @echo "clean..."
    @rm -rf app/
    @echo "success clean"
build-db:
    @echo "run build db..."
    @env GOOS=linux go build -o app/db cmd/db_init/main.go
    @echo "success build db"
build-server:
    @echo "begin build server..."
    @env GOOS=linux go build -o app/server main.go
    @echo "success build server"
upload-vps: clean build-db build server
    @echo "begin run upvs upload..."
    @echo "Host: $(host)"
    @scp -r app root@($host):go_mega
    @scp config.yml root@$(host):go_mega/
    @scp -r templates root@$(host):go_mega/templates
    @echo "success vps upload"