bulkclone
===

Get repositories you want in bulk

## How to use

This tool requires YAML file like below:

```yaml
repos:
  - https://github.com/fooUser/barRepo
  - ...
```

After preparing that kinda config file,

```console
$ cat repos.yaml | bulkclone
```

## License

MIT

## Author

@b4b4r07
