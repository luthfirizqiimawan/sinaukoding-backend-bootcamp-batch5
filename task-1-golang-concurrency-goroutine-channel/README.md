# 🧩 Task 1 — Golang Concurrency (Goroutines + Channel + WaitGroup)

Tugas ini merupakan latihan dasar penggunaan **Concurrency di Golang** dengan memanfaatkan:

- Goroutines
- Channel
- WaitGroup

Program akan memproses sebuah slice angka secara paralel,  
mengalikan setiap angka dengan `2`, lalu mengumpulkan hasilnya kembali.

---

## ⚙️ How It Works

- Setiap angka diproses oleh `worker()` melalui goroutine
- Hasil dikirim melalui **channel**
- `WaitGroup` memastikan seluruh proses selesai
- Channel ditutup setelah semua worker selesai
- Hasil digabungkan dalam slice `results`

---

## 📌 Source Code

File utama:

```
main.go
```

Bahasa utama:

```
Golang
```

---

## 🚀 Run Program

Jalankan program dengan perintah:

```bash
go run main.go
```

---

## 📝 Example Output

```
Angka awal: [3 6 9 11 14]
Hasil proses: [6 12 18 22 28]
```

> Catatan: Urutan hasil dapat berbeda karena proses berjalan **secara concurrent**.

---

## 🎯 Learning Objectives

Dari tugas ini, dipelajari konsep:

- Dasar Goroutines
- Komunikasi antar goroutine menggunakan Channel
- Sinkronisasi proses dengan WaitGroup
- Dasar paralelisme pada Golang
