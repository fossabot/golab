## golab login

Login to a Gitlab server

### Synopsis


Log in to a Gitlab server with your username and password.

```
golab login [flags]
```

### Options

```
  -h, --help              help for login
  -s, --host string       (required) URL to Gitlab server eg. http://gitlab.org
  -p, --password string   (optional) password, if not given, you'll be prompted interactively
  -u, --user string       (required) username
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab](golab.md)	 - Gitlab CLI written in Go

