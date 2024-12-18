# CINEMA SEATING

## Description

````
You are required to complete one take-home assignment, which is a Coding Task
designed to help us evaluate your coding style. Use this opportunity to showcase your
understanding of coding best practices.
Please note that the assignment is confidential. Do not share or publish the questions,
your code, design, or any related information with the public or anyone not involved in
your hiring process with VulcanLabs

Coding & Implementation Task

Due to the ongoing pandemic, seating arrangements must follow specific distancing
rules:
• People within the same group can sit together without any restrictions.
• Different groups must maintain a minimum distance from one another. This
distance is measured using the Manhattan distance (diagonal moves are not
allowed).
For example, in the seating layout below, the shortest distance between the two 1’s is
marked with x’s, and the distance is 7. If min_distance = 6, this seating arrangement is
acceptable.

[[1, X, X, X, X],
[0, 0, 0, 0, X],
[0, 0, 0, 0, X],
[0, 0, 0, 0, 1]]

We are looking for a solution to ensure cinema patrons follow social distancing rules.
Your task is to develop a gRPC or Restful service that we can use for its cinema
operations.
Requirements:
• The service should be configurable to support different cinema sizes and
minimum distance rules, specified as: (rows, columns, min_distance).

• The service must provide a RPC method or API for querying a set of available
seats that can be purchased together.
• Another RPC method or API should allow the reservation of specific seats,
identified by their (row, column) coordinates ((0, 0) being the top left).
• Another RPC method or API should allow the cancellation of specific seats,
identified by their (row, column) coordinates, with (0, 0) representing the top left
corner.
Important Notes:
• Clearly state any assumptions you make regarding business requirements or
preferences.
• Ensure that your solution highlights your coding style and technical knowledge.
````

#### Mục tiêu của bài toán
1. Xây dựng API quản lý chỗ ngồi trong rạp chiếu phim:
- Quy định khoảng cách tối thiểu giữa các nhóm khách hàng (Manhattan distance).
- Cho phép khách hàng:
  - Tìm kiếm ghế có thể mua chung mà vẫn đảm bảo giãn cách.
  - Đặt trước các vị trí ghế cụ thể.
  - Hủy đặt chỗ với các vị trí đã chọn.

2. Yêu cầu kỹ thuật 
- Xây dựng gRPC hoặc RESTful API.
- Xử lý logic khoảng cách Manhattan.
- Dữ liệu rạp
- Đảm bảo tính linh hoạt

#### Triển khai
1. Chọn công nghệ
- Golang: Ưu tiên vì hiệu năng và khả năng triển khai nhanh.
- Framework: Gin
- Triển khai theo chuẩn Restful API
  - `POST /cinema/configure`: Cấu hình kích thước rạp và khoảng cách
  - `GET /cinema/available-seats`: Truy vấn các cụm ghế khả dụng và không khả dụng
  - `POST /cinema/reserve`: Đặt ghế cụ thể.
  - `POST /cinema/cancel`: Hủy đặt ghế.
- Database: Lưu trữ tạm file JSON (mục tiêu nhanh gọn trong quá trình làm test, không cần setup phức tạp), phần này có thể thay thế bằng database khác như PostgreSQL, MySQL, MongoDB. Để lựa chọn được database phù hợp cần đánh giá nghiệp vụ của sản phẩm, hiệu nặng, bảo mật, chi phí ...

2. Tính khoảng cách
* Manhattan distance: Tính khoảng cách giữa 2 điểm. Với p1 tại (x1, y1) và p2 tại (x2, y2) thì sẽ thực hiện tính khoảng cách theo công thức bên dưới
```text
distance = |x1 - x2| + |y1 - y2|
```

3. Document API
- Lựa chọn Swagger vì thay thế cho việc viết document bằng tay (Word, Confluent ...), giúp tạo ra document tự động từ code vì nhanh hoàn thiện.

4. Hạn chế
- Chưa xử lý trường hợp nhiều rạp phim.
- Mỗi lần cập nhật lại cấu hình rạp phim sẽ reset lại toàn bộ dữ liệu

### Requirement
- Golang go1.23.3 or highest

### Setup

1. Clone project:
```bash
$ git clone https://github.com/thienkb1123/cinema-seating.git
```

1. Install dependencies:
```bash
$ go mod download
```

2. Run project:
```bash
$ make run or go run cmd/app
```

### How to Test

1. Run test:
```bash
$ make test
```

### How to Use Swagger

1. Swag init if not have folder `docs`:
```bash
$ make swag-v1
```

2. Open browser and access to `http://localhost:8080/swagger/index.html` for testing API.
