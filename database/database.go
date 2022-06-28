package main

import (
	db "github.com/sonyarouje/simdb"
)

type Album struct {
	AlbumID string  `json:"albumid"`
	Title   string  `json:"title"`
	Artist  string  `json:"artist"`
	Price   float32 `json:"price"`
}

func (c Album) ID() (jsonField string, value interface{}) {
	value = c.ID
	jsonField = "albumid"
	return
}

func Inser(c Album) error {

	driver, err := db.New("data")

	if err != nil {
		return err
	}

	driver.Insert(c)

	return nil
}

// func main() {
// 	driver, err := db.New("data")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var defaultDatabase = []Album{
// 		{AlbumID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 		{AlbumID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 		{AlbumID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// 	}

// 	for _, v := range defaultDatabase {
// 		// album := Album{
// 		// 	AlbumID: v.AlbumID,
// 		// 	Title: v.Title,
// 		// 	Artist: v.Artist,
// 		// 	Price: v.Price,
// 		// }
// 		driver.Insert(v)
// 	}

// 	fmt.Println("INSERT OK")

// }
