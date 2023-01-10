# mdf
Markdown for file links

![image](https://user-images.githubusercontent.com/30147836/211445570-994d5266-ccea-48a5-95d9-513d581dbd80.png)

I couldn't find an easy way to do it in a shell script, so I wrote this tool. Hey `awk`/`sed`/`xargs` guru, challenge you to a script in comments! :smile:

## Usage

```
./mdf --help
usage: mdf [<flags>] [<path>]

Flags:
  --help                     Show context-sensitive help (also try --help-long and --help-man).
  --link-prefix=LINK-PREFIX  Link prefix
  --column=COLUMN ...        List of columns

Args:
  [<path>]  Path with files to list
```

## Examples

```
./mdf ~/<some-repo>/<sub-dir> --link-prefix=https://github.com/xornivore/<some-repo>/tree/main/<sub-dir> --column File --column Review | pbcopy
```


