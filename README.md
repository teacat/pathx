

dir := "foo/bar"
dir2 := hello/world/"

pathx.Join(dir, dir2) // foo/bar/hello/world
pathx.JoinDir(dir, dir2) // foo/bar/hello/world/

file := "my/dir/main.go"

pathx.Name(file) // main
pathx.Ext(file)  // go

wtf := "foo/bar/hello/foo/bar/world"
pathx.Merge() // foo/bar/hello/world

```
a := "foo/bar/world"
pathx.Contains("/bar/world/")
```

pathx.Current()

pathx.Current(pathx.Join("a", "b"))

pathx.CurrentDir