package main

func ProccessMsg(b []byte) []byte {
	return []byte(string(b) + " " + "sucessfully processed")
}
