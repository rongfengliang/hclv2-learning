job db mysql2 {
    jobinfo {
       webhook = "http://127.0.0.1:4195"
       driver = "mysql"
       jsengine = "otto"
       location = env.LOC2
       query = "select * from mysql2"
       schedule = "*/20 * * * * *"
    }
}

job http http1 {
    jobinfo {
       webhook = "http://127.0.0.1:4195"
       driver = "mysql"
       jsengine = "otto"
       location = env.LOC2
       query = "select * from mysql2"
       schedule = "*/20 * * * * *"
    }
}