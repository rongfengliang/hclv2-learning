
job http demo1{
    jobinfo {
       webhook = "http://127.0.0.1:4195"
       driver = "resty"
       jsengine = "otto"
       location = env.LOC1
    }
}

job db mysql1 {
    jobinfo {
       webhook = "http://127.0.0.1:4195"
       driver = "mysql"
       jsengine = "otto"
       location = env.LOC2
    }
}


job http demo2{
    jobinfo {
       webhook = "http://127.0.0.1:4195"
       driver = "resty"
       jsengine = "otto"
       location = env.LOC1
    }
}

job db mysql2 {
    jobinfo {
       webhook = "http://127.0.0.1:4195"
       driver = "mysql"
       jsengine = "otto"
       location = env.LOC2
    }
}