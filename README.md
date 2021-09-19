# Loghighlight

For terminals that support ANSI color codes.
Pipe a log file through this tool to highlight words that haven't appeared in the last n lines.

## Description

Some log files contain repeating lines that can make it hard to notice other lines in between. This tool will leave repetitive words alone and highlight words as they appear for the first time.

The input are tokenized and tokens ("words") with only unicode letters may be highlighted, not numbers or special characters. The number of lines to consider when highlighting can be specified as a parameter. Stdin and stdout are used.

## Example

Install
```
go install github.com/larschri/loghighlight
```

Run
```
loghighlight < /var/log/syslog | less -R
```
