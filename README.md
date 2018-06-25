# Queuelib

This is basic rabbitmq library, which pushed data to directly on queue or to an exchange. It makes use of a singleton type connection, i.e once connection got established, same connection handler will be used for the lifetime of the app to publish messages to queue.

# How to use

* Import library as `import ( "queuelib" )`.
* Define config of the queue to be connected as below:
    ``` config := queuelib.Config{
		Username:     "guest",
		Password:     "guest",
		Host:         "localhost",
		Vhost:        "/",
		Queuename:    "test",
		Exchange:     "",
		ExchangeType: "fanout",
		Routingkey:   "test",
	} ```
    `if data is being pushed to exchange, no need to pass queuename`.

* Initialize to eastablish the connection:
    ``` initErr := queuelib.Init(&config)
	if initErr != nil {
		return initErr
	} ```

* get the connection handler as `queuelib.Conn`.
* Publish message to queue:
    ``` publishError := queue.Publish(config, data)
	if publishError != nil {
		return publishError
	} ```


