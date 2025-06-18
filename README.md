# PAY SLIP
<div id="top"></div>
<!-- PROJECT LOGO -->
<br/>
<div align="center">
  <h3 align="center">Project-Task "Pay Slip" </h3>

</div>

<!-- ABOUT THE PROJECT -->
### 💻 &nbsp;About The Project

Pay Slip merupakan project task untuk sistem penggajian menggunakan tech-stack RESTful API dengan menggunakan bahasa Golang.    
Dilengkapi dengan berbagai fitur yang memungkinkan user employee dan user admin melakukan proses record kehadiran(attendance), lembur(overtime), reimbursement, sampai payroll, pay slip, dan summary dari salaries employee. 
Adapun fitur yang ada dalam RESTful API kami antara lain :
<div>
<details>
<summary>🙎 Users</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
 Fitur users digunakan untuk melihat detail list users dan juga login agar bisa mengakses fitur slip pay.
 
<div>
  
| Feature Users | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| GET | users  | YES | - | Melihat detail list user |
| POST | users/login | - | - | Melakukan login pada sistem slip pay |
</details>

<details>
<summary>🙎 Attendance</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
 Merupakan fitur agar employee bisa record kehadirannya mulai dari jam masuk(check-in) kantor dan jam pulang(check-out) yang hanya bisa dilakukan saat weekdays. Di sini juga ada attendance-period untuk admin agar bisa mengatur periode tanggal mulai dan selesainya periode kehadiran untuk penggajian tertentu.
 
<div>
  
| Feature Attendance | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | employee/login-time  | - | YES | Melakukan proses check-in kantor |
| POST | employee/logout-time | - | YES | Melakukan proses check-out kantor |
| POST | /admin/attendance-period | - | YES | Mengatur periode tanggal mulai dan tanggal selesai untuk penggajian |

</details>

<details>
<summary>⏱ &nbsp;Overtime</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
Ovetime merupakan fitur yang digunakan oleh employee untuk pengajuan lembur. Bisa dilakukan pengajuan lembur ketika employee sudah memenuhi jam masuk.

| Feature Overtime | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /employee/overtime | - | YES | Mengajukan pengajuan lembur |

</details>

<details>
<summary>📝 &nbsp;Reimbursement</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
Reimbursement merupakan fitur yang digunakan ketika sejumlah uang perlu direimburs. Di sini terdapat deskripsi agar employee menyampaikan hal-hal reimburs tersebut.

| Feature Reimbursement | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | employee/reimbursement | - | YES | Mengajukan jumlah uang yang perlu direimburs |

</details>

<details>
<summary>🧧 &nbsp;Payroll</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
Payroll nerupakan fitur khusus admin untuk melakukan payroll/pembayaran gaji sesuai attendance period.

| Feature Payroll | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /admin/payroll | - | YES | Melakukan payroll/pembayaran gaji |

</details>

<details>
<summary>📜 &nbsp;Payslip</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
Payslip merupakan fitur yang digunakan oleh employee untuk melihat take home pay sesuai dengan kehadiran(attendance), lembur(overtime), dan reimbursement.

| Feature Contractor | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /employee/payslip | - | YES | Mencetak/melihat slip gaji |

</details>

<details>
<summary>📓 &nbsp;Summary Salaries</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
Summary Salaries merupakan fitur yang digunakan oleh admin untuk melihat take home pay masing-masing employee serta take home pay total dari seluruh employee.

| Feature Contractor | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| GET | /admin/summary-salaries | - | YES | Mencetak/melihat take home pay employee |

</details>

### Swagger
<a href="https://app.swaggerhub.com/apis-docs/faizalsundara/Test-BuddyKu/1.0.0/"><strong> Link »</strong></a>

<!-- IMAGES -->
<!-- ### 🖼&nbsp;Images -->

<!-- CONTACT -->
### Contact

[![GitHub Faizal](https://img.shields.io/badge/-Faizal-white?style=flat&logo=github&logoColor=black)](https://github.com/faizalsundara)

<p align="center">:copyright: 2025 | Faizal </p>
</h3>