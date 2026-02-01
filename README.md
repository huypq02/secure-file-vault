# S3 Secure File Storage & Sharing Mini-Project

## 🎯 Project Overview

This project provides a highly secure, scalable solution for file storage and sharing, designed to be easily deployable for personal use or as a mini SaaS offering. It leverages AWS S3 for encrypted file storage, PostgreSQL for metadata management, and follows clean architecture principles to ensure maintainability and rapid extensibility.

---

## 🔧 Features

### 1. 📤 Secure File Upload

- Users can upload files (PDFs, images, documents, etc.).
- Files are encrypted using AES before being stored in AWS S3.
- Metadata (file name, owner, creation date, expiry, etc.) is stored in PostgreSQL.

### 2. 🔐 End-to-End Encryption

- Server-side encryption guarantees that files are protected before storage.
- AES keys can be further encrypted using RSA/public key cryptography for advanced sharing scenarios.
- Only authorized recipients can decrypt and access file contents.

### 3. 📥 Secure File Download

- Users can retrieve and download previously uploaded files.
- Files are fetched from S3 and decrypted on the server before delivery.

### 4. 📎 File Sharing via Expiring Links

- Generate shareable links with configurable expiration times.
- Support for access limitation (number of downloads) and optional authentication requirements.

### 5. 📅 Automatic File Expiry & Cleanup

- Each file can have an expiry time set.
- Background job (cron or worker) automatically deletes expired files from S3 and removes metadata from the database.
- Ensures no stale or lingering data remains.

### 6. 🛡️ Authentication & Authorization _(Extensible)_

- Token-based authentication (JWT) for user validation.
- Each file is tied to a specific user (owner).
- Sharing can be restricted to specific users for enhanced access control.

---

## 🏗️ Architecture Principles

- **Clean Architecture:** Decoupled layers (Domain, Usecase, Interface, Infrastructure) enable easy testing, maintenance, and scaling.
- **Security-First:** Data is encrypted at rest and in transit; keys are protected using best practices.
- **Scalability:** Designed for both single-user deployment and SaaS multi-tenancy.
- **Maintainability:** Simple, modular codebase with clear separation of concerns.

---

## 🚀 Getting Started

1. **Clone the repository**
2. **Configure AWS, PostgreSQL, and environment variables**
3. **Run the server and background job scheduler**
4. **Access the API for file operations**

_See the [docs/SETUP.md](docs/SETUP.md) for detailed instructions._

---

## 🛡️ Security Considerations

- **AES Encryption** is performed server-side before files are uploaded to S3.
- **Key Management:** AES keys may be encrypted with RSA/public keys for secure sharing.
- **Access Control:** All API endpoints require authentication; sharing links are time- and usage-limited.
- **Data Lifecycle:** Expired files and metadata are purged automatically by scheduled jobs.

---

## 💡 Extensibility Questions

- How will key management evolve as you add more sharing and multi-user features?
- What strategies will you use to efficiently handle large file uploads/downloads?
- How will you ensure atomic deletion of files and metadata under high concurrency?
- How will you adapt authentication and authorization as the system scales to multi-tenant SaaS?

---

## 📚 References

- [Clean Architecture by Robert C. Martin](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)
- [AWS S3 Security Best Practices](https://docs.aws.amazon.com/AmazonS3/latest/userguide/security-best-practices.html)
- [Go Crypto Libraries](https://pkg.go.dev/golang.org/x/crypto)
- [PostgreSQL Official Documentation](https://www.postgresql.org/docs/)

---

## 🏆 Contributing

Pull requests, feature suggestions, and security reviews are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.
