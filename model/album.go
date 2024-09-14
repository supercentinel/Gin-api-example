package model

type Album struct {
    ID uint `json:"id"`
    Title string `json:"title"`
    Artist string `json:"artist"`
}

func GetAlbums() []Album {
    var albums []Album

    DB.Find(&albums)

    return albums
}

func GetAlbum(id int) Album {
    var result Album

    DB.First(&result, id)

    return result
}

func PostAlbum(album Album) Album {
    DB.Create(&album)

    return album
}

func UpdateAlbum(album Album) Album {
    DB.Save(&album)

    return album
}

func DeleteAlbum(id int) Album {
    var album Album

    DB.First(&album, id)
    DB.Delete(&album)

    return album
}
