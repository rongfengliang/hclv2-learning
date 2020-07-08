job http demo1{
    jobinfo {
       webhook = "http://127.0.0.1:4195"
       driver = "resty"
       query = "select * from demo1"
       location = env.LOC1
       schedule = "*/5 * * * * *"
    }
}

job db mysql1 {
    jobinfo {
       driver = "mysql"
       query = "select * from mysql1"
       location = env.LOC2
       schedule = "*/10 * * * * *"
    }
}



job http demo2{
    jobinfo {
       webhook = "http://127.0.0.1:4195"
       driver = "resty"
       location = env.LOC1
       query = "select * from demo2"
       schedule = "*/15 * * * * *"
    }
}

job db mysql2 {
    jobinfo {
       driver = "mysql"
       location = env.LOC2
       query = "select * from mysql2"
       schedule = "*/20 * * * * *"
    }
}


job all mysql3 {
    jobinfo {
       webhook = "http://127.0.0.1:4195"
       driver = "mysql"
       location = env.LOC2
       query = "select * from mysql2"
       schedule = "*/20 * * * * *"
    }
}