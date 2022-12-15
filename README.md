# Objectives

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of your last project, [ascii-art](https://01.kood.tech/git/root/public/src/branch/master/subjects/ascii-art).

Your webpage must allow the use of the different banners:

+ [shadow](https://01.kood.tech/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
+ [standard](https://01.kood.tech/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
+ [tinkertoy](https://01.kood.tech/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)

Implement the following HTTP endpoints:

1. GET `/`: Sends HTML response, the main page.
   1.1 GET Tip: go templates to receive and display data from the server.
v
2. POST `/ascii-art`: that sends data to Go server (text and a banner)
   2.1 POST Tip: use form and other types of tags to make the post request.\

The way you display the result from the POST is up to you. What we recommend are one of the following :

+ Display the result in the route `/ascii-art` after the POST is completed. So going from the home page to another page.
+ Or display the result of the POST in the home page. This way appending the results in the home page.

The main page must have:

+ text input
+ radio buttons, select object or anything else to switch between banners
+ button, which sends a POST request to '/ascii-art' and outputs the result on the page.

## HTTP status code

Your endpoints must return appropriate HTTP status codes.

+ OK (200), if everything went without errors.
+ Not Found, if nothing is found, for example templates or banners.
+ Bad Request, for incorrect requests.
+ Internal Server Error, for unhandled errors.

In the root project directory create a README.MD file with the following sections and contents:

+ Description
+ Authors
+ Usage: how to run
+ Implementation details: algorithm

## Instructions

+ HTTP server must be written in **Go**.
+ HTML templates must be in the project root directory *templates*.
+ The code must respect the [good practices](https://01.kood.tech/git/root/public/src/branch/master/subjects/good-practices/README.md).

## Allowed packages

+ Only the [standard go](https://golang.org/pkg/) packages are allowed

## Usage

[Here's an example.](http://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20)

# Description
# Authors
# Usage: how to run
# Implementation details: algorithm

### Functional for audit

Has the requirement for the allowed packages been respected? (Reminder for this project: (only standard packages)

Does the project contain HTML files?

Try inputting with the standard template/banner the following example:

In the first line `{123}`

In the second line `<Hello> (World)!`
```

   __                     __
  / /  _   ____    _____  \ \
 | |  / | |___ \  |___ /   | |
/ /   | |   __) |   |_ \    \ \
\ \   | |  / __/   ___) |   / /
 | |  |_| |_____| |____/   | |
  \_\                     /_/

   __  _    _          _   _          __            __ __          __                 _       _  __    _
  / / | |  | |        | | | |         \ \          / / \ \        / /                | |     | | \ \  | |
 / /  | |__| |   ___  | | | |   ___    \ \        | |   \ \  /\  / /    ___    _ __  | |   __| |  | | | |
< <   |  __  |  / _ \ | | | |  / _ \    > >       | |    \ \/  \/ /    / _ \  | '__| | |  / _` |  | | | |
 \ \  | |  | | |  __/ | | | | | (_) |  / /        | |     \  /\  /    | (_) | | |    | | | (_| |  | | |_|
  \_\ |_|  |_|  \___| |_| |_|  \___/  /_/         | |      \/  \/      \___/  |_|    |_|  \__,_|  | | (_)
                                                   \_\                                           /_/
```

Does it display the right result as above?

Try to input `"123??"` using the template/banner standard.
```

                     ___    ___
 _   ____    _____  |__ \  |__ \
/ | |___ \  |___ /     ) |    ) |
| |   __) |   |_ \    / /    / /
| |  / __/   ___) |  |_|    |_|
|_| |_____| |____/   (_)    (_)

```
Does it display the right result as above?

Try to input `"$% "="` using the template/banner shadow.
```

                        _|  _|
  _|   _|_|    _|       _|  _|
_|_|_| _|_|  _|                _|_|_|_|_|
_|_|       _|
  _|_|   _|  _|_|              _|_|_|_|_|
_|_|_| _|    _|_|
  _|
```
Does it display the right result as above?

Try to input `"123 T/fs#R"` using the template/banner thinkertoy.
```

  0    --  o-o        o-O-o     o  o-o      | |  o--o
 /|   o  o    |         |      /   |       -O-O- |   |
o |     /   oo          |     o   -O-  o-o  | |  O-Oo
  |    /      |         |    /     |    \  -O-O- |  \
o-o-o o--o o-o          o   o      o   o-o  | |  o   o

```
Does it display the right result as above?

Does it display an understandable graphical representation of the result?

Try to navigate between all the available pages in the website.

Are all the pages working? Does the project implement [404 status](https://www.restapitutorial.com/httpstatuscodes.html)?

Does the project handle [HTTP status 400 - Bad Request](https://kinsta.com/knowledgebase/400-bad-request/#causes)?

Does the project handle [HTTP status 500 - Internal Server Errors](https://www.restapitutorial.com/httpstatuscodes.html)?

Try making a request to the server (clicking a button to generate the ascii-art representation on the website)

Is the communication between [server and client](https://www.geeksforgeeks.org/client-server-model/) well established?

Does the server uses the right [HTTP method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods)?

Did the site work without crashing at any time?

Is the server written in Go?

As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)

### General

+Does the server present all the needed [handlers and patterns](https://golang.org/pkg/net/http/#HandleFunc) for the http requests?

Basic

+Does the server run quickly and effectively? (Favoring recursive, no unnecessary data requests, etc)

+Does the code obey the [good practices](https://github.com/01-edu/public/blob/master/subjects/good-practices/README.md)?

+Is there a test file for this code?

+Are the tests checking each possible case?

+Are the instructions in the website clear?

+Does the project run using an API?

### Social

+Did you learn anything from this project?

+Can it be open-sourced / be used for other sources?

+Would you recommend/nominate this program as an example for the rest of the school?