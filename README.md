# gofixtures
version: 1.0.0

A small command line tool written in Go lang, that loads yaml files
and insert it's records to a database.

#### Install

This will install gofixtures to your $GOPATH/bin

```bash
$ go get github.com/objectizer/gofixtures
$ go install github.com/objectizer/gofixtures
```

#### Usage

1. Prepare YAML file, let's call it example.yaml
2. Start writing records, each record must start with "table_name" field, as following

```yaml
- table_name: countries
  name: Egypt
- table_name: countries
  name: Germany
- table_name: cities
  name: Cairo
  country_id: 1
- table_name: cities
  name: Munich
  country_id: 2
- table_name: cities
  name: Alexandria
  country_id: 1 
```

3. the previous yaml file creates two countries, and tree cities
4. gofixtures will parse each record and insert it to postgresql
5. in order to use gofixutres, change directory to where the yaml file exists, run command

```bash
$ gofixtures -file example.yaml -dbconf "user=eslam dbname=mydb sslmode=disable"
```

6. gofixtures expects the database configuration file in "db/dbconf.yml", but you can override this by either:
	a. supply db conf file (yaml file) in the command line `$ gofixtures -dbconffile mydbconf.yml`
	b. supply connection string in the command line `$ gofixtures -driver postgres -dbconf "user=eslam dbname=mydb sslmode=disable"


7. gofixtures expecte yaml files to exists in "fixtures/" directory, but you can override this by either:
	a. specify a yaml file, e.g: `$ gofixtures -file myfile.yaml`
	b. specify a directory and loal all the yaml files inside it, e.g: `$ gofixtures -directory /home/eslam/myfixtures`

A combination of all the available flags can be used, e.g:

```bash
$ gofixtures -dbconffile mydbconf.yaml -directory /home/eslam/fixtures 
```

##### Avialable Command Line Flags
1. dbconf "database connection string"
2. dbconffile "database configuration yaml file"
3. driver "defaults to postgres"
4. file "a yaml file to load"
5. directory "a directory contains fixtures"


##### Loading Database Configuration From YAML file

In order to load database configuration from a YAML file, gofixtures expects the files to contain two records, driver and open, e.g:
```yaml
driver: postgres
open: "user=eslam dbname=mydb sslmode=disable"
```


#### Supported Field Types

only the following field types are supported so far

```go
string
int
Time
```
more field types to be added


#### TODO
1. support JSON files
2. support different sql databases like MySQL
3. support more field types like Date and DateTime, float, JSON
4. ~~load multiple yaml fiels, or load folders~~
5. ~~ability to load configuration from file instead of kwargs~~
