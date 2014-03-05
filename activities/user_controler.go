package activities

import (
  "./models"
  "github.com/coocood/qbs"
  _ "github.com/lib/pq"
  "github.com/stretchr/goweb"
  "github.com/stretchr/goweb/context"
  "net/http"
  "strconv"
)

/*
	RESTful example
*/

type UsersController struct {
}

// Before gets called before any other method.
func (r *UsersController) Before(ctx context.Context) error {

  // set a Things specific header
  ctx.HttpResponseWriter().Header().Set("X-Things-Controller", "true")

  return nil

}

func (r *UsersController) Create(ctx context.Context) error {

  data, dataErr := ctx.RequestData()

  if dataErr != nil {
    return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, dataErr.Error())
  }

  dataMap := data.(map[string]interface{})
  //models.CreateUser(dataMap)
  _, err := models.Create(dataMap, new(models.User))
  if err != nil {
    return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
  }

  return goweb.Respond.WithStatus(ctx, http.StatusCreated)

}

func (r *UsersController) ReadMany(ctx context.Context) error {
  q, err := qbs.GetQbs()
  if err != nil {
    return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
  }

  var users []*models.User
  err = q.FindAll(&users)

  return goweb.API.RespondWithData(ctx, users)
}

func (r *UsersController) Read(ids string, ctx context.Context) error {
  var id int64
  var user *models.User = new(models.User)
  id, err := strconv.ParseInt(ids, 0, 64)
  _, err = models.FindById(id, user)

  if err != nil {
    return goweb.Respond.WithStatus(ctx, 404)
    //return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
  }

  if user != nil {
    return goweb.API.RespondWithData(ctx, user)
  }
  return goweb.Respond.WithStatus(ctx, http.StatusNotFound)

}

func (r *UsersController) DeleteMany(ctx context.Context) error {

  return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
  //return goweb.Respond.WithOK(ctx)
}

func (r *UsersController) Delete(id string, ctx context.Context) error {

  return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
  //return goweb.Respond.WithOK(ctx)

}
