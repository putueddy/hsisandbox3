package main

import (
	"fmt"

	tugas3 "sandboxhsi3.com/golang/tugas"
)

func cetakItem(item tugas3.Item) {
	fmt.Printf("Item ID: %v, Nama: %s, Status: %s, Amount: %v\n", item.ID, item.Name, item.Status, item.Amount)
}

func cetakItemDetail(itemDetail tugas3.ItemDetail) {
	fmt.Printf("Item Detail ID: %v, Nama: %s, ItemID: %v\n", itemDetail.ID, itemDetail.Name, itemDetail.ItemID)
}

func main() {
	// 1. Buatlah 2 tabel dengan struktur one-to-many!
	//
	// CREATE SEQUENCE IF NOT EXISTS item_id_seq;
	// CREATE TABLE item (
	//    id integer PRIMARY KEY DEFAULT nextval('item_id_seq'),
	//    name text NOT NULL,
	//    status text NOT NULL,
	//    amount integer NOT NULL,
	//    CHECK (amount >= 0)
	// );
	//
	// CREATE SEQUENCE IF NOT EXISTS item_detail_id_seq;
	// CREATE TABLE item_detail (
	//    id integer PRIMARY KEY DEFAULT nextval('item_detail_id_seq'),
	//    item_id INTEGER NOT NULL REFERENCES item ON DELETE CASCADE,
	//    name text NOT NULL
	// );

	// 2. Buatlah CRUD untuk 2 tabel tersebut!
	db, err := tugas3.ConnectDatabase()
	if err != nil {
		panic("Gagal terkoneksi ke database!")
	}

	tugas3.DBMigrate(db)

	// (C) Create
	item := tugas3.Item{
		Name:   "Laptop",
		Status: "Kosong",
		Amount: 0,
		ItemDetails: []tugas3.ItemDetail{
			{Name: "ThinkPad"},
			{Name: "Macbook"},
		},
	}
	err = tugas3.CreateItem(db, &item)
	if err != nil {
		panic("Gagal membuat item!")
	}

	// (R) Read
	bacaItem, err := tugas3.GetItemByID(db, 1)
	if err != nil {
		panic("Gagal membaca item berdasarkan ID")
	}
	cetakItem(bacaItem)

	// 3. Buatlah transactions untuk menambah (insert) item_detail dan memperbarui (update) item.status dan item.amount!
	itemDetail := tugas3.ItemDetail{
		Name:   "Zenbook",
		ItemID: bacaItem.ID,
	}
	err = tugas3.CreateItemDetail(db, &itemDetail)
	if err != nil {
		panic("Gagal membuat item detail!")
	}
	// (U) Update
	bacaItem.Status = "Ada"
	bacaItem.Amount = 3
	err = tugas3.UpdateItem(db, &bacaItem)
	if err != nil {
		panic("Gagal membuat item detail!")
	}
	cetakItem(bacaItem)

	// (D) Sebelum
	itemDetails, err := tugas3.GetItemDetails(db)
	if err != nil {
		panic("Gagal membaca item detail!")
	}
	for _, element := range itemDetails {
		cetakItemDetail(element)
	}
	// (D) Delete
	err = tugas3.DeleteItemDetailByID(db, itemDetails[len(itemDetails)-1].ID)
	if err != nil {
		panic("Gagal menghapus item detail!")
	}
	// (D) Setelah
	itemDetails, err = tugas3.GetItemDetails(db)
	if err != nil {
		panic("Gagal membaca item detail!")
	}
	for _, element := range itemDetails {
		cetakItemDetail(element)
	}

	// 4. Buatlah batch insert untuk tugas nomor 3!
	item2 := tugas3.Item{
		Name:   "Gawai",
		Status: "Ada",
		Amount: 5,
		ItemDetails: []tugas3.ItemDetail{
			{Name: "Samsung"},
		},
	}
	err = tugas3.CreateItem(db, &item2)
	if err != nil {
		panic("Gagal membuat item!")
	}
	allItem, err := tugas3.GetItems(db)
	if err != nil {
		panic("Gagal membaca item!")
	}

	lastItem := allItem[len(allItem)-1]
	itemDetail2 := []tugas3.ItemDetail{
		{Name: "Apple", ItemID: lastItem.ID},
		{Name: "Oppo", ItemID: lastItem.ID},
		{Name: "Xiaomi", ItemID: lastItem.ID},
		{Name: "Huawei", ItemID: lastItem.ID},
	}
	// Batch Insert
	db.CreateInBatches(itemDetail2, len(itemDetail2))
	// Tampilkan item detail terkini
	itemDetails, err = tugas3.GetItemDetails(db)
	if err != nil {
		panic("Gagal membaca item detail!")
	}
	for _, element := range itemDetails {
		cetakItemDetail(element)
	}
}
