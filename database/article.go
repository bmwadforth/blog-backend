package database

import (
	"blog-backend/mapper"
	"blog-backend/models"
	"blog-backend/util"
	"context"
	"errors"
)

func GetArticle(articleId string) util.DataResponse[models.ArticleModel] {
	var article models.ArticleModel
	dataResponse := util.NewDataResponse("success", article)
	ctx := context.Background()
	client, _ := createClient(ctx)
	defer client.Close()

	docs, err := client.Collection("articles").Where("ArticleId", "==", articleId).Documents(ctx).GetAll()
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	if len(docs) == 0 {
		dataResponse.SetError(errors.New("no articles found"), util.DbresultNotFound)
		return dataResponse
	}

	if len(docs) > 1 {
		dataResponse.SetError(errors.New("error multiple articles found"), util.DbresultError)
		return dataResponse
	}

	err = docs[0].DataTo(&article)
	if err != nil {
		dataResponse.SetError(errors.New("error unable to deserialize record"), util.DbresultError)
		return dataResponse
	}
	dataResponse.SetData(article)

	return dataResponse
}

func GetArticles() util.DataResponse[[]models.ArticleModel] {
	var articles []models.ArticleModel
	var article models.ArticleModel
	dataResponse := util.NewDataResponse("successfully read articles", articles)
	ctx := context.Background()
	client, _ := createClient(ctx)
	defer client.Close()

	docs, err := client.Collection("articles").Documents(ctx).GetAll()
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	for _, doc := range docs {
		doc.DataTo(&article)
		articles = append(articles, article)
	}

	dataResponse.SetData(articles)

	return dataResponse
}

func CreateArticle(request models.CreateArticleRequest) util.DataResponse[string] {
	dataResponse := util.NewDataResponse("successfully created article", "")
	ctx := context.Background()
	client, _ := createClient(ctx)
	defer client.Close()

	article := mapper.MapArticleCreatRequest(request)
	_, _, err := client.Collection("articles").Add(ctx, article)
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	dataResponse.SetData(article.ArticleId)

	return dataResponse
}