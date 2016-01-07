# convey-actions

convey-actions is server for action testing using customized [goconvey](https://github.com/smartystreets/goconvey) dashboard

this is a really dirty implementation just to pass the convey api compability, and it works

![Dashboard](https://bitbucket.org/actions/convey-dashboard/raw/master/screenshot.png)

## functionality

* Stories
* Coverage
* Tests watching
* Run tests from dashboard

## watching

watching for changes in tests is done using `github.com/go-fsnotify/fsnotify`

# Authors

This package is a for testing [actions](https://bitbucket.org/actions).
It is customized [GoConvey](https://github.com/smartystreets/goconvey) dashboard.

GoConvey was originally created by [SmartyStreets](https://github.com/smartystreets); in particular:

 - [Michael Whatcott](https://github.com/mdwhatcott)
 - [Matt Holt](https://github.com/mholt)

Modifications:
 - [Łukasz Kurowski](https://github.com/crackcomm)
