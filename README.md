Gitlab CLI written in Go [![Build Status](https://travis-ci.org/michaellihs/golab.svg?branch=master "Travis CI status")](https://travis-ci.org/michaellihs/golab)
======================== 

This project provides a Command Line Interface (CLI) for Gitlab written in Go. The project uses the [go-gitlab client](https://github.com/xanzy/go-gitlab) for the communication with Gitlab.

It allows you to run Gitlab administration tasks from your command line. Examples:

* create a user

   ``` bash
   golab user create --email username@company.com --username username --password 12341234 --name "User McArthur" --skipConfirmation
   ```

* modify a user

   ``` bash
   golab user modify -i 41 --admin true
   ```

* create a new project / repository

   ``` bash
   golab project create -g my-group -n my-project
   ```

* add an ssh key for a user

   ``` bash
   golab user ssh-keys add --key "`cat ~/.ssh/id_rsa.pub`" --title "my dsa key"
   ```

For a complete documentation of features, check the [generated documentation](doc/golab.md)

Setup & Configuration
---------------------

Install the CLI tool with

    go get github.com/michaellihs/golab

First create a Gitlab [access token for your user](https://docs.gitlab.com/ce/user/profile/personal_access_tokens.html) in Gitlab (most likely a admin user).

Create a file `.golab.yml` in either `~/` or the directory you want to use golab with the following content:

    ---
    url: "http(s)://<YOUR GITLAB URL>"
    token: "<YOUR PERSONAL ACCESS TOKEN>"

Test your configuration - e.g. by running `golab project` to get a list of projects from your Gitlab server.


Development
===========

API Debugging
-------------

Run `curl` requests against the API:

    curl --header "PRIVATE-TOKEN: FqBiTTJ4oRPdskWDTktr" -H "Content-Type: application/json" -X PUT -d '{"admin": true}' http://localhost:8080/api/v4/users/41


Build and run the application
-----------------------------

There is a `makefile` included that can build and test the application and render the automatically generated documentation:

*  `make` - build the application

*  `make test` - run the tests

* `make gendoc` - render the documentation
    

Gitlab Docker Image
-------------------

For local development, you can use a [Gitlab Docker image](https://docs.gitlab.com/omnibus/docker/README.html). There are some pitfalls, when using Gitlab Docker image on a Mac:

* You cannot properly mount the `/var/opt/gitlab` directory due to issues with NFS mounts on Mac
* The ssh port `22` is already in use on the Mac, if a ssh server is running

Therefore adapt the provided run command to the following:

    sudo docker run --detach \
        --hostname gitlab.example.com \
        --publish 443:443 --publish 80:80 --publish 8022:22 \
        --name gitlab \
        --volume /tmp/gitlab/config:/etc/gitlab \
        --volume /tmp/gitlab/logs:/var/log/gitlab \
        gitlab/gitlab-ce:latest

Afterwards you can start the (existing) container with:

    sudo docker start gitlab


TODOs
=====

Support multiple Targets
------------------------

`golab login` should take a parameter `-e` that sets an environment which we can later on select with `-e` in each command or read from `$golab_env`.

Therefore we also need to change the structure of the `.golab.yml` like this:

    ---
    gitlab.com:
      url: "https://gitlab.com"
      token: "gitlab_com_token"

    localhost:
      url: "http://localhost:12345"
      token: "localhost_token"

This allows working with multiple Gitlab servers at the same time.


Further Resources
=================

* [Gitlab API Documentation](https://docs.gitlab.com/ee/api/README.html)
* [Cobra Library (Go CLI Library)](https://github.com/spf13/cobra)
* [Viper Library (Go Flags Library)](https://github.com/spf13/viper)
* [Sling (Go HTTP Library)](https://github.com/dghubble/sling)
* [Ginkgo (Go Testing Library)](https://onsi.github.io/ginkgo/)
* [GoMock (Go Mocking Library)](https://github.com/golang/mock)
* [go-gitlab (Go Gitlab Library)](https://github.com/xanzy/go-gitlab)
