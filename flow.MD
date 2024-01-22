## Flow System

- Partner Send Transaction Request
- Our server consume those transaction request and send into kafka
- Kafka worker will process those transction

### Problems

#### How to ask the server if transaction is invalid?

- First request is for asking a available balance
- Second request is for create a new transaction

#### How we update maximum tenor each customer

- Create 1 table for configuration and then update all customer based on their profiling

## Consumer

- Data Keterangan
- NIK Nomor KTP Konsumen
- Full name Nama Lengkap Konsumen
- Legal name Nama Konsumen Di KTP
- Tempat Lahir Tempat lahir konsumen sesuai KTP
- Tanggal Lahir Tanggal lhir konsumen sesuai KTP
- Gaji Gaji Konsumen
- Foto KTP Foto KTP Konsumen
- Foto Selfie Foto Selfie Konsumen
- Loan Risk, generated based on Risk Configuration

## Partner Transaction

- Nomor Kontrak Nomor Kontrak untuk setiap transaksi konsumen
- OTR Angka On The Road transaksi barang baik itu White Godds, Motor atau Mobil konsumen
- Admin Fee Angka admin transaksi barang baik itu White Godds, Motor atau Mobil konsumen
- Jumlah Cicilan Angka jumlah cicilan transaksi barang baik itu White Godds, Motor atau Mobil konsumen
- Jumlah Bunga Angka bunga yang ditagihkan setiap transaksi barang baik itu White Godds, Motor atau Mobil konsumen
- Nama Asset Nama Asset yang dibeli konsumen

## LoanRisk

- Loan Risk

## LoanRiskLimit

- Risk Configuration
- Tenor Duration
- Limit
