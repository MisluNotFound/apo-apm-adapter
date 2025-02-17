package skywalking

func getServer(componmentName string) string {
	if server, ok := nameToServerMap[componmentName]; ok {
		return server
	}
	return componmentName
}

var nameToServerMap = map[string]string{
	"mongodb-driver":                        "MongoDB",
	"rocketMQ-producer":                     "RocketMQ",
	"rocketMQ-consumer":                     "RocketMQ",
	"kafka-producer":                        "Kafka",
	"kafka-consumer":                        "Kafka",
	"activemq-producer":                     "ActiveMQ",
	"activemq-consumer":                     "ActiveMQ",
	"rabbitmq-producer":                     "RabbitMQ",
	"rabbitmq-consumer":                     "RabbitMQ",
	"postgresql-jdbc-driver":                "PostgreSQL",
	"Xmemcached":                            "Memcached",
	"Spymemcached":                          "Memcached",
	"h2-jdbc-driver":                        "H2",
	"mysql-connector-java":                  "Mysql",
	"Jedis":                                 "Redis",
	"FreeRedis":                             "Redis",
	"StackExchange.Redis":                   "Redis",
	"Redisson":                              "Redis",
	"Lettuce":                               "Redis",
	"Zookeeper":                             "Zookeeper",
	"SqlClient":                             "SqlServer",
	"Npgsql":                                "PostgreSQL",
	"MySqlConnector":                        "Mysql",
	"EntityFrameworkCore.InMemory":          "InMemoryDatabase",
	"EntityFrameworkCore.SqlServer":         "SqlServer",
	"EntityFrameworkCore.Sqlite":            "SQLite",
	"Pomelo.EntityFrameworkCore.MySql":      "Mysql",
	"Npgsql.EntityFrameworkCore.PostgreSQL": "PostgreSQL",
	"transport-client":                      "Elasticsearch",
	"SolrJ":                                 "Solr",
	"cassandra-java-driver":                 "Cassandra",
	"pulsar-producer":                       "Pulsar",
	"pulsar-consumer":                       "Pulsar",
	"rest-high-level-client":                "Elasticsearch",
	"mariadb-jdbc":                          "Mariadb",
	"Mysqli":                                "Mysql",
	"influxdb-java":                         "InfluxDB",
	"Predis":                                "Redis",
	"PyMysql":                               "Mysql",
	"spring-kafka-consumer":                 "kafka-consumer",
	"mssql-jdbc-driver":                     "SqlServer",
	"Psychopg":                              "PostgreSQL",
	"GoMysql":                               "Mysql",
	"ClickHouse-jdbc-driver":                "ClickHouse",
	"Apache-Kylin-jdbc-driver":              "Apache-Kylin",
	"MysqlClient":                           "Mysql",
	"AsyncPG":                               "PostgreSQL",
	"AIORedis":                              "Redis",
	"Impala-jdbc-driver":                    "Impala",
	"eventmesh-producer":                    "EventMesh",
	"eventmesh-consumer":                    "EventMesh",
	"amqp-producer":                         "amqp",
	"amqp-consumer":                         "amqp",
	"GoRedis":                               "Redis",
}
