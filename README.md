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

## Development
If you want to help the development of the tool, you can do it by submitting
your PR.

### Functionalities
This tool is very simple yet, so it still doesn't have all the functionalities of
the CF web interface, but it helps get the job done. Below are the list of
the currently implemented functionalities and what I expect to implement in the
near future:

- [x] List zones
- [x] List records of a zone
- [x] Edit the value and TTL of a record
- [ ] Create a new record
- [ ] Delete a record
- [ ] Edit the whole record (type, name, CF proxy, etc...)
