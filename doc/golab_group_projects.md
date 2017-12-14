## golab group projects

List a group's projects

### Synopsis


Get a list of projects in this group. When accessed without authentication, only public projects are returned.

```
golab group projects [flags]
```

### Options

```
      --archived            (optional) Limit by archived status
  -h, --help                help for projects
      --id string           (required) The ID or URL-encoded path of the group owned by the authenticated user
      --order_by string     (optional) Return projects ordered by id, name, path, created_at, updated_at, or last_activity_at fields. Default is created_at
      --owned               (optional) Limit by projects owned by the current user
      --page int            (optional) Page of results to retrieve
      --per_page int        (optional) The number of results to include per page (max 100)
      --search string       (optional) Return list of authorized projects matching the search criteria
      --simple              (optional) Return only the ID, URL, name, and path of each project
      --sort string         (optional) Return projects sorted in asc or desc order. Default is desc
      --starred             (optional) Limit by projects starred by the current user
      --visibility string   (optional) Limit by visibility public, internal, or private
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab group](golab_group.md)	 - Manage Gitlab Groups

