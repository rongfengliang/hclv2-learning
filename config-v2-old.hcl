
job http demo1{
    source {
       dsn = "demo:demo@tcp(127.0.0.1:3306)/demo"
       driver = "mysql"
       query = "select * from demo1"
       schedule = "*/5 * * * * *"
    }
    transform {
       jsengine = "otto"
       message = <<JS
        if ( $rows.length < 1 ) {
            return
        }
        log("this is a demo")
        var msg =  "";
         _.chain($rows).pluck('name').each(function(name){
            msg += name+"--------demo--from otto----";
        })
         var info = {
            msgtype: "text",
            text: {
                content: msg
            }
        }
        log(JSON.stringify(info))
        send(JSON.stringify(info))
        JS
    }
    sink {
       webhook = "http://127.0.0.1:4195"
       driver = "resty",
       header = []
    }

}

job db demo2{
    source {
       dsn = "demo:demo@tcp(127.0.0.1:3306)/demo"
       driver = "mysql"
       query = "select * from demo1"
       schedule = "*/5 * * * * *"
    }
    transform {
       jsengine = "otto"
       message = "select * from demo1"
    }
    sink {
       dsn = "demo:demo@tcp(127.0.0.1:3306)/demo"
       driver = "mysql"
    }

}