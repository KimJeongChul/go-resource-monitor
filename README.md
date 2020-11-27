# go-resource-monitor
Resource monitor displays information about the use of hardware (CPU, memory, disk, and network) resources in real time.

![image](https://user-images.githubusercontent.com/10591350/100466717-b6264380-3114-11eb-9cf7-f400d70548fa.png)

### Start
```shell
go build
./go-resource-monitor
```

### Config file
config.json
```json
{
    "port": {HTTP_SERVER_PORT},
    "period": {MONITORING_PERIOD},
    "logPath": {LOG_FILE_PATH}
}
```