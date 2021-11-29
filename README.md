# aScan
A high-performance concurrent scanner written by go, which can be used for survival detection, tcp port detection, and web service detection.

## Function

1. Survival detection of arbitrary permissions
2. TCP port detection
3. Web asset detection

Convenient asset sorting, all results are output after sorting !

## Usage

```
# View instructions
./aScan -h

  -t string
    	Scanned IP address(192.168.1.1 || 192.168.1.1-255 || 192.168.1.1/24), 192.168.1.1/12 is not allowed! (default "127.0.0.1")
  -r int
    	The number of open coroutines for Port detection (default 600)
  -wr int
    	The number of open coroutines for Web detection (default 100)
  -p string
    	Which ports range to scan(-p 22,23,8080-8081)
  -m string
    	Which ports dict to scan(lite, top100, top1000, all, custom) (default "top100")
  -ds int
    	Disable Survival Detection(-ds 1)
  -dw int
    	Disable Web Scanning(-dw 1)
  -f string
    	Scan the ip address in the file(-f ips.txt)
    	
 Notice:
      It is best not to adjust the number of coroutines, the default number balances stability and speed
```

## Example

> Scan a single ip

```
./aScan -t 192.168.1.1
```

> Scan the ip within the specified range

```
./aScan -t 192.168.1.1-10
```

> Scan IP in c class

```
./aScan -t 192.168.1.1/24
```

> Scan ip and specify the port to be scanned

````
./aScan -t 192.168.1.1 -p 22,80,8000-8080
````

> Set up a dictionary for scanning ports

````
./aScan -t 192.168.1.1 -m lite
````

> Set the number of scanning coroutines

````````````
./aScan -t 192.168.1.1 -r 500
````````````

> Set the number of coroutines to scan web pages

```
./aScan -t 192.168.1.1 -wr 50
```

> Scan the ip in the specified file

```
./aScan -f ips.txt
```

> Disable Survival Detection

```
./aScan -t 192.168.1.1 -ds 1
```

> Disable Web Scanning

```
./aScan -t 192.168.1.1 -dw 1
```

## Example

![image-1](https://raw.githubusercontent.com/seventeenman/aScan/main/img/image-1.jpg)

![image-2](https://raw.githubusercontent.com/seventeenman/aScan/main/img/image-2.png)

## Todo List

- [ ] Get port banner
- [ ] More feature-rich web detection
- [ ] Code performance optimization
- [ ] Possibility to obtain intranet host information
- [ ] Code structure adjustment to facilitate adding functions

## Issues

The scanner may have some unpredictable bugs in the web detection module written in the fasthttp library. If you encounter a bug, you can post the front-end source code of the target that caused the problem in the issue.
