package controller

import (
	"gin-example/model"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
    var albums []model.Album

    albums = model.GetAlbums()

    c.JSON(http.StatusOK, gin.H{
        "albums": albums,
    })
}

func GetAlbum(c *gin.Context) {
    var result model.Album
    id, err := strconv.Atoi(c.Param("id"))

    if err != nil {
        log.Println("GetAlbum(): id should be an integer")
        c.JSON(http.StatusBadRequest,
            gin.H{
                "message": "id should be an integer",
            })
        return
    }

    if id < 0 {
        log.Println("GetAlbum(): %d should be a positive integer", id)
        c.JSON(http.StatusBadRequest,
            gin.H{
                "message": "id should be a positive integer",
            })
        return
    }


    result = model.GetAlbum(id)
    log.Println("GetAlbum(): {", result.ID, result.Title, result.Artist, "}")

    if result.Artist == "" && result.Title == "" {
        c.JSON(http.StatusNotFound, gin.H{
            "message": "album not found",
        })
        return
    }


    c.JSON(http.StatusOK, gin.H{
        "ID": result.ID,
        "Title": result.Title,
        "Artist": result.Artist,
    })
}

func PostAlbum(c *gin.Context) {
    var album model.Album

    album.Title = c.PostForm("title")
    album.Artist = c.PostForm("artist")

    if album.Title == "" || album.Artist == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "title and artist are required",
        })
        return
    }

    model.PostAlbum(album)

    c.JSON(http.StatusAccepted, gin.H{
        "status": "accepted",
        "ID": album.ID,
        "Title": album.Title,
        "Artist": album.Artist,
    })
}

func PutAlbum(c *gin.Context)  {
    var album model.Album
    id, err := strconv.Atoi(c.Param("id"))

    if err != nil {
        c.JSON(http.StatusBadRequest,
            gin.H{
                "message": "id should be an integer",
            })
        return
    }

    if id < 0 {
        c.JSON(http.StatusBadRequest,
            gin.H{
                "message": "id should be a positive integer",
            })
        return
    }

    newTitle := c.PostForm("title")
    newArtist := c.PostForm("artist")

    if newTitle == "" || newArtist == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "title and artist are required",
        })
        return
    }

    album = model.GetAlbum(id)

    album.ID = uint(id)
    album.Title = newTitle
    album.Artist = newArtist

    model.UpdateAlbum(album)
    c.JSON(http.StatusAccepted, gin.H{
        "status": "accepted",
        "ID": album.ID,
        "Title": album.Title,
        "Artist": album.Artist,
    })
}

func DeleteAlbum(c *gin.Context) {
    var album model.Album
    id, err := strconv.Atoi(c.Param("id"))

    if err != nil {
        log.Println("GetAlbum(): id should be an integer")
        c.JSON(http.StatusBadRequest,
            gin.H{
                "message": "id should be an integer",
            })
        return
    }

    if id < 0 {
        log.Println("GetAlbum(): %d should be a positive integer", id)
        c.JSON(http.StatusBadRequest,
            gin.H{
                "message": "id should be a positive integer",
            })
        return
    }

    album = model.DeleteAlbum(id)
    c.JSON(http.StatusOK, gin.H{
        "status": "deleted",
        "ID": album.ID,
        "Title": album.Title,
        "Artist": album.Artist,
    })

}
