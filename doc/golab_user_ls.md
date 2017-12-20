## golab user ls

List users

### Synopsis


Get a list of users.

```
golab user ls [flags]
```

### Options

```
      --active                          (optional) Filter users based on state active
      --blocked                         (optional) Filter users based on state blocked
      --created_after string            (optional) Search users by creation date time range, e.g. 2001-01-02T00:00:00.060Z (admin only)
      --created_before string           (optional) Search users by creation date time range, e.g. 2001-01-02T00:00:00.060Z (admin only)
      --custom_attribute_key string     (optional) Filter by custom attribute key (admin only)
      --custom_attribute_value string   (optional) Filter by custom attribute value (admin only)
      --extern_uid string               (optional) Lookup users by external UID and provider (admin only)
      --external                        (optional) Search for users who are external (admin only)
  -h, --help                            help for ls
      --page int                        (optional) Page of results to retrieve
      --per_page int                    (optional) The number of results to include per page (max 100)
      --provider string                 (optional) Lookup users by external UID and provider (admin only)
      --search string                   (optional) Search for users by email or username (admin only)
      --username string                 (optional) Lookup users by username (admin only)
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab user](golab_user.md)	 - Manage Gitlab users

