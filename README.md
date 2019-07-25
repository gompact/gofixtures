# gofixtures

version: 3.0.0

[![Build Status](https://travis-ci.org/schehata/gofixtures.svg)](https://travis-ci.org/schehata/gofixtures)

A small command line tool written in Go lang, that loads fixtures
and insert it's records to a database. Currently supports YAML and JSON files
through CLI, and supports PostgreSQL only.

#### Supported File Types

- [X] JSON
- [X] YAML
- [X] CSV

#### Supported Databases

- [x] PostgreSQL
- [ ] MySQL
- [ ] SQLServer 
- [ ] MongoDB
- [ ] Redis
- [ ] Cassandra
- [ ] Firebase

#### Install

This will install gofixtures to your $GOPATH/bin

```bash
$ go get github.com/schehata/gofixtures/v3
```

#### Usage
    
To start using gofixtures you need to have a configuration file, that includes information
on which db driver should it use, and what are the credentials
It could be written
in YAML or JSON, here is an example for `.gofixtures.yaml` which will be automatically loaded
by gofixtures:

```go
import (
    "github.com/schehata/gofixutres/v3"
    "github.com/schehata/gofixutres/v3/entity"
)

func main() {
    config := entity.Config {
      entity.DBConfig{
        Driver: "postgres"
      }
    }
    gf, err := gofixutres.New(config)
    if err != nil {
      log.Fatal(err)
    }
    err = gf.LoadFromFiles(string[]{"./fixtures/countries.yml"})
    if err != nil {
      log.Fatal(err)
    }
}
```


example of `./fixtures/countries.yml`:

```yaml
table: countries
records:
  - name: Egypt
  - capital: Cairo
  - name: Germany
    capital: Berlin
  - name: Netherlands
    capital: Amsterdam
```

3. the previous yaml file inserts three records into table `countries`
4. gofixtures will parse each record and insert it into the database


#### Notes on CSV Support

- Column names are read from the first row
- Filename will be used as a tablename
- For now only comma ',' is allowed as a separator, will work on providing CLI flags to change that as needed

## Usage from Command Line 

You can use the gofixtures CLI tool to quickly load data into datastores from the command line, check the CLI docs.