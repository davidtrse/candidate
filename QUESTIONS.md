# Câu hỏi (40 điểm)

Trả lời ngắn gọn, viết thẳng dưới mỗi câu. Không cần dài, quan trọng là hiểu đúng bản chất — có thể tham chiếu đến phần code em vừa sửa.

## 1. Struct & composition (5đ)

Trong Go, `struct` là gì? Tại sao Go dùng struct + composition + interface thay vì class/inheritance như Java/C#?

_Trả lời:_

## 2. Value vs Pointer receiver (5đ)

Giải thích tại sao `UpdateName` ở bản gốc (trước khi em sửa) không có tác dụng. Ngoài lý do "muốn thay đổi object", còn lý do nào khác để dùng pointer receiver?

_Trả lời:_

## 3. `defer` trong thực tế (5đ)

Cho 3 ví dụ thực tế mà backend Go thường dùng `defer`.

_Trả lời:_

## 4. Panic vs Error (5đ)

Tại sao không nên dùng `panic()` để xử lý lỗi business thông thường (ví dụ "user not found")? Panic nên dùng cho trường hợp nào?

_Trả lời:_

## 5. `sync.Once` (5đ)

`sync.Once` dùng để làm gì? Cho 1 use case thực tế khác ngoài việc khởi tạo DB connection.

_Trả lời:_

## 6. Singleton & connection pool (5đ)

Nếu nói `*sql.DB` là "singleton" trong application, điều đó có nghĩa là app chỉ mở đúng một TCP connection tới MySQL không? Giải thích.

_Trả lời:_

## 7. Authentication vs Authorization vs Validation (5đ)

Giải thích sự khác nhau giữa 3 khái niệm này, cho ví dụ mỗi loại.

_Trả lời:_

## 8. Authorization scenario (5đ)

User A đã login thành công, gọi `GET /users/123`, nhưng `123` là data của User B. Request hợp lệ về format. Đây là lỗi authentication, authorization, hay validation? Vì sao?

_Trả lời:_
