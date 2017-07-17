+++
title = "File Download Recipe"
description = "File download recipe for Echo."
[menu.main]
  name = "File Download"
  parent = "cookbook"
+++

## How to download a file?

### Server

`server.go`

{{< embed "file-download/server.go" >}}

### Client

`index.html`

{{< embed "file-download/index.html" >}}

## How to download a file as inline, opening it in the browser?

### Server

`server.go`

{{< embed "file-download/inline/server.go" >}}

### Client

`index.html`

{{< embed "file-download/inline/index.html" >}}

## How to download a file as attachment, prompting client to save the file?

### Server

`server.go`

{{< embed "file-download/attachment/server.go" >}}

### Client

`index.html`

{{< embed "file-download/attachment/index.html" >}}
