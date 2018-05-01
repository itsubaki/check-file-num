# check-fileagedir

# Help

```

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
