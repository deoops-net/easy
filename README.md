# easy

golang sugar library


## OS Methods

`import "github.com/deoops-net/easy"`

* RootDir() string
returns the absolute root directory

```golang
fmt.Println(RootDir())
```

## Hash Methods

`import "github.com/deoops-net/easy/hash"`

* MD5(string) string

```golang
fmt.Println(hash.MD5("foo"))
```

## Database

based on official mongodb driver, for using mongo easily

`import db "github.com/deoops-net/easy/db/mongo"`


* db.IniConn(db.MongoConf) nil
* db.Table(string) *mongo.Collection

```golang
conf := db.MongoConf{}
db.IniConn(conf)
db.Table("users").FindOne(context.Background(), bson.M{})
```
