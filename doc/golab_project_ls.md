## golab project ls

List all projects

### Synopsis


Get a list of all visible projects across GitLab for the authenticated user.

```
golab project ls [flags]
```

### Options

```
  -a, --archived            (optional) Limit by archived status
  -h, --help                help for ls
  -m, --membership          (optional) Limit by projects that the current user is a member of
  -o, --order_by string     (optional) Return projects ordered by id, name, path, created_at, updated_at, or last_activity_at fields. Default is created_at
      --owned               (optional) Limit by projects owned by the current user
      --search string       (optional) Return list of projects matching the search criteria
  -s, --simple              (optional) Return only the ID, URL, name, and path of each project
      --sort string         (optional) Return projects sorted in asc or desc order. Default is desc
      --starred             (optional) Limit by projects starred by the current user
      --statistics          (optional) Include project statistics
  -v, --visibility string   (optional) Limit by visibility public, internal, or private
```

### Options inherited from parent commands

```
      --config string   (optional) CURRENTLY NOT SUPPORTED config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab project](golab_project.md)	 - Manage projects

###### Auto generated by spf13/cobra on 27-Nov-2017