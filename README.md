# book-api-gorm

Repo ini adalah sebuah api serderhana menggunakan GO (gin web framework) dan PostgreSQL. Project ini mirip dengan <a href="https://github.com/togihon/book-api"> Project sebelumnya</a> tetapi
terdapat beberapa perbedaan, diantaranya:  
1. Tidak membuat table di postgre secara manual, melainkan memanfaatkan migrasi menggunakan `gorm`
2. terdapat perubahan pada struct book dimana `description` dihilangkan dan penambahan `createdAt` dan `updatedAt`  
  
  Untuk menjalankan ketik `go run main.go` diterminal
# ubah config database

Jangan lupa untuk menyesuaikan config database dengan mengubah data di bagian  
pkg/database/database.go  

# body request postman untuk create dan update
```
 {  
  "title": "insert title",  
  "author": "insert author"   
 }  
``` 

 # test postman
Berikut ini adalah HTTP method beserta route-nya  
  
GET `localhost:8080/books/` &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//get all books  
GET `localhost:8080/books/:bookID` &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;  //get book by id  
POST `localhost:8080/books/` &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; //create book  
PUT `localhost:8080/books/:bookID` &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;   //update book by id  
DELETE `localhost:8080/books/:bookID` &nbsp;&nbsp;  //delete book by id  
