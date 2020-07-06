
job http demo{
    webhook = "http://127.0.0.1:4195"
    driver = "mysql"
    dsn = "demo:demo@tcp(127.0.0.1:3306)/demo"
    jsengine = "otto"
    myinfo = env.USERNAME2
    query = <<SQL
        SELECT users.* FROM users
    SQL
    schedule = "* * * * * *"
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

job db mysql1 {
    webhook = "http://127.0.0.1:4195"
    driver = "mysql"
    dsn = "demo:demo@tcp(127.0.0.1:3306)/demo"
    jsengine = "otto"
    myinfo = env.USERNAME2
    query = <<SQL
        SELECT users.* FROM users
    SQL
    schedule = "* * * * * *"
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