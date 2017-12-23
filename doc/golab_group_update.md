## golab group update

Update group

### Synopsis


Updates the project group. Only available to group owners and administrators.

```
golab group update [flags]
```

### Options

```
      --description string       (optional) The description of the group
  -h, --help                     help for update
      --id int                   (required) The ID of the group
      --lfs_enabled              (optional) Enable/disable Large File Storage (LFS) for the projects in this group
      --name string              (optional) The name of the group
      --path string              (optional) The path of the group
      --request_access_enabled   (optional) Allow users to request member access.
      --visibility string        (optional) The visibility level of the group. Can be private, internal, or public.
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab group](golab_group.md)	 - Manage Gitlab Groups

