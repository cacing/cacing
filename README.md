![screenshot](https://user-images.githubusercontent.com/16364286/105848033-046e9a80-6011-11eb-9b32-80f6e8ce1838.gif)

# Table of Contents
1. [Introduction](#introduction)
2. [Features](#features)
3. [How to Use](#how-to-use)

# Introduction
Cacing is a simple in memory cache engine.

# Features
* Intuitive API
* Key-value store
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

## Available Command
* **SET** a key with value.

  example:
  ```
  SET user1 cacing
  SET user2 hadihammurabi
  SET user3 needkopi
  ```
* **GET** value of a key

  example:
  ```
  GET user1
  GET user2
  GET user3
  ```
* **DEL** to delete a key

  example:
  ```
  DEL user1
  DEL user2
  DEL user3
  ```
* **EXISTS** to check key existence

  example:
  ```
  EXISTS user1
  EXISTS user2
  EXISTS user3
  ```
