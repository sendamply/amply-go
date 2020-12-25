package amply

var amplyClient client = client{url: "https://sendamply.com/api/v1", accessToken: ""}

var Email email = email{client: &amplyClient}

func SetAccessToken(token string) {
	amplyClient.accessToken = token
}
