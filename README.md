# cfdnsmanager
CLI tool to manage DNS records on CloudFlare

## Install
To install it, use:
```
go get -u github.com/DennyLoko/cfdnsmanager
```

The tool reads a configuration file from `$HOME/cfdnsmanager.yaml` or
`/etc/cfdnsmanager.yaml` with your CloudFlare credentials. You can find a sample
of the credentials at [`config.yaml.sample`](./blob/master/config.yaml.sample).
