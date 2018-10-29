# gofixtures

version: 2.1.0

[![Build Status](https://travis-ci.org/ishehata/gofixtures.svg)](https://travis-ci.org/ishehata/gofixtures)

A small command line tool written in Go lang, that loads fixtures
and insert it's records to a database. Currently supports YAML and JSON files
through CLI, and supports PostgreSQL only.

#### Supported File Types

- [X] JSON
- [X] YAML
- [ ] CSV

#### Supported Databases

- [X] PostgreSQL
- [ ] MySQL
- [ ] SQLServer 
- [ ] MongoDB
- [ ] Redis
- [ ] Cassandra
- [ ] Firebase

#### Install

This will install gofixtures to your $GOPATH/bin

```bash
$ go get github.com/emostafa/gofixtures
$ go install github.com/emostafa/gofixtures
```

#### Usage

To start using gofixtures you need to have a configuration file for the database connection, It could be written
in YAML or JSON, here is an example for `db.yaml`

```yaml
driver: "postgres"
database: "mydb"
user: "foo"
password: "bar"
host: "localhost"
```

The next step is to prepare your fixtures:

1. Prepare YAML file (or json file), let's call it example.yaml
2. Start writing your fixture, it should have a "table" which declares that table name
that we are going to insert data into, and then a list of records, as following:

```yaml
table: countries
records:
  - name: Egypt
  - name: Germany
  - name: Netherlands
```

3. the previous yaml file inserts three records into table `countries`
4. gofixtures will parse each record and insert it into the database
5. in order to use gofixutres, change directory to one level above where the yaml file exists, run command

```bash
$ gofixtures --file fixtures/example.yaml --dbconf db.yaml
```


6. by default, gofixtures expecte yaml files to exists in "fixtures/" directory, but you can override this by either:
	a. specify a file to load, e.g: `$ gofixtures -file myfixture.json`
	b. specify a directory and loal all the fixtures files inside it, e.g: `$ gofixtures -directory /home/foo/myfixtures`

A combination of all the available flags can be used, e.g:

```bash
$ gofixtures --dbconf mydbconf.yaml --dir ./my_fixtures 
```

##### Avialable Command Line Flags
1. dbconf "database configuration YAML (or JSON) file"
3. file "a yaml or json file to load"
4. dir "a directory contains fixtures"



#### TODO
1. ~~support JSON files~~
2. support different sql databases like MySQL
3. ~~load multiple yaml fiels, or load folders~~
4. ~~ability to load configuration from file instead of kwargs~~
