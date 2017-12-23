## golab project hooks delete

Delete project hook

### Synopsis


Removes a hook from a project. This is an idempotent method and can be called multiple times. Either the hook is available or not.

```
golab project hooks delete [flags]
```

### Options

```
  -h, --help          help for delete
      --hook_id int   The ID of the project hook
  -i, --id string     The ID or URL-encoded path of the project
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab project hooks](golab_project_hooks.md)	 - Manage project hooks.

