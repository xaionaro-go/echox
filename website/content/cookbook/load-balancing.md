+++
title = "Load Balancing"
description = "Load balancing multiple Echo servers using a reverse proxy server like Nginx, HAProxy and Armor."
[menu.main]
  name = "Load Balancing"
  parent = "cookbook"
+++

This recipe demonstrates how you can use Nginx, HAProxy or Armor as a reverse proxy server and load balance between multiple Echo servers.

## How to setup Nginx wth Echo?

### Step 1: Install Nginx

https://www.nginx.com/resources/wiki/start/topics/tutorials/install

### Step 2: Configure Nginx

Create a file `/etc/nginx/sites-enabled/localhost` with the following content:

```nginx
upstream localhost {
  server localhost:8081;
  server localhost:8082;
}

server {
  listen          8080;
  server_name     localhost;
  access_log      logs/localhost.access.log main;

  location / {
    proxy_pass      http://localhost;
  }
}
```
> Replace localhost with your domain e.g. api.labstack.com.

### Step 3: Start upstream servers

- `cd upstream`
- `go run server.go server1 :8081`
- `go run server.go server2 :8082` 

### Step 4: Start Nginx

`nginx`

### Step 5: Browse to https://localhost:8080

You should see a webpage being served from "server 1" or "server 2".

```sh
Hello from upstream server server1
```

## [Source Code]({{< source "load-balancing" >}})

`upstream/server.go`

{{< embed "load-balancing/upstream/server.go" >}}

`nginx.conf`

{{< embed "load-balancing/nginx.conf" >}}

## Maintainers

- [vishr](https://github.com/vishr)
