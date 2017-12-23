## golab group-members edit

Edit a member of a group or project

### Synopsis


Updates a member of a group or project.

  Access Levels:

	10 = Guest Permissions
	20 = Reporter Permissions
	30 = Developer Permissions
	40 = Master Permissions
	50 = Owner Permissions

```
golab group-members edit [flags]
```

### Options

```
  -a, --access_level int    (required) a valid access level
  -e, --expires_at string   (optional) expiry date of membership (yyy-mm-dd)
  -h, --help                help for edit
  -i, --id int              (required) id of group to change membership for
  -u, --user_id int         (required) id the user to change membership for
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab group-members](golab_group-members.md)	 - Access group members

