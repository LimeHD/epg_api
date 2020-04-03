package helpers

func GetDbConnectionString(user string, pass string, host string, db string) string {
	return user + ":" + pass + host + "/" + db
}
