# Go click click click
Fast, Continuous, Automatic mouse click tool.

# Installation
```
go get github.com/walkingmask/go3click/cmd/go3click
```

# Usage
Get mouse location.

```
$ go3click loc
```

Click at 120,590, wait 200 milliseconds, click at 840,670 3 times.

```
$ go3click 1:120,590 w:200 3:840,670
```

Show help

```
$ go3click help
```

# Referenced
- [GoでCocoa APIを使う](http://unknownplace.org/archives/cgo-and-eventloop.html)
- [objecttive-cでのマウスイベント](http://scalpingroulett.mailorder-site.com/wordpress/archives/148)
- [[Objective-C]マウスイベント/キーボードイベントの作成](http://blog.springdawn.info/post/48712194308)
- [Code-Hex/battery](https://github.com/Code-Hex/battery/blob/master/battery_darwin.go)
- [BlueM/cliclick](https://github.com/BlueM/cliclick)
