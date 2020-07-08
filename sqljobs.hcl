job db  {
     sqls database1{
        sqltype = "slowquery"
        sql = <<SQL
            select * from users
        SQL
     }
     sqls database1{
        sqltype = "slowquery"
        sql = <<SQL
            select * from users
        SQL
     }
}

job http  {
     sqls database1{
        sqltype = "slowquery"
        sql = <<SQL
            select * from users
        SQL
     }
     sqls database1{
        sqltype = "slowquery"
        sql = <<SQL
            select * from users
        SQL
     }
}