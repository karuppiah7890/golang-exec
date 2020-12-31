# Story

```bash
$ go test -v ./...
=== RUN   TestSomething
--- FAIL: TestSomething (0.00s)
panic: fork/exec ../test/script: permission denied [recovered]
        panic: fork/exec ../test/script: permission denied

goroutine 18 [running]:
testing.tRunner.func1.1(0x112fd80, 0xc000092330)
        /usr/local/Fish/Barrel/go/1.15.6/go/src/testing/testing.go:1072 +0x30d
testing.tRunner.func1(0xc000082600)
        /usr/local/Fish/Barrel/go/1.15.6/go/src/testing/testing.go:1075 +0x41a
panic(0x112fd80, 0xc000092330)
        /usr/local/Fish/Barrel/go/1.15.6/go/src/runtime/panic.go:969 +0x1b9
_/Users/karuppiahn/oss/github.com/karuppiah7890/windows-golang-exec-bash-script/pkg.Something(0xc0000a6080, 0xe)
        /Users/karuppiahn/oss/github.com/karuppiah7890/windows-golang-exec-bash-script/pkg/cmd.go:13 +0xfd
_/Users/karuppiahn/oss/github.com/karuppiah7890/windows-golang-exec-bash-script/pkg.TestSomething(0xc000082600)
        /Users/karuppiahn/oss/github.com/karuppiah7890/windows-golang-exec-bash-script/pkg/cmd_test.go:9 +0xaa
testing.tRunner(0xc000082600, 0x115aab8)
        /usr/local/Fish/Barrel/go/1.15.6/go/src/testing/testing.go:1123 +0xef
created by testing.(*T).Run
        /usr/local/Fish/Barrel/go/1.15.6/go/src/testing/testing.go:1168 +0x2b3
FAIL    _/Users/karuppiahn/oss/github.com/karuppiah7890/windows-golang-exec-bash-script/pkg     0.536s
FAIL
```

```bash
$ chmod +x test/script
```

```bash
$ go test -v ./...
=== RUN   TestSomething
[68 111 105 110 103 32 115 111 109 101 32 114 97 110 100 111 109 32 119 111 114 107 46 46 46 10 68 111 110 101 33 32 58 41 10]
--- PASS: TestSomething (0.01s)
PASS
ok      _/Users/karuppiahn/oss/github.com/karuppiah7890/windows-golang-exec-bash-script/pkg     0.074s
```
