# golang-crud-spa


If you don't have Docker/Docker-Compose check **Setup** section below


<details>
<summary><b>Setup</b></summary>
<p>

## Install Docker

###### Uninstall old versions

```
$ sudo apt-get remove docker docker-engine docker.io containerd runc
```

###### Set up the repository

```
$ sudo apt-get update

$ sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

$ sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
```

###### Install Docker Engine - Community

```
$ sudo apt-get update

$ sudo apt-get install docker-ce docker-ce-cli containerd.io

$ sudo apt-get install docker-ce=<VERSION_STRING> docker-ce-cli=<VERSION_STRING> containerd.io

$ sudo docker run hello-world
```

## Install Docker Compose 

```
$ sudo curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

$ sudo chmod +x /usr/local/bin/docker-compose

$ sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```

</p>
</details>

## Running the application

Run the application is very simple:

* check your $GOPATH

```
$ echo $GOPATH
```

* clone this repository to any directory outside $GOPATH

```
$ git clone https://github.com/albertojnk/golang-crud-spa.git
```
 or

```
$ git clone git@github.com:albertojnk/golang-crud-spa.git
```

* cd to the repository:

```
$ cd golang-crud-spa
```

* run docker-compose:

```
$ docker-compose up
```
