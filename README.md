# PWL_TUGAS2_122140017
# Saga Microservices - Order, Payment, Shipping System

## Deskripsi Tugas

Sistem ini dikembangkan menggunakan pola **Saga** dengan 3 service utama yang terintegrasi, yaitu **Order Service**, **Payment Service**, dan **Shipping Service**. Setiap service memiliki tugas tertentu dalam urutan yang telah ditentukan. Jika salah satu layanan gagal, kompensasi dilakukan untuk mengembalikan sistem ke keadaan semula.

### 1. **Order Service**:
Order Service bertugas untuk membuat order berdasarkan permintaan dari klien. Ketika order berhasil dibuat, statusnya akan berubah menjadi **PAID**, dan Order ID yang dihasilkan akan digunakan untuk melanjutkan ke langkah berikutnya.

### 2. **Payment Service**:
Payment Service menerima request dari Order Service untuk memproses pembayaran. Jika pembayaran berhasil, Payment Service akan mengkonfirmasi pembayaran yang sukses, dan proses akan diteruskan ke Shipping Service untuk pengiriman barang. Jika pembayaran gagal, Order Service akan dipanggil untuk membatalkan order yang sudah dibuat, dan sistem akan mengirimkan respons kegagalan.

### 3. **Shipping Service**:
Shipping Service menerima request dari Payment Service untuk memproses pengiriman barang. Jika pengiriman berhasil dilakukan, maka order dianggap selesai dan sistem akan memberikan respons yang menyatakan bahwa pesanan telah diproses dengan sukses. Jika pengiriman gagal, kompensasi dilakukan secara berurutan dimulai dari Shipping Service akan dibatalkan terlebih dahulu, kemudian Payment Service akan dipanggil untuk mengembalikan dana, dan akhirnya Order Service akan membatalkan order yang telah dibuat.

### **Proses Kompensasi**:
Proses kompensasi dilakukan secara berurutan jika terjadi kegagalan dalam pengiriman barang. Jika pengiriman gagal, maka pengembalian dana dilakukan melalui Payment Service, dan Order Service akan dibatalkan. Hal ini memastikan bahwa transaksi yang gagal dapat dibatalkan sepenuhnya tanpa meninggalkan data yang tidak konsisten di dalam sistem.

## Alur Pengujian (Testing)

### 1. **Memastikan Semua Service Dapat Berjalan di Port yang Sesuai**:
- **Order Service** berjalan di `http://localhost:8000`
- **Payment Service** berjalan di `http://localhost:8001`
- **Shipping Service** berjalan di `http://localhost:8002`
- **Orchestrator Service** berjalan di `http://localhost:8003`

### 2. **Membuat Order Baru**:
Mennggunakan software **Postman** untuk mensimulasikan alur proses. Untuk membuat order baru, lakukan request ke **Order Service** dengan atribut `id`, `jumlah`, dan `status produk`. Setelah request berhasil, Order Service akan memberikan response dengan status yang sesuai.

### 3. **Proses Pembayaran**:
Setelah order berhasil dibuat, melakukan pemanggilan ke **Payment Service** untuk memproses pembayaran. Jika pembayaran berhasil, maka status order akan berubah menjadi "success".

### 4. **Proses Pengiriman (Shipping)**:
Setelah pembayaran berhasil, lakukan pemanggilan ke **Shipping Service**. Pada proses ini, hanya informasi `id` dan `nama produk` yang ditampilkan karena jumlah produk sudah diproses sebelumnya di Payment Service.

### 5. **Proses Kompensasi (Jika Shipping Gagal)**:
Jika terjadi kesalahan dalam **Shipping Service** dan pengiriman gagal, maka sistem akan otomatis melakukan kompensasi. Proses kompensasi dilakukan sebagai berikut:
  1. **Cancel Order**: Order Service akan dibatalkan terlebih dahulu.
  2. **Refund Payment**: Payment Service akan melakukan refund untuk mengembalikan dana.
  3. **Cancel Order Again**: Order Service akan membatalkan pesanan sebagai langkah akhir dari proses kompensasi.

### 6. **Langkah-langkah Kompensasi**:
  - **Shipping Service Cancelled**: Jika pengiriman gagal, Shipping Service akan dibatalkan.
  - **Refund Payment**: Payment Service akan menghapus jumlah produk yang telah dibayar dan mengembalikan dana.
  - **Cancel Order**: Order Service akan membatalkan pesanan yang telah dibuat.


Menjalankan service masing-masing pada terminal yang terpisah menggunakan perintah:
```bash
go run main.go
