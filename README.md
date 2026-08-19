# Go Backend Assessment — Giáp

Bài test này thay cho buổi phỏng vấn trực tiếp. Gồm 2 phần:

1. **Code** — sửa các bug trong một Go service nhỏ (không cần viết mới từ đầu).
2. **QUESTIONS.md** — trả lời một số câu hỏi ngắn.

## Setup

1. Repo này được push lên GitHub (private). Em được thêm làm collaborator, hoặc fork về tài khoản của em.
2. Tạo branch mới, ví dụ `giap/assessment`.
3. Sửa code trực tiếp trên branch đó, commit, rồi mở Pull Request về `main`.

## Chạy thử

```
go build ./...
go test ./... -race -v
```

`go build` sẽ pass ngay từ đầu — code hiện tại compile được, chỉ là logic có bug.
`go test` sẽ **fail hoặc không compile** ở một vài package — đó chính là các bug cần sửa.
Mỗi bug đều có comment `TODO(candidate)` ngay tại vị trí cần sửa.

## Danh sách bug cần sửa

| File | Vấn đề |
|---|---|
| `internal/model/user.go` | `UpdateName` dùng value receiver nên không có tác dụng lên struct gốc |
| `internal/db/db.go` | `GetDB()` có race condition khi nhiều goroutine gọi cùng lúc |
| `internal/repository/user_repository.go` | `GetUser` không tôn trọng `context` (cancel/timeout) |
| `internal/service/user_service.go` | panic khi không tìm thấy user, thay vì trả về `error` |
| `internal/handler/user_handler.go` | không validate `id`, không set timeout, không map lỗi ra đúng HTTP status |

Có thể cần sửa nhiều file cùng lúc vì signature thay đổi (ví dụ sửa `service.GetUser` để trả thêm `error` thì `handler` cũng phải sửa theo) — đó là chủ ý, không phải lỗi của bộ đề.

Không cần thêm feature, không cần viết thêm test — test đã có sẵn, mục tiêu là làm chúng pass mà vẫn đúng bản chất (không hack test).

## Sau khi sửa code

Trả lời các câu hỏi trong `QUESTIONS.md`, viết thẳng vào file đó.

## Nộp bài

Mở Pull Request. Repo có GitHub Actions tự chạy `go build` + `go vet` + `go test -race` trên mỗi PR — nhìn cột check ở PR để biết đã pass hết chưa trước khi báo lại.
