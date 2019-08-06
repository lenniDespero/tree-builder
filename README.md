### Утилита tree.

It's print directories and files(if flag -f enable)

For tests run `go test -v` in folder

```
$ go test -v
=== RUN   TestTreeFull
--- PASS: TestTreeFull (0.00s)
=== RUN   TestTreeDir
--- PASS: TestTreeDir (0.00s)
PASS
ok      coursera/homework/tree     0.127s
```

```
go run main.go . -f
├───main.go (1881b)
├───main_test.go (1318b)
└───testdata
	├───project
	│	├───file.txt (19b)
	│	└───gopher.png (70372b)
	├───static
	│	├───css
	│	│	└───body.css (28b)
	│	├───html
	│	│	└───index.html (57b)
	│	└───js
	│		└───site.js (10b)
	├───zline
	│	└───empty.txt (empty)
	└───zzfile.txt (empty)
go run main.go .
└───testdata
	├───project
	├───static
	│	├───css
	│	├───html
	│	└───js
	└───zline
```