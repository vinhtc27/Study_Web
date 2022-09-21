# Welcome Go Framework 

Hi! Đây là **GO Framework** mà team Backend sẽ build để phục vụ cho mục đích viết các service để làm các nhiệm vụ REST API, giao tiếp với Kafka, giao tiếp với Database. Ngoài ra sẽ hỗ trợ các dev xây dựng khung chương trình, cách để dễ tìm kiếm và xử lý lỗi khi cần. Dễ tiếp cận lúc mới học. 


# Structure

```
1. cmd 
	 2. main
		 3. main.go
 2. config
	 3. Keys
		 4. private.key
		 5. public.key 
	 4. development.yaml 
	 5. production.yaml 
 3. docs
 4. pkg
	 5. auth
		 6. basic.go
		 7. jwt.go
	 6. crypt
		 7. crypt.go
	 7. db
		 8. db.go
		 9. mongo.go
		 10. mysql.go 
		 11. postgres.go 
	 8. log 
		 9. log.go
	 9. router 
		 10. handler.go 
		 11. middleware.go 
		 12. response.go 
		 13. router.go 
	 10. server 
		 11. config.go 
		 12. server.go 
	 11. **service**
		 12. index
			 13. index.go
			 14. auth.go 
		 13. **new-service**
			 14.  **controller**
				 15. **new-service.go**
			 15. **model**
				 16. **new-service.go**
			 16. **router.go**
	 12. docker-compose.yml 
	 13. Makefile 
```

### Giải thích tổng quan 
Cấu trúc Service sẽ được chia làm 3 phần : 
Phần 1 : Phần Core hỗ trợ việc phát triển các service bao gồm các folder **config**, **pkg**, **docs**
Phần 2 : Phần liên quan đến việc chạy và setting Service như **cmd** dùng chứa phần hàm main, chạy service. **docker-compose.yml** dùng để cài đặt để giao tiếp với các **service** khác. 
Phần 3: Phần này là phần quan trọng nhất, các dev sẽ chủ yếu xử lý các nhiệm vụ loggic tại đây, ở folder **new-service** hiện nay mình để **service** để chạy kiểm tra service **users** ở mức cơ bản nhưng đầy đủ. 

## Thiết kế Service của mình 
Như hiện tại các bạn có thể thấy mình có để 2 folder và 1 file **router.go** trong thư mục Service. Ý nghĩa và cách thiết kế Service như thế nào : 

 1. **Router.go** là nơi mình sẽ thiết kế tất cả các đường dẫn của mình để cho bên ngoài có thể sử dụng, có thể public với tất cả, hoặc tuỳ từng quyền mới có thể truy cập vào API đó.  
```
    router.Router.Get(router.RouterBasePath+"/users", controller.GetUser) 
     // API method GET, path : /users , handler function : controller.GetUser 
     
     router.Router.Post(router.RouterBasePath+"/users", controller.AddUser) 
     // API method POST, path: /users, handler function : controller.AddUser 
	
	router.Router.With(auth.JWT).Get(router.RouterBasePath+"/users", users.GetUser)
	// Check request with auth.JWT , after handler request API GET 
```



2. **users** hoặc một cái tên khác ví dụ : **new-service** tuỳ vào nghiệp vụ bạn xử lý. \
Vì chúng ta sẽ xử lý dưới dạng microservice do đó. 1 Service sẽ chỉ có một Folder này, ở đây tôi đang xử lý **users** service. \
Bên trong nó chúng ta sẽ thiết kế dạng MVC truyền thống tuy nhiên sẽ không có folder **View**]
3. **new-service/model** : 	\
Ở đây chúng ta sẽ khai báo các struct đại diện cho các đối tượng chúng ta cần xử lý ví dụ : User, Asset, Price ... cho từng nghiệp vụ của mình 
4. **new-service/controller**: Ở đây chúng ta sẽ khai báo các hàm để xử lý tương ứng với các **path** mà chúng ta đã khai báo ở trong **router.go**. \
Ví dụ: Users thì sẽ gồm các hàm GetAUsers, DeleteUser... 

## Đọc hiểu cấu trúc và đóng góp nâng cao 
## config 
Chứa các thông tin config của dự án, các key, biến môi trường phục vụ mục đích phát triển và deploy. 

## pkg 
Đây là **Folder** quan trọng nhất của cả **Project** chứa tất cả các **core** và **tiện ích** để sử dụng cho phần **service** như : 
1. package db : Nơi xử lý các phần kết nối, tương tác với các **database** 
2. package log : Nơi thiết kế các method **tiện ích ** cho việc **log** dễ hiểu và đầy đủ hơn 
3. package server : Nơi mình setup **Server**
4. package router : Nơi mình setup **Router** để có thể tương tác với các **middleware** dễ dàng, Xây dựng sẵn các hàm cho việc **response** dễ dàng hơn. 
5. package auth: Nơi mình thiết kế các hàm **auth cơ bản** hỗ trợ cho việc xác thực với các request 
6. package crypt: Nơi thiết kế các hàm xử lý đến việc mã hoá, hiện tại mình đang để tương tác với **RSA** 

Sắp tới đây thì folder **pkg** này không đủ cho việc có thêm các tiện ích được build sẵn để sử dụng cho mục **service** như **cache, kafka, email, crypto ** mong mọi người đóng góp, nhận xét, nâng cấp **code-base** tốt hơn nữa. 

## Run with Docker 
1. docker compose up 
2. docker compose stop