# `godocv`

`godocv` wraps godoc in `/usr/bin/less` and tries to prepend the package name
with `./vendor/` (if that path exists). This means that `godocv` shows the local
vendored package docs if they exist instead of the `$GOPATH` docs. Maybe one day
this will not be necessary: https://github.com/golang/go/issues/21939
