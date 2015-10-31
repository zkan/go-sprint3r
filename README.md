# go-sprint3r
Go at Sprint3r

file: `*.go`

run: `go run <file.go>` ซึ่งจริงๆ คือมันไป compile เป็น binary แล้วรัน แล้วก็ลบทิ้ง

build: `go build <file.go>` เพื่อสร้าง binary ซึ่งสามารถเอาไปรันที่ไหนก็ได้ ตัวนี้จะเป็น static library ซึ่งทุก package จะรวมอยู่ในไฟล์ binary นี้อยู่แล้ว

package ชื่อ main จะเป็น entry point (เปลี่ยนเป็นชื่ออื่นไม่ได้)

เวลา import package เข้ามา เวลาจะเรียกใช้ func ข้างใน ให้ใส่นำหน้า func นั้นด้วย ชื่อ package เสมอ (จริงๆ มีวิธีเรียกโดยไม่ต้องใช้ชื่อ package นำหน้าก็ได้ โดยใช้ . นำหน้าชื่อ package ตอน import แต่วิธีนี้ไม่แนะนำ จะทำให้อ่านโค้ดลำบาก)

ตัวอักษรแรกของ func จะเป็นตัวอักษรใหญ่หมด (ยกเว้น main) จะทำให้ func นั้นสามารถใช้นอก package ได้ แต่ถ้า func เป็นชื่อตัวเล็กหมด จะหมายความว่า func นั้นสามารถใช้ได้แค่ใน package นั้นเท่านั้น (ชื่อตัวแปรก็เหมือนกัน)

เซต GOPATH

`export GOPATH=/Users/zkan/Projects/go-sprint3r/web`

หรือ

`export GOPATH=$HOME/Projects/go-sprint3r/web`

`go env` ดู environment variables

https://github.com/roofimon/Contact/blob/master/server.go
