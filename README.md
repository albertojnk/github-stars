# golang-crud-spa

[![Build Status](https://travis-ci.com/albertojnk/golang-crud-spa.svg?branch=master)](https://travis-ci.com/albertojnk/golang-crud-spa)

##

If you don't have Docker/Docker-Compose check **Setup Docker** section

<details>
<summary><b>Setup Docker</b></summary>
<p>

## Docker
macOS: <a href="https://docs.docker.com/docker-for-mac/install/"> https://docs.docker.com/docker-for-mac/install/ </a>

linux: <a href="https://docs.docker.com/install/linux/docker-ce/ubuntu/"> https://docs.docker.com/install/linux/docker-ce/ubuntu/ </a>




## Docker Compose

linux: <a href="https://docs.docker.com/compose/install/"> https://docs.docker.com/compose/install/ </a>
</p>
</details>

## Setup application & Run it

As we are using golang with go modules, it's recommended to clone this repository outside your $GOPATH, so:

* check your $GOPATH

```
    $ echo $GOPATH
```

* clone repository to any directory outside $GOPATH

```
    $ git clone https://github.com/albertojnk/golang-crud-spa.git
```

* cd to the repository:

```
    $ cd your-chosen-path/golang-crud-spa
```

* run docker-compose:

```
    $ docker-compose up
```
