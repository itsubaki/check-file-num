# check-fileagedir

# Help

```
$ check-file-num -h
Usage:
  check-file-num [OPTIONS]

Application Options:
  -b, --base=         the base directory(required)
  -w, --warning-num=  warning if more number of files(count) (default: 10)
  -c, --critical-num= critical if more number of files(count) (default: 100)
  -d, --debug         debug print

Help Options:
  -h, --help          Show this help message
```

# Install

```
sudo mkr plugin install itsubaki/check-file-num@v0.1
```

```
# /etc/mackerel-agent/mackerel-agent.conf
[plugin.checks.fnum_td-agent_buffer]
command = "/opt/mackerel-agent/plugins/bin/check-file-num -b /var/log/td-agent/buffer/ -w 10 -c 100"
```
