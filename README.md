Go Sort
=======

The purpose of this project is to implement some functionalities of [GNU sort](https://man7.org/linux/man-pages/man1/sort.1.html).

Build
-----

To build this project run the following command:

```bash
go build cmd/sort/sort.go
```

Usage
-----

Actually there is only one parameter supported `-r` that will reverse the sorting order, this parameter could be in any position.
Is possible to read a file or from stding:

```bash
./sort /etc/passwd
cat /etc/passwd | ./sort
```
