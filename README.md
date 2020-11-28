# go-resource-monitor
Resource monitor displays information about the use of hardware (CPU, memory, disk, and network) resources in real time.

![image](https://user-images.githubusercontent.com/10591350/100493393-d0ded380-3179-11eb-9351-a6ff3aff1dfd.png)

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