# readme

go 1.17

examples from:
- [《在Docker中体验Go1-17中的泛型》的源码](https://github.com/JasonkayZK/Go_Learn/tree/go-v1.17-rc-generic)

## run
```bash
go run -gcflags=-G=3 1_print/t.go
```

## FAQ
### `Cannot (yet) export a generic type`/`Cannot export a generic function (yet)`
当前generic不支持导出， 因此对应的generic var/function用大写字母开头的命名都会报该错.