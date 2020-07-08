dbconf clickhousedb {
    driver = "clickhouse"
    dsn = "tcp://host1:9000?database=clicks&read_timeout=10&write_timeout=20"
    bind = ":8080"
}