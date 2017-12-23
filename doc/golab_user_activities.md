## golab user activities

Get user activities (admin only)

### Synopsis


Get the last activity date for all users, sorted from oldest to newest.

The activities that update the timestamp are:

* Git HTTP/SSH activities (such as clone, push)
* User logging in into GitLab

By default, it shows the activity for all users in the last 6 months, but this can be amended by using the from parameter.

```
golab user activities [flags]
```

### Options

```
      --from string   (optional) Date string in the format YEAR-MONTH-DAY, e.g. 2016-03-11. Defaults to 6 months ago.
  -h, --help          help for activities
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user](golab_user.md)	 - Manage Gitlab users

