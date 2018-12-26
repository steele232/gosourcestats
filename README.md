# gosourcestats
Simple CLI Command to give statistics on the size of a Go repository.  

Written entirely in Go (GoLang) and [Simplified Mandarin using ZouYu](https://github.com/steele232/zouyu) (also maintained by the Author)

## 路线图 (Roadmap)
- List out Number of Go files
- Directory depth command flag/parameter
- Support 'all' (more) operating systems. Right now, it works on Linux/MacOS because I'm using '/' and not handling other more sophisticated methods of handling directory traversing. 

## 安装指示 (Installation Instructions)
- Should be ```go get``` -able ... then install it with ```go install```
- ```gosourcestats``` command works on the directory that you are in and any directories below it, but works only on ```.go``` files
