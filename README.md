# Add.go
Given a JSON dictionary containing `x` and `y`, outputs the sum of x and y.

Extra keys in the dictionary will be ignored.  Missing keys will default to zero.

## Usage
```
add [-v] <input file path>
```
The program will also read directly from standard input, e.g.:
```
echo '{"x": 5.5, "y": -1.542}' | add
```
```
add < somefile.json
```

#### Verbose flag
If the `-v` or `--verbose` flag is provided on the command line, the entire expression will be output rather than just the result, for example:

```
> echo '{"x": 5.5, "y": -1.542}' | add
3.958

> echo '{"x": 5.5, "y": -1.542}' | add -v
5.5 + -1.542 = 3.958
```

