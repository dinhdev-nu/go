## Playlist GO MEMBER

[Danh sách GO MEMBER](https://www.youtube.com/playlist?list=PLw0w5s5b9NK6qiL9Xzki-mGbq_V8dBQkY)

## go-backend-vetautet

Trên thực tế, ứng dụng bán hàng lượng đồng thời CAO về cơ bản là ứng dụng toàn diện các công nghệ có tính đồng thời cao trong các tình huống cụ thể . Bạn phải quen thuộc với các tình huống đồng thời cao và khả năng thiết kế kiến ​​trúc đồng thời cao là một năng lực quan trọng không thể thiếu đối với các lập trình viên hiện nay. 

Do đó, ngay cả khi doanh nghiệp của bạn chưa đụng tới hoặc không có kịch bản ứng dụng bán hàng lượng đồng thời CAO, chúng ta vẫn nên tìm hiểu về cấu trúc ứng dụng bán hàng lượng đồng thời CAO và càng sớm càng tốt.

Đây là dự án vetautet được phát triền bởi TEAM JAVA và nay chuyển qua GO. 
Dự án này có hai phần chính.

1 - Với chi phí tiết kiệm làm sao có thể tăng hiệu quả trong vấn đề READ. Nghĩa là từ 10K lên tới 30K req/sec. Thì nhiệm vụ của một BACKEND sẽ kết hợp với nhiều stack công nghệ như thế nào?
2 - Vấn đề tiếp theo là WRIRE chính là đặt vé tàu vào ngày tết or ngày BlackFriday thì hệ thống phải đảm bảo hiệu suất tốt nhất và nhất quán hàng tồn kho.
...

Các anh chị có thể đọc thêm tại đây: [Series: vetautet](https://github.com/anonystick/anonystick)

## Dành cho AI?

1 - Sinh viên và lập trình viên đã có một số kinh nghiệm và kiến ​​thức cơ bản về lập trình go và hiểu basic về backend
2- Những lập trình viên muốn đi phỏng vấn trong vòng 2 tuần thì có thể nắm vững kiến ​​trúc đồng thời CAO trong thời gian ngắn;
3 - Các Lập trình viên muốn nâng cao hiểu biết của mình về kiến ​​trúc của một hệ thống có lượng đồng thời CAO.

## How to run 

Trong folder có thư mục `environment` ở đây chứa tất cả các Stacks công nghệ bạn phải nằm thông qua như
`Kafka`, `redis`, `master/slave`, `sentinel`, `RateLimiter`, `CircuitBreaker`, `MySQL`...

Chú ý ở đây có hai file `docker-compose` nếu các bạn máy có cấu hình trung bình, vui lòng 

```bash
~ docker-compose -f environment/docker-compose-dev.yml up
```

Thì docker sẽ run những container cần thiết cho dự án.

Ngược lại nếu máy của bạn khác mạnh thì có thể run bản full của docker.

```bash
~ docker-compose -f environment/docker-compose-pro.yml up
```

Ngoài ra hãy chuyên nghiệp khi sử dụng `makefile`.

```bash
docker_up:
	docker-compose -f environment/docker-compose-dev.yml up
docker_down:
	docker-compose -f environment/docker-compose-dev.yml down
```

## How to test

```bash
~ curl http://localhost:8002/v1/2024/ticket/item/4
```

Output:

```go 
{
    "code":20001,
    "message":"success",
    "data":
        {
            "ticket_Id":1,
            "ticket_Name":
            "Vé Sự Kiện 12/12 - Hạng Phổ Thông",
            "stock_Available":1000,
            "stock_Initial":1000
        }
    }
```

