# Goenvs

## Description

Package simply loads all environment variable from .env file placed at the project root.

## Instalation

```
go get github.com/jackoski/goenvs
```

## Usage

Just at the top of your main function implement package, nad put yuor `.env` file at the project root

```golang
package main

import (
  env "github/jackoski/goenvs"
) 

func main(){
  env.Load()

  // rest of your code...
}
```

## example of .env file

```golang

# Comments are allowed in separate line
# Each KEY=value need to be separated by `=`

# valid
KEY=value

# invalid
KEYvalue

# inline comments are prohibited
Key=value # this will cause a bug

```

