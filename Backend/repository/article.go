package repository

import (
	"linblog/model"
	"linblog/repository/articleSource"
	"time"
)

type ArticleRepository struct {
}

func (article *ArticleRepository) GetArticles(page, size int) ([]*model.Article, int, error) {
	//先拿到所有的分类
	cates, err := articleSource.Article.GetCategories()
	if err != nil {
		return nil, 0, err
	}
	//按照所有分类拿到所有的文章名
	res := make([]*model.Article, 0)
	min_index, max_index := (page-1)*size, page*size
	totalLength := 0
	for _, cate := range cates {
		articles, err := articleSource.Article.GetArticleNames(cate)
		if err != nil {
			return nil, 0, err
		}
		articleLength := len(articles)
		//小于所需起始条目的直接continue
		if articleLength+totalLength < min_index {
			totalLength += articleLength
			continue
		}

		for _, article := range articles {
			if totalLength > max_index {
				return res, totalLength, nil
			}
			//每个分类对应的文章
			if totalLength >= min_index && totalLength <= max_index {
				content, err := articleSource.Article.GetArticleHtml(cate, article)
				if err != nil {
					return nil, 0, err
				}
				newArticle := &model.Article{
					Id:            totalLength,
					IsTop:         false,
					Banner:        "https://picsum.photos/200",
					IsHot:         true,
					PubTime:       int(time.Now().UnixMilli()),
					Title:         article,
					Summary:       "testtesttesttest",
					Content:       content,
					ViewsCount:    1000,
					CommentsCount: 100,
				}
				res = append(res, newArticle)
			}
			totalLength++
		}
	}
	return res, totalLength, nil
}