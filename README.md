![screenshot](https://user-images.githubusercontent.com/16364286/105629391-4bc02400-5e75-11eb-8b27-e2b31957f1f1.png)

# Table of Contents
1. [Introduction](#introduction)
2. [Features](#features)
3. [How to Use](#how-to-use)

# Introduction
Cacing is quite simple in memory cache engine.

# Features
* Intuitive API
* Single executable for server and client
* Pluggable storage engine
* Single server instance for single application client

# How to Use
Starting the server by run:
```bash
$ cacing run --user root --password 123 --port 8081
```

Connect to server (as client) by run:
```bash
$ cacing connect --dsn cacing://root:123@localhost:8081
```
