package tugas3

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Deklarasi model Item
type Item struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:128"`
	Status      string `gorm:"size:64"`
	Amount      int
	ItemDetails []ItemDetail `gorm:"foreignKey:ItemID"`
}

// Deklarasi model ItemDetail
type ItemDetail struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:128"`
	ItemID uint   `gorm:"index"`
	Item   Item
}

func ConnectDatabase() (Db *gorm.DB, err error) {
	// Implementasi Gorm logger, mencetak ke stdout jika terjadi kesalahan atau query lambat
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		})

	// Buat koneksi ke database postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=127.0.0.1 user=hsisandbox3 password=hsisandbox3 dbname=hsisandbox3 port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		PreferSimpleProtocol: true,
	}),
		&gorm.Config{
			Logger: newLogger,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})

	// Kembalikan koneksi dan error
	return db, err
}

func DBMigrate(db *gorm.DB) {
	// Migrasikan tabel Item dan ItemDetail
	db.AutoMigrate(&Item{}, &ItemDetail{})
}

// (C) Create
func CreateItem(db *gorm.DB, item *Item) (err error) {
	// Buat item baru
	return db.Create(item).Error
}

func CreateItemDetail(db *gorm.DB, itemDetail *ItemDetail) (err error) {
	// Buat item detail baru
	return db.Create(itemDetail).Error
}

// (R) Read
func GetItems(db *gorm.DB) (items []Item, err error) {
	// Baca semua item
	return items, db.Find(&items).Error
}

func GetItemByID(db *gorm.DB, id int) (item Item, err error) {
	// Baca item berdasarkan ID
	return item, db.First(&item, id).Error
}

func GetItemDetails(db *gorm.DB) (itemDetails []ItemDetail, err error) {
	// Baca semua item detail
	return itemDetails, db.Find(&itemDetails).Error
}

func GetItemDetailByID(db *gorm.DB, id int) (itemDetail ItemDetail, err error) {
	// Baca item detail berdasarkan ID
	return itemDetail, db.First(&itemDetail, id).Error
}

// (U) Update
func UpdateItem(db *gorm.DB, item *Item) (err error) {
	// Perbarui item
	return db.Save(item).Error
}

func UpdateItemDetail(db *gorm.DB, itemDetail *ItemDetail) (err error) {
	// Perbarui item detail
	return db.Save(itemDetail).Error
}

// (D) Delete
func DeleteItemByID(db *gorm.DB, id uint) (err error) {
	// Hapus item berdasarkan ID
	return db.Delete(&Item{}, id).Error
}

func DeleteItemDetailByID(db *gorm.DB, id uint) (err error) {
	// Hapus item detail berdasarkan ID
	return db.Delete(&ItemDetail{}, id).Error
}
