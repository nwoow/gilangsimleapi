# Golang Divide Assignment

After installing packages simply run this:

	go run main.go


# Run the code using Docker:

If you wish to tun this code simple run this command to build image:

	docker-compose build


If you wish to run the code simple run:

	docker-compose up

To access the endpoint use the URL:

	0.0.0.0:8080/api/divide

And the payload should be:

	{
	"a":3,
	"b":9
	}

Run the code in CURL use:

	curl --location --request POST 'localhost:8080/api/divide' \
	--header 'Content-Type: application/json' \
	--data-raw '{
	"a":1,
	"b":1
	}'
