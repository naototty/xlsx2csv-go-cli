# xlsx2csv-go-cli

A simple program that converts Microsoft XLSX files into CSV files.

Run `build.sh` to build the executable using Docker.


## original is 

this CLI is based on https://github.com/colinmollenhour/xlsx2csv
and original is https://github.com/tealeg/xlsx2csv

originalを改変したものです

## how to use 

### usage

show usage;

```shell
MacBook-Pro:testdata chroum$ ./xlsx2csv-go-darwin64
./xlsx2csv-go-darwin64
	dumps the given xlsx file's chosen sheet as a CSV,
	with the specified delimiter, into the specified output.

Usage:
	./xlsx2csv-go-darwin64 [flags] <xlsx-to-be-read>
  -d string
    	Delimiter to use between fields(TAB is tab) (default ",")
  -i int
    	Index of sheet to convert, zero based
  -o string
    	filename to output to. -=stdout (default "-")
```
### convert to stdout

example) set delimiter to '|'

```
MacBook-Pro:testdata chroum$ ./xlsx2csv-go-darwin64 -d '|' testfile.xlsx
Foo|Bar
Baz |Quuk
```

example) set delimiter to ','

```
MacBook-Pro:testdata chroum$ ./xlsx2csv-go-darwin64 -d ',' testfile2.xlsx
Bob,Alice,Sue
Yes,No,Yes
No,,Yes
```

example) set delimiter to ',' and set excel tab to 0

```
acBook-Pro:testdata chroum$ ./xlsx2csv-go-darwin64 -d ',' -i 0 testfile2.xlsx
Bob,Alice,Sue
Yes,No,Yes
No,,Yes
```

### convert to csvfile

example) out to file : out.csv

```
MacBook-Pro:testdata chroum$ ./xlsx2csv-go-darwin64 -d ',' -i 0 -o out.csv testfile2.xlsx
2nd-ConoHa-MacBook-Pro:testdata chroum$ cat out.csv
Bob,Alice,Sue
Yes,No,Yes
No,,Yes
```
