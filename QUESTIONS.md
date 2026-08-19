# Câu hỏi (40 điểm)

Trả lời ngắn gọn, viết thẳng dưới mỗi câu. Không cần dài, quan trọng là hiểu đúng bản chất — có thể tham chiếu đến phần code em vừa sửa.

## 1. Struct & composition (5đ)

Trong Go, `struct` là gì? Tại sao Go dùng struct + composition + interface thay vì class/inheritance như Java/C#?

_Trả lời:_ struct is a composition data type in go. it can has many fields of different types. Go use struct composition interface to avoid problems with hierachies like diamond problem, deep hierachies that is hard to modify

## 2. Value vs Pointer receiver (5đ)

Giải thích tại sao `UpdateName` ở bản gốc (trước khi em sửa) không có tác dụng. Ngoài lý do "muốn thay đổi object", còn lý do nào khác để dùng pointer receiver?

_Trả lời:_  UpdateName doesn't work before fixing is because a value receiver pass a copy of the variable into the function their for any changes made doesn't affect the original. Aside from mutation, memory is also a reason to use pointer receiver, if the copied object has a large size, we will quickly run out of memory.

## 3. `defer` trong thực tế (5đ)

Cho 3 ví dụ thực tế mà backend Go thường dùng `defer`.

_Trả lời:_ 3 use case of defer: cancel() to clean context, unlock mutex, recover from panic

## 4. Panic vs Error (5đ)

Tại sao không nên dùng `panic()` để xử lý lỗi business thông thường (ví dụ "user not found")? Panic nên dùng cho trường hợp nào?

_Trả lời:_ business logic error are expected in the flow and the flow will not stop but if panic is used then program will crash. We should only use panic for unrecoverable programmer errors that might stop the program from continue running.

## 5. `sync.Once` (5đ)

`sync.Once` dùng để làm gì? Cho 1 use case thực tế khác ngoài việc khởi tạo DB connection.

_Trả lời:_ sync once is used for a one time run. Other use case might be running one time initialization of some sort of metrics or handler at start of program

## 6. Singleton & connection pool (5đ)

Nếu nói `*sql.DB` là "singleton" trong application, điều đó có nghĩa là app chỉ mở đúng một TCP connection tới MySQL không? Giải thích.

_Trả lời:_ no, *sql.DB is a connection pool which means there are many connections, the only thing that should only exist once is the instance itself, creating many instance will create many pools and exhaust the db.

## 7. Authentication vs Authorization vs Validation (5đ)

Giải thích sự khác nhau giữa 3 khái niệm này, cho ví dụ mỗi loại.

_Trả lời:_ authentication is to check who you are, authorization is to check if what you can do, validation is to check if things are correct

## 8. Authorization scenario (5đ)

User A đã login thành công, gọi `GET /users/123`, nhưng `123` là data của User B. Request hợp lệ về format. Đây là lỗi authentication, authorization, hay validation? Vì sao?

_Trả lời:_  authorization because authorization is to check if you can do something or not and a user is not supposed to get data of another one
