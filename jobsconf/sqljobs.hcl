job db  {
     sqls database1{
        sqltype = "slowquery1"
        sql = <<SQL
            select * from users
        SQL
     }
     sqls database2{
        sqltype = "slowquery2"
        sql = <<SQL
            select * from users
        SQL
     }
}

job http  {
     sqls database3{
        sqltype = "slowquery3"
        sql = <<SQL
            select * from users
        SQL
     }
     sqls database4{
        sqltype = "slowquery4"
        sql = <<SQL
            select * from users
        SQL
     }
}