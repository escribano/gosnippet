# Go Snippet

This is a sample web app designed to help show a real world deployable quick prototype of a web app to those wishing to use Go and wondering how to knock something out in a few hours.

You are encouraged to fork this, copy it, submit pull requests for improvements (that help other people get started or fix bugs).

For that reason, rather than create a complex app with caveats like "you shouldn't do this in live" or "// TODO: handle errors", this example is going to be an extremely simple app that does show you all of the error handling and what to do in live, and it makes some excuses to be sure that we demonstrate each type of thing you might want to do (database integration, object caching, scheduled tasks, config file for deployment, etc).

This app is not a framework... it's just a pulling together of libraries with a set of conventions that will help developers with knowledge of Rails or Django to make the jump to Go.

This app uses:
* [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
* [cron](https://github.com/robfig/cron)
* [libpq](https://github.com/lib/pq)
* [memcache](https://github.com/bradfitz/gomemcache)
* [toml-go](https://github.com/laurent22/toml-go)

What this app does... not much, it's just a GitHub Gist type thing... post code, have it stored in the database, have it formatted prettily when it is displayed. It will only keep code for 1 month, after which it will auto-delete it (needed to find some excuse to put scheduled tasks in).

# Dependencies

You will need to have installed: memcached, postgresql. Assumption that you know how to install that stuff and can create a database using the SQL file included in this project.

It is also assumed that you have at least Go1.1 and that you've setup your $GOPATH to a place you're happy to work from (mine is ~/Dev/Code feel free to use whatever you want).

Then:
```
go get github.com/bradfitz/gomemcache/memcache
go get github.com/gorilla/mux
go get github.com/laurent22/toml-go/toml
go get github.com/lib/pq
go get github.com/robfig/cron
```

Which will have pulled stuff into $GOPATH/src/* and then built the packages and shoved those in $GOPATH/pkg/*

# Pre-Requisites

GoSnippet expects to find a configuration file either in the directory it is run from or ideally at /etc/gosnippet/gosnippet.toml with the following configuration keys (values obviously change per environment):

```
[listen]

interface = "" # Empty string for all interfaces
port = 8080

[database]

host = "localhost"
port = 5432
database = "gosnippet"
username = "gosnippet"
password = "gosnippet"

[memcached]

host = "127.0.0.1"
port = 11211

[directories]

# Relative to the directory gosnippet is run from, trailing slash
# is important. Ideally these directories are in /srv/gosnippet/*/
# in a prod environment. Note the trailing slash is required.
templates = "./views/"
static = "./static/"
```
