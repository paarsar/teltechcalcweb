Setup
-------------------------

#### Navigate to project directory in your command line

### Grab Dependencies
```
make grab-dependencies
```

### run project
```
make run
```

### Navigate
```
http://localhost/add?x=1&y=1
```


By default the server runs on port 80 if you want to change the port number alter the `httpport` property in `conf/app.conf` 

```
httpport=<desired port>
````