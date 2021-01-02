# [awacha](https://awacha.com/)

"awacha" is a simple note web application.

You can upload and download text data on the web from the terminal.

## Usage

You can POST to any URL after slash of `awacha.com/`.

If POST to same URL, the saved text will be overwrite.

### Example of use

- You want to copy the configuration file to my new computer
- You have multiple computers and want to share a (long) URL
- You want to share text data temporarily

#### Note

**Do not upload sensitive documents.** Awacha is a note application that does not require authentication. In other words, anyone can update, view, or delete the note you created.

### Example

#### Create note

```bash
$ curl -X POST -d "hello world" awacha.com/helloworld
> Created
```

You can also send a text file.

```bash
# Create a file to send in advance. (e.g., .vimrc)
$ curl -X POST --data-binary @$HOME/.vimrc awacha.com/vimrc
> Created
```

#### Get note

```bash
$ curl awacha.com/helloworld
> hello world
```

You can also save the file using the wget command.

```bash
$ wget awacha.com/vimrc -o .vimrc
$ cat .vimrc
> syntax on
> set fenc=utf-8
> set nobackup
> set noswapfile
```

#### Delete note

```bash
$ curl -X DELETE awacha.com/helloworld
> Deleted
```

## Licence

[MIT](https://github.com/iPolyomino/awacha/blob/master/LICENSE)

## Copyright

© 2019-2021 [Hagi](https://github.com/iPolyomino)
