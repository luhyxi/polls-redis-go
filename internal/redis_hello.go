package internal

func get_client() (string, error) {
	url, err := get_redis_url()
	if err != nil {
		return "", err
	}

	return url, nil
}
