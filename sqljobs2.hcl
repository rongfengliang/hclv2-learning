job db  {
     sqls database5{
        sqltype = "slowquery5"
        sql = <<SQL
            select * from users
        SQL
     }
     sqls database6{
        sqltype = "slowquery6"
        sql = <<SQL
            select * from users
        SQL
     }
}

job http  {
     sqls database7{
        sqltype = "slowquery7"
        sql = <<SQL
            select * from users
        SQL
     }
     sqls database8{
        sqltype = "slowquery8"
        sql = <<SQL
            select * from users
        SQL
     }
}