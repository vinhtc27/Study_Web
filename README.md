# Chatella Backend 

## Các công nghệ sử dụng trong repository
+ Language: [golang](https://www.postgresql.org/docs/) 
+ Database: [postgresql](https://www.postgresql.org/docs/) 
+ Security: [jwt](https://jwt.io/)
+ Cache: [dgraph-io/ristretto](https://github.com/dgraph-io/ristretto)
+ Router: [go-chi](https://go-chi.io/#/)
+ Websocket: [gorilla/websocket](https://github.com/gorilla/websocket)
## Database
![Database](https://gcdnb.pbrd.co/images/7jF9iJD2Tp0z.png?o=1)
+ Id của các bảng đều được đánh index.
## Các bước build và chạy trên server
+ Do backend có kết nối với gmail để gửi email nên phải chắc chắn kết nối smtp không bị unreachable
+ Sử dụng lệnh ```ping smtp.gmail.com``` cho đến khi hết bị lỗi ```Network is unreachable``` và nhận được response
+ Bấm ```Ctrl + C``` để ngừng ping tới mail server của Google
+ Di chuyển đến thư mục chứa file main.go ```cd cmd/main/```
+ Build file execution cho linux ```env GOOS=linux GOARCH=amd64 go build main.go```
+ Di chuyển file ```main``` ra thư mục ngoài cùng của backend ```mv main ../..```
+ Di chuyển ra thư mục ngoài cùng của backend ```cd ../..```
+ Mở port 8080 ```expose 8080```
+ Chạy server ```./main```
+ Truy cập vào trang ```w42g11.int3306.freeddns.org``` để sử dụng trang web

# Các tài khoản đã có sẵn

## Tài khoản 1
+ email: ```trinhcongvinh27022002@gmail.com```
+ password: ```Vinh123@```

## Tài khoản 2
+ email: ```20021389@vnu.edu.vn```
+ password: ```Tholoc_2002```

## Tài khoản 3
+ email: ```mhoanganh25@gmail.com```
+ password: ```Hoanganh253```





