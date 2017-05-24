## Golang : custom cookie

paket dibutuhkan : 

	import "github.com/gorilla/securecookie"
	import "net/http"

penggunaan : 

jadikan 1 direktori dengan program go lainya agar bisa digunakan

program sederhana untuk menetapkan nama cookie dan juga valuenya
cara penggunaan sangat sederhana, terdapat 3 fungsi dalam program
ini yaitu

* menetapkan cookie


	set_cookie(name_cookie string,value_cookie string,res http.ResponseWriters)

contoh : 
	
	set_cookie("cookie_1","akses",res)


* memanggil cookie


	get_cookie(name_cookie string,req *http.Request) string

contoh : 

	get_cookie("cookie_1",req) string


* menghapus cookie


	clear_cookie(name_cookie string,res http.ResponseWriters)

contoh : 

	clear_cookie("cookie_1",res)


semoga bermanfaat