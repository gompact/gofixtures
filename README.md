# gofixtures

a small command line tool written in Go lang, that loads a yaml file
and insert it's records to postgresql database.

#### Install

this will install gofixtures  to your $GOPATH/bin

```
go get github.com/eslammostafa/gofixtures
```

#### Usage

1. Prepare YAML file, let's call it example.yaml
2. Start writing records, each record must start with "table_name" field, as following

```
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

```
gofixtures exmaple.yaml database=your_db_name user=db_user_name
```

6. gofixtures depend on keyword arguments for the database connection, available arguemnts

```
database: defaults to "postgres"
user: defaults to "postgres"
password: defaults to ""
port: defaults to 5432
host: defaults to localhost
```

7. the gofixture command will load the yaml file, parse records and insert them into database,
   and it will print the insertion queries


#### Supported Field Types

only the following field types are supported so far

```
string
int
Time
```
more field types to be added


#### TODO
1. support JSON files
2. support different sql databases like MySQL
3. support nosql databases
4. support more field types like Date and DateTime, float, JSON
5. load multiple yaml fiels, or load folders
6. ability to load configuration from file instead of kwargs
