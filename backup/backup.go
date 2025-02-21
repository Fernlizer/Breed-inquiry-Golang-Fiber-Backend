package backup

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/config"
)

// BackupDatabase ทำการ Backup Database ถ้าถูกเปิดใช้งานใน config
func BackupDatabase(cfg *config.Config) error {
	if !cfg.Backup.Enable {
		log.Println("⚠️ Database backup is disabled in config.")
		return nil
	}

	// ตรวจสอบและเลือก pg_dump ตาม OS
	pgDumpPath := cfg.Backup.PgDumpPath
	if pgDumpPath == "" {
		switch runtime.GOOS {
		case "windows":
			pgDumpPath = `"C:\Program Files\PostgreSQL\17\bin\pg_dump.exe"`
		case "linux", "darwin":
			pgDumpPath = "pg_dump" // ใช้จาก PATH
		default:
			return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
		}
	}

	// ตรวจสอบว่า pg_dump มีอยู่จริง
	if _, err := exec.LookPath(pgDumpPath); err != nil {
		return fmt.Errorf("pg_dump not found at: %s", pgDumpPath)
	}

	// กำหนดโฟลเดอร์ Backup
	backupDir := "backup"
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		_ = os.Mkdir(backupDir, 0755)
	}

	// ตั้งชื่อไฟล์ Backup
	timestamp := time.Now().Format("20060102_150405")
	backupFile := fmt.Sprintf("%s/%s_backup_%s.sql", backupDir, cfg.Database.Name, timestamp)

	// คำสั่ง `pg_dump`
	cmd := exec.Command(pgDumpPath,
		"-h", cfg.Database.Host,
		"-p", fmt.Sprintf("%d", cfg.Database.Port),
		"-U", cfg.Database.User,
		"-d", cfg.Database.Name,
		"-F", "c", "-f", backupFile,
	)

	// ตั้งค่าให้ `pg_dump` ใช้ Password จาก ENV
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", cfg.Database.Password))

	//รันคำสั่ง Backup
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("database backup failed: %v", err)
	}

	log.Println("✅ Database backup completed:", backupFile)
	return nil
}
