# recap coding challenge

## Getting started

For the CLI to work you require at least the environment variables `API_KEY`
containing your API key for the NASDAQ End of Day US Stock Prices API.

If a `.env` file is present, the CLI will automatically read its values into
environment variables.

```
go run main.go -symbol=AAPL -start=2017-10-01 -end=2017-10-31
```

### Email

To send results via email you can use the boolean flag `-email`.

```
go run main.go -symbol=AAPL -start=2017-10-01 -end=2017-10-31 -email
```

It uses [AWS Simple Email Service](https://aws.amazon.com/ses/) and requires the
following environment variables:

```
EMAIL_SENDER=
EMAIL_RECEIVER=
AWS_ACCESS_KEY=
AWS_SECRET_KEY=
```

Note that the access key needs to belong to an AWS user with an
`AmazonSESFullAccess` IAM policy attached.
