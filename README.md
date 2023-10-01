# API Documentation

## User

### Register User

Mendaftarkan pengguna baru.

- **URL**

  `/api/v1/users/register`

- **Method**

  `POST`

- **Request Body**
  
  - Format: JSON
  - Contoh:
    ```json
    {
      "username": "john_doe",
      "email": "john.doe@example.com",
      "password": "password123"
    }
    ```

- **Response**

  - Jika sukses:

    - Status: `200 OK`
    - Body:
      ```json
      {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjAsImV4cCI6MTY5NjI1MDczMX0.MDbmuWxkuFM4F1gDVSCM1AwB-_t-pIB5emy0EKcp7rY",
        "user": {
          "ID": 1,
          "username": "john_doe",
          "email": "john.doe@example.com",
          "created_at": "2023-10-01T19:45:31.673517+07:00",
          "updated_at": "2023-10-01T19:45:31.673517+07:00"
        }
      }
      ```

  - Jika terdapat kesalahan validasi (misalnya, username atau email sudah terdaftar):

    - Status: `400 Bad Request`
    - Body:
      ```json
      {
        "error": "Username sudah terdaftar"
      }
      ```

### Login User

Login pengguna.

- **URL**

  `/api/v1/users/login`

- **Method**

  `POST`

- **Request Body**
  
  - Format: JSON
  - Contoh:
    ```json
    {
      "email": "john.doe@example.com",
      "password": "password123"
    }
    ```

- **Response**

  - Jika sukses:

    - Status: `200 OK`
    - Body:
      ```json
      {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjAsImV4cCI6MTY5NjI1MDczMX0.MDbmuWxkuFM4F1gDVSCM1AwB-_t-pIB5emy0EKcp7rY",
        "user": {
          "ID": 1,
          "username": "john_doe",
          "email": "john.doe@example.com",
          "created_at": "2023-10-01T19:45:31.673517+07:00",
          "updated_at": "2023-10-01T19:45:31.673517+07:00"
        }
      }
      ```

  - Jika terdapat kesalahan validasi (misalnya, email atau password tidak valid):

    - Status: `400 Bad Request`
    - Body:
      ```json
      {
        "error": "Email atau kata sandi salah"
      }
      ```

### Update User

Memperbarui informasi pengguna.

- **URL**

  `/api/v1/users/:userId`

- **Method**

  `PUT`

- **Request Header**

  - Authorization: Bearer {Token}

- **Request Parameters**

  - userId: ID pengguna

- **Request Body**
  
  - Format: JSON
  - Contoh:
    ```json
    {
      "username": "john_doe_updated",
      "email": "john.doe.updated@example.com",
      "password": "newpassword123"
    }
    ```

- **Response**

  - Jika sukses:

    - Status: `200 OK`
    - Body:
      ```json
      {
        "message": "Informasi pengguna berhasil diperbarui"
      }
      ```

  - Jika terdapat kesalahan validasi (misalnya, username atau email sudah terdaftar):

    - Status: `400 Bad Request`
    - Body:
      ```json
      {
        "error": "Username sudah terdaftar"
      }
      ```

### Delete User

Menghapus akun pengguna.

- **URL**

  `/api/v1/users/:userId`

- **Method**

  `DELETE`

- **Request Header**

  - Authorization: Bearer {Token}

- **Request Parameters**

  - userId: ID pengguna

- **Response**

  - Jika sukses:

    - Status: `200 OK`
    - Body:
      ```json
      {
        "message": "Akun pengguna berhasil dihapus"
      }
      ```

  - Jika terdapat kesalahan validasi (misalnya, pengguna tidak ditemukan

 atau tidak diizinkan):

    - Status: `404 Not Found` atau `401 Unauthorized`
    - Body:
      ```json
      {
        "error": "Pengguna tidak ditemukan"
      }
      ```

## Photos

### Upload Photo

Mengunggah foto baru.

- **URL**

  `/api/v1/photos`

- **Method**

  `POST`

- **Request Header**

  - Authorization: Bearer {Token}

- **Request Body**
  
  - Format: JSON
  - Contoh:
    ```json
    {
      "title": "Sunset",
      "url": "https://example.com/sunset.jpg"
    }
    ```

- **Response**

  - Jika sukses:

    - Status: `200 OK`
    - Body:
      ```json
      {
        "message": "Foto berhasil diunggah",
        "photo": {
          "ID": 1,
          "title": "Sunset",
          "url": "https://example.com/sunset.jpg",
          "created_at": "2023-10-01T19:45:31.673517+07:00",
          "updated_at": "2023-10-01T19:45:31.673517+07:00"
        }
      }
      ```

### Get Photos

Mendapatkan daftar foto.

- **URL**

  `/api/v1/photos`

- **Method**

  `GET`

- **Response**

  - Jika sukses:

    - Status: `200 OK`
    - Body:
      ```json
      {
        "photos": [
          {
            "ID": 1,
            "title": "Sunset",
            "url": "https://example.com/sunset.jpg",
            "created_at": "2023-10-01T19:45:31.673517+07:00",
            "updated_at": "2023-10-01T19:45:31.673517+07:00"
          },
          // ...
        ]
      }
      ```

### Update Photo

Memperbarui informasi foto.

- **URL**

  `/api/v1/photos/:photoId`

- **Method**

  `PUT`

- **Request Header**

  - Authorization: Bearer {Token}

- **Request Parameters**

  - photoId: ID foto

- **Request Body**
  
  - Format: JSON
  - Contoh:
    ```json
    {
      "title": "Updated Sunset",
      "url": "https://example.com/updated-sunset.jpg"
    }
    ```

- **Response**

  - Jika sukses:

    - Status: `200 OK`
    - Body:
      ```json
      {
        "message": "Informasi foto berhasil diperbarui"
      }
      ```

### Delete Photo

Menghapus foto.

- **URL**

  `/api/v1/photos/:photoId`

- **Method**

  `DELETE`

- **Request Header**

  - Authorization: Bearer {Token}

- **Request Parameters**

  - photoId: ID foto

- **Response**

  - Jika sukses:

    - Status: `200 OK`
    - Body:
      ```json
      {
        "message": "Foto berhasil dihapus"
      }
      ```

  - Jika terdapat kesalahan validasi (misalnya, foto tidak ditemukan atau tidak diizinkan):

    - Status: `404 Not Found` atau `401 Unauthorized`
    - Body:
      ```json
      {
        "error": "Foto tidak ditemukan"
      }
      ```
