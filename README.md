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
