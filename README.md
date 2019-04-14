# [awacha](https://awacha.com/)

simple note web application

## Description

You can send and receive text from command line.

## Usage

```bash
$ curl  -X POST -d "hello world" awacha.com/helloworld
-> Created

$ curl awacha.com/helloworld
-> hello world
```

You can POST to any URL after slash of `awacha.com/`.

If POST to same URL, the saved text will be overwrite.

## Licence

[MIT](https://github.com/iPolyomino/awacha/blob/master/LICENSE)

## Author

[Hagi](https://github.com/iPolyomino)

