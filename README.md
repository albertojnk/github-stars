# github-stars

[![Build Status](https://travis-ci.com/albertojnk/github-stars.svg?branch=master)](https://travis-ci.com/albertojnk/github-stars)

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
    $ git clone https://github.com/albertojnk/github-stars.git
```

* cd to the repository:

```
    $ cd your-chosen-path/github-stars
```

### before you run compose

As docker compose creates/uses a `.data` folder in the root of the directory and because of that, sometimes when running compose it throws an permission error. To avoid this the best you can do is creating giving `.data` folder permissions:

```
    $ sudo mkdir .data
    $ sudo chmod 777 -R .data
```

If you are still getting permissions error try deleting the folder:

```
    $ sudo rm -rf .data
```

### Run 

* run docker-compose:

```
    $ docker-compose up
```

* go to your browser and type:

```
    http://localhost:8090
```

#

Have fun tagging your starred repositories :D