package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/backup"
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/config"
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/internal/domain"
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/internal/repository"
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/internal/repository/database"
	"gorm.io/gorm"
)

// backupDatabase ทำการ Backup Database และลบไฟล์เก่า
func backupDatabase(cfg *config.Config) error {
	timestamp := time.Now().Format("20060102_150405") // YYYYMMDD_HHMMSS
	backupDir := "backup"
	backupFile := fmt.Sprintf("%s/db_backup_%s.sql", backupDir, timestamp)

	// ตรวจสอบว่าโฟลเดอร์ backup มีอยู่หรือไม่ ถ้าไม่มีให้สร้าง
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		if err := os.Mkdir(backupDir, os.ModePerm); err != nil {
			return fmt.Errorf("❌ Failed to create backup directory: %v", err)
		}
	}

	var cmd *exec.Cmd
	if cfg.Database.Driver == "postgres" {
		// ใช้ pg_dump สำหรับ PostgreSQL
		cmd = exec.Command("pg_dump", "-h", cfg.Database.Host, "-p", fmt.Sprintf("%d", cfg.Database.Port),
			"-U", cfg.Database.User, "-d", cfg.Database.Name, "-F", "c", "-f", backupFile)
		cmd.Env = append(os.Environ(), "PGPASSWORD="+cfg.Database.Password) // ใช้สำหรับ PostgreSQL
	} else if cfg.Database.Driver == "mysql" {
		// ใช้ mysqldump สำหรับ MySQL
		cmd = exec.Command("mysqldump", "-h", cfg.Database.Host, "-P", fmt.Sprintf("%d", cfg.Database.Port),
			"-u", cfg.Database.User, "-p"+cfg.Database.Password, cfg.Database.Name, "--result-file="+backupFile)
	} else {
		return fmt.Errorf("❌ Unsupported database driver: %s", cfg.Database.Driver)
	}

	// รันคำสั่ง Backup
	log.Println("🚀 Backing up database before migration...")
	if err := backup.BackupDatabase(cfg); err != nil {
		log.Fatal("❌", err)
	}

	log.Println("✅ Database backup completed:", backupFile)

	// **เรียกฟังก์ชันลบไฟล์ Backup เก่ากว่า X วัน**
	retentionDays := cfg.Backup.RetentionDays
	if retentionDays > 0 {
		deleteOldBackups(backupDir, retentionDays)
	}

	return nil
}

// deleteOldBackups ลบไฟล์ backup ที่มีอายุมากกว่าจำนวนวันที่กำหนด
func deleteOldBackups(backupDir string, retentionDays int) {
	log.Printf("🧹 Cleaning up backups older than %d days...\n", retentionDays)

	// คำนวณ timestamp ของไฟล์ที่ควรลบ
	expiration := time.Now().AddDate(0, 0, -retentionDays)

	err := filepath.Walk(backupDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ข้าม directories
		if info.IsDir() {
			return nil
		}

		// ถ้าไฟล์เก่ากว่า retentionDays → ลบไฟล์
		if info.ModTime().Before(expiration) {
			log.Println("🗑 Deleting old backup:", path)
			if err := os.Remove(path); err != nil {
				log.Printf("❌ Failed to delete %s: %v\n", path, err)
			} else {
				log.Println("✅ Deleted:", path)
			}
		}

		return nil
	})

	if err != nil {
		log.Println("❌ Error cleaning backups:", err)
	} else {
		log.Println("✅ Backup cleanup completed!")
	}
}

// dropTables ลบ Table ทั้งหมด
func dropTables(db *gorm.DB) error {
	log.Println("⚠️ Dropping all tables...")
	err := db.Migrator().DropTable(&domain.Breed{}) // ลบ Table ทั้งหมด
	if err != nil {
		return err
	}
	log.Println("✅ All tables dropped successfully!")
	return nil
}

func main() {
	// โหลด Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("❌ Failed to load config:", err)
	}

	// เชื่อมต่อ Database
	dbInstance, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal("❌ Database setup failed:", err)
	}

	db, err := dbInstance.Connect(cfg)
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	// **1️⃣ Backup Database**
	if err := backupDatabase(cfg); err != nil {
		log.Fatal(err)
	}

	// **2️⃣ Drop Tables**
	if err := dropTables(db); err != nil {
		log.Fatal("❌ Failed to drop tables:", err)
	}

	// **3️⃣ ทำ Migration ใหม่**
	log.Println("🚀 Running database migrations...")
	if err := db.AutoMigrate(&domain.Breed{}); err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	log.Println("✅ Migration completed successfully!")

	// **4️⃣ Insert Master Data**
	log.Println("🚀 Seeding master data...")
	if err := repository.SeedBreeds(db); err != nil {
		log.Fatal("❌ Failed to seed master data:", err)
	}
	log.Println("✅ Master data inserted successfully!")
}
