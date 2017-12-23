## golab project edit

Edit project

### Synopsis


Updates an existing project.

```
golab project edit [flags]
```

### Options

```
      --ci_config_path string                              (optional) The path to CI config file
      --container_registry_enabled                         (optional) Enable container registry for this project
      --default_branch string                              (optional) master by default
      --description string                                 (optional) Short project description
  -h, --help                                               help for edit
  -i, --id string                                          (required) The ID or URL-encoded path of the project
      --import_url string                                  (optional) URL to import repository from
      --issues_enabled                                     (optional) Enable issues for this project
      --jobs_enabled                                       (optional) Enable jobs for this project
      --lfs_enabled                                        (optional) Enable LFS
      --merge_requests_enabled                             (optional) Enable merge requests for this project
  -n, --name string                                        (required) The name of the project
      --only_allow_merge_if_all_discussions_are_resolved   (optional) Set whether merge requests can only be merged when all the discussions are resolved
      --only_allow_merge_if_pipeline_succeeds              (optional) Set whether merge requests can only be merged with successful jobs
      --path string                                        (optional) Custom repository name for the project. By default generated based on name
      --public_jobs                                        (optional) If true, jobs can be viewed by non-project-members
      --request_access_enabled                             (optional) Allow users to request member access
      --resolve_outdated_diff_discussions                  (optional) Automatically resolve merge request diffs discussions on lines changed with a push
      --shared_runners_enabled                             (optional) Enable shared runners for this project
      --snippets_enabled                                   (optional) Enable snippets for this project
      --tag_list stringArray                               (optional) The list of tags for a project; put array of tags, that should be finally assigned to a project
      --visibility string                                  (optional) See project visibility level
      --wiki_enabled                                       (optional) Enable wiki for this project
```

### Options inherited from parent commands

```
      --ca-file string   (optional) provides a .pem file to be used in certificates pool for SSL connection
      --ca-path string   (optional) provides a directory with .pem certificates to be used for SSL connection
      --config string    (optional) golab config file (default is ./.golab.yml and $HOME/.golab.yml)
```

### SEE ALSO
* [golab project](golab_project.md)	 - Manage projects

