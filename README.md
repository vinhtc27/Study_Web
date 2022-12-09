# Chattela Backend 

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
