package database

/*
Program need to init database 1st before it can be used.

dbConfig := new(database.Config)
dbConfig.DSN["TxDB"] = "txdbdsn"
err := database.Init(dbConfig)
if err != nil {
	log.Fatal("Failed to connect to database")
}
*/

/*
After database is connected, database connection can be used using Get(databaseType) function

conn, err := database.Get(database.TxDB)
if err != nil {
	return err
}

Error need to be checked everytime database connection object retrieved.
I assume this approach is more make sense rather than creating a global variable to access the database.append
By using global variable, we cannot control the db connection object and don't know wether database is connected or not.

or you can let the database lib to check the error for you, but it will give you fatal once database connection object
is not exists

conn := database.GetFatal(database.TxDB)
*/

/*
As this wrapper is using sqlt, we can watch the database connection using its DoHeartBeat function.
HeartBeat function is to watch all databases connection and StopBeat to stop them all.

New connection need to be added to HeartBeat and StopBeat in order to watch the database connection and stop it
gracefully.
*/
