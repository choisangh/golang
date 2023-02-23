package api

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/choisangh/gin_project/internal/global"
	"github.com/gin-gonic/gin"
)

type RequestFile struct {
	File *multipart.FileHeader `form:"file"`
}

type RequestURI struct {
	ImageID string `uri:"id"`
}

type Response struct {
	Res string `json:"res"`
}

func CreateImage(c *gin.Context) {
	req := RequestFile{}
	//인자 주소값을 줘야함 //Bind는 리스폰스하고 에러 -> ShouldBind가 에러 핸들링 자유도가 높아서 ShouldBind 추천
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid"})

		return
	}

	if err := c.SaveUploadedFile(req.File, "images/"+strconv.FormatInt(global.MaxImageNumber, 10)); err != nil { //int -> string 파일명
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid"})

		return
	}
	global.MaxImageNumber++

	c.JSON(http.StatusOK, Response{Res: "success"})

}

func ReadImage(c *gin.Context) {
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid"})
		return
	}
	fmt.Println(reqURI.ImageID)
	c.File("images/" + reqURI.ImageID)

}
