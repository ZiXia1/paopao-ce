// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v3.1.1

package v1

import (
	"net/http"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

type Priv interface {
	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	UnfollowTopic(*web.UnfollowTopicReq) mir.Error
	FollowTopic(*web.FollowTopicReq) mir.Error
	StickTopic(*web.StickTopicReq) (*web.StickTopicResp, mir.Error)
	ThumbsDownTweetReply(*web.TweetReplyThumbsReq) mir.Error
	ThumbsUpTweetReply(*web.TweetReplyThumbsReq) mir.Error
	ThumbsDownTweetComment(*web.TweetCommentThumbsReq) mir.Error
	ThumbsUpTweetComment(*web.TweetCommentThumbsReq) mir.Error
	DeleteCommentReply(*web.DeleteCommentReplyReq) mir.Error
	CreateCommentReply(*web.CreateCommentReplyReq) (*web.CreateCommentReplyResp, mir.Error)
	DeleteComment(*web.DeleteCommentReq) mir.Error
	CreateComment(*web.CreateCommentReq) (*web.CreateCommentResp, mir.Error)
	VisibleTweet(*web.VisibleTweetReq) (*web.VisibleTweetResp, mir.Error)
	StickTweet(*web.StickTweetReq) (*web.StickTweetResp, mir.Error)
	LockTweet(*web.LockTweetReq) (*web.LockTweetResp, mir.Error)
	CollectionTweet(*web.CollectionTweetReq) (*web.CollectionTweetResp, mir.Error)
	StarTweet(*web.StarTweetReq) (*web.StarTweetResp, mir.Error)
	DeleteTweet(*web.DeleteTweetReq) mir.Error
	CreateTweet(*web.CreateTweetReq) (*web.CreateTweetResp, mir.Error)
	DownloadAttachment(*web.DownloadAttachmentReq) (*web.DownloadAttachmentResp, mir.Error)
	DownloadAttachmentPrecheck(*web.DownloadAttachmentPrecheckReq) (*web.DownloadAttachmentPrecheckResp, mir.Error)
	UploadAttachment(*web.UploadAttachmentReq) (*web.UploadAttachmentResp, mir.Error)

	mustEmbedUnimplementedPrivServant()
}

type PrivBinding interface {
	BindUnfollowTopic(*gin.Context) (*web.UnfollowTopicReq, mir.Error)
	BindFollowTopic(*gin.Context) (*web.FollowTopicReq, mir.Error)
	BindStickTopic(*gin.Context) (*web.StickTopicReq, mir.Error)
	BindThumbsDownTweetReply(*gin.Context) (*web.TweetReplyThumbsReq, mir.Error)
	BindThumbsUpTweetReply(*gin.Context) (*web.TweetReplyThumbsReq, mir.Error)
	BindThumbsDownTweetComment(*gin.Context) (*web.TweetCommentThumbsReq, mir.Error)
	BindThumbsUpTweetComment(*gin.Context) (*web.TweetCommentThumbsReq, mir.Error)
	BindDeleteCommentReply(*gin.Context) (*web.DeleteCommentReplyReq, mir.Error)
	BindCreateCommentReply(*gin.Context) (*web.CreateCommentReplyReq, mir.Error)
	BindDeleteComment(*gin.Context) (*web.DeleteCommentReq, mir.Error)
	BindCreateComment(*gin.Context) (*web.CreateCommentReq, mir.Error)
	BindVisibleTweet(*gin.Context) (*web.VisibleTweetReq, mir.Error)
	BindStickTweet(*gin.Context) (*web.StickTweetReq, mir.Error)
	BindLockTweet(*gin.Context) (*web.LockTweetReq, mir.Error)
	BindCollectionTweet(*gin.Context) (*web.CollectionTweetReq, mir.Error)
	BindStarTweet(*gin.Context) (*web.StarTweetReq, mir.Error)
	BindDeleteTweet(*gin.Context) (*web.DeleteTweetReq, mir.Error)
	BindCreateTweet(*gin.Context) (*web.CreateTweetReq, mir.Error)
	BindDownloadAttachment(*gin.Context) (*web.DownloadAttachmentReq, mir.Error)
	BindDownloadAttachmentPrecheck(*gin.Context) (*web.DownloadAttachmentPrecheckReq, mir.Error)
	BindUploadAttachment(*gin.Context) (*web.UploadAttachmentReq, mir.Error)

	mustEmbedUnimplementedPrivBinding()
}

type PrivRender interface {
	RenderUnfollowTopic(*gin.Context, mir.Error)
	RenderFollowTopic(*gin.Context, mir.Error)
	RenderStickTopic(*gin.Context, *web.StickTopicResp, mir.Error)
	RenderThumbsDownTweetReply(*gin.Context, mir.Error)
	RenderThumbsUpTweetReply(*gin.Context, mir.Error)
	RenderThumbsDownTweetComment(*gin.Context, mir.Error)
	RenderThumbsUpTweetComment(*gin.Context, mir.Error)
	RenderDeleteCommentReply(*gin.Context, mir.Error)
	RenderCreateCommentReply(*gin.Context, *web.CreateCommentReplyResp, mir.Error)
	RenderDeleteComment(*gin.Context, mir.Error)
	RenderCreateComment(*gin.Context, *web.CreateCommentResp, mir.Error)
	RenderVisibleTweet(*gin.Context, *web.VisibleTweetResp, mir.Error)
	RenderStickTweet(*gin.Context, *web.StickTweetResp, mir.Error)
	RenderLockTweet(*gin.Context, *web.LockTweetResp, mir.Error)
	RenderCollectionTweet(*gin.Context, *web.CollectionTweetResp, mir.Error)
	RenderStarTweet(*gin.Context, *web.StarTweetResp, mir.Error)
	RenderDeleteTweet(*gin.Context, mir.Error)
	RenderCreateTweet(*gin.Context, *web.CreateTweetResp, mir.Error)
	RenderDownloadAttachment(*gin.Context, *web.DownloadAttachmentResp, mir.Error)
	RenderDownloadAttachmentPrecheck(*gin.Context, *web.DownloadAttachmentPrecheckResp, mir.Error)
	RenderUploadAttachment(*gin.Context, *web.UploadAttachmentResp, mir.Error)

	mustEmbedUnimplementedPrivRender()
}

// RegisterPrivServant register Priv servant to gin
func RegisterPrivServant(e *gin.Engine, s Priv, b PrivBinding, r PrivRender) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("POST", "/topic/unfollow", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindUnfollowTopic(c)
		if err != nil {
			r.RenderUnfollowTopic(c, err)
			return
		}
		r.RenderUnfollowTopic(c, s.UnfollowTopic(req))
	})

	router.Handle("POST", "/topic/follow", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindFollowTopic(c)
		if err != nil {
			r.RenderFollowTopic(c, err)
			return
		}
		r.RenderFollowTopic(c, s.FollowTopic(req))
	})

	router.Handle("POST", "/topic/stick", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindStickTopic(c)
		if err != nil {
			r.RenderStickTopic(c, nil, err)
			return
		}
		resp, err := s.StickTopic(req)
		r.RenderStickTopic(c, resp, err)
	})

	router.Handle("POST", "/tweet/reply/thumbsdown", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindThumbsDownTweetReply(c)
		if err != nil {
			r.RenderThumbsDownTweetReply(c, err)
			return
		}
		r.RenderThumbsDownTweetReply(c, s.ThumbsDownTweetReply(req))
	})

	router.Handle("POST", "/tweet/reply/thumbsup", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindThumbsUpTweetReply(c)
		if err != nil {
			r.RenderThumbsUpTweetReply(c, err)
			return
		}
		r.RenderThumbsUpTweetReply(c, s.ThumbsUpTweetReply(req))
	})

	router.Handle("POST", "/tweet/comment/thumbsdown", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindThumbsDownTweetComment(c)
		if err != nil {
			r.RenderThumbsDownTweetComment(c, err)
			return
		}
		r.RenderThumbsDownTweetComment(c, s.ThumbsDownTweetComment(req))
	})

	router.Handle("POST", "/tweet/comment/thumbsup", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindThumbsUpTweetComment(c)
		if err != nil {
			r.RenderThumbsUpTweetComment(c, err)
			return
		}
		r.RenderThumbsUpTweetComment(c, s.ThumbsUpTweetComment(req))
	})

	router.Handle("DELETE", "/post/comment/reply", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindDeleteCommentReply(c)
		if err != nil {
			r.RenderDeleteCommentReply(c, err)
			return
		}
		r.RenderDeleteCommentReply(c, s.DeleteCommentReply(req))
	})

	router.Handle("POST", "/post/comment/reply", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindCreateCommentReply(c)
		if err != nil {
			r.RenderCreateCommentReply(c, nil, err)
			return
		}
		resp, err := s.CreateCommentReply(req)
		r.RenderCreateCommentReply(c, resp, err)
	})

	router.Handle("DELETE", "/post/comment", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindDeleteComment(c)
		if err != nil {
			r.RenderDeleteComment(c, err)
			return
		}
		r.RenderDeleteComment(c, s.DeleteComment(req))
	})

	router.Handle("POST", "/post/comment", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindCreateComment(c)
		if err != nil {
			r.RenderCreateComment(c, nil, err)
			return
		}
		resp, err := s.CreateComment(req)
		r.RenderCreateComment(c, resp, err)
	})

	router.Handle("POST", "/post/visibility", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindVisibleTweet(c)
		if err != nil {
			r.RenderVisibleTweet(c, nil, err)
			return
		}
		resp, err := s.VisibleTweet(req)
		r.RenderVisibleTweet(c, resp, err)
	})

	router.Handle("POST", "/post/stick", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindStickTweet(c)
		if err != nil {
			r.RenderStickTweet(c, nil, err)
			return
		}
		resp, err := s.StickTweet(req)
		r.RenderStickTweet(c, resp, err)
	})

	router.Handle("POST", "/post/lock", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindLockTweet(c)
		if err != nil {
			r.RenderLockTweet(c, nil, err)
			return
		}
		resp, err := s.LockTweet(req)
		r.RenderLockTweet(c, resp, err)
	})

	router.Handle("POST", "/post/collection", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindCollectionTweet(c)
		if err != nil {
			r.RenderCollectionTweet(c, nil, err)
			return
		}
		resp, err := s.CollectionTweet(req)
		r.RenderCollectionTweet(c, resp, err)
	})

	router.Handle("POST", "/post/star", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindStarTweet(c)
		if err != nil {
			r.RenderStarTweet(c, nil, err)
			return
		}
		resp, err := s.StarTweet(req)
		r.RenderStarTweet(c, resp, err)
	})

	router.Handle("DELETE", "/post", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindDeleteTweet(c)
		if err != nil {
			r.RenderDeleteTweet(c, err)
			return
		}
		r.RenderDeleteTweet(c, s.DeleteTweet(req))
	})

	router.Handle("POST", "/post", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindCreateTweet(c)
		if err != nil {
			r.RenderCreateTweet(c, nil, err)
			return
		}
		resp, err := s.CreateTweet(req)
		r.RenderCreateTweet(c, resp, err)
	})

	router.Handle("GET", "/attachment", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindDownloadAttachment(c)
		if err != nil {
			r.RenderDownloadAttachment(c, nil, err)
			return
		}
		resp, err := s.DownloadAttachment(req)
		r.RenderDownloadAttachment(c, resp, err)
	})

	router.Handle("GET", "/attachment/precheck", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindDownloadAttachmentPrecheck(c)
		if err != nil {
			r.RenderDownloadAttachmentPrecheck(c, nil, err)
			return
		}
		resp, err := s.DownloadAttachmentPrecheck(req)
		r.RenderDownloadAttachmentPrecheck(c, resp, err)
	})

	router.Handle("POST", "/attachment", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindUploadAttachment(c)
		if err != nil {
			r.RenderUploadAttachment(c, nil, err)
			return
		}
		resp, err := s.UploadAttachment(req)
		r.RenderUploadAttachment(c, resp, err)
	})

}

// UnimplementedPrivServant can be embedded to have forward compatible implementations.
type UnimplementedPrivServant struct {
}

func (UnimplementedPrivServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedPrivServant) UnfollowTopic(req *web.UnfollowTopicReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) FollowTopic(req *web.FollowTopicReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) StickTopic(req *web.StickTopicReq) (*web.StickTopicResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) ThumbsDownTweetReply(req *web.TweetReplyThumbsReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) ThumbsUpTweetReply(req *web.TweetReplyThumbsReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) ThumbsDownTweetComment(req *web.TweetCommentThumbsReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) ThumbsUpTweetComment(req *web.TweetCommentThumbsReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) DeleteCommentReply(req *web.DeleteCommentReplyReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) CreateCommentReply(req *web.CreateCommentReplyReq) (*web.CreateCommentReplyResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) DeleteComment(req *web.DeleteCommentReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) CreateComment(req *web.CreateCommentReq) (*web.CreateCommentResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) VisibleTweet(req *web.VisibleTweetReq) (*web.VisibleTweetResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) StickTweet(req *web.StickTweetReq) (*web.StickTweetResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) LockTweet(req *web.LockTweetReq) (*web.LockTweetResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) CollectionTweet(req *web.CollectionTweetReq) (*web.CollectionTweetResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) StarTweet(req *web.StarTweetReq) (*web.StarTweetResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) DeleteTweet(req *web.DeleteTweetReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) CreateTweet(req *web.CreateTweetReq) (*web.CreateTweetResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) DownloadAttachment(req *web.DownloadAttachmentReq) (*web.DownloadAttachmentResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) DownloadAttachmentPrecheck(req *web.DownloadAttachmentPrecheckReq) (*web.DownloadAttachmentPrecheckResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) UploadAttachment(req *web.UploadAttachmentReq) (*web.UploadAttachmentResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPrivServant) mustEmbedUnimplementedPrivServant() {}

// UnimplementedPrivRender can be embedded to have forward compatible implementations.
type UnimplementedPrivRender struct {
	RenderAny func(*gin.Context, any, mir.Error)
}

func (r *UnimplementedPrivRender) RenderUnfollowTopic(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedPrivRender) RenderFollowTopic(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedPrivRender) RenderStickTopic(c *gin.Context, data *web.StickTopicResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) RenderThumbsDownTweetReply(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedPrivRender) RenderThumbsUpTweetReply(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedPrivRender) RenderThumbsDownTweetComment(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedPrivRender) RenderThumbsUpTweetComment(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedPrivRender) RenderDeleteCommentReply(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedPrivRender) RenderCreateCommentReply(c *gin.Context, data *web.CreateCommentReplyResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) RenderDeleteComment(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedPrivRender) RenderCreateComment(c *gin.Context, data *web.CreateCommentResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) RenderVisibleTweet(c *gin.Context, data *web.VisibleTweetResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) RenderStickTweet(c *gin.Context, data *web.StickTweetResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) RenderLockTweet(c *gin.Context, data *web.LockTweetResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) RenderCollectionTweet(c *gin.Context, data *web.CollectionTweetResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) RenderStarTweet(c *gin.Context, data *web.StarTweetResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) RenderDeleteTweet(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedPrivRender) RenderCreateTweet(c *gin.Context, data *web.CreateTweetResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) RenderDownloadAttachment(c *gin.Context, data *web.DownloadAttachmentResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) RenderDownloadAttachmentPrecheck(c *gin.Context, data *web.DownloadAttachmentPrecheckResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) RenderUploadAttachment(c *gin.Context, data *web.UploadAttachmentResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPrivRender) mustEmbedUnimplementedPrivRender() {}

// UnimplementedPrivBinding can be embedded to have forward compatible implementations.
type UnimplementedPrivBinding struct {
	BindAny func(*gin.Context, any) mir.Error
}

func (b *UnimplementedPrivBinding) BindUnfollowTopic(c *gin.Context) (*web.UnfollowTopicReq, mir.Error) {
	obj := new(web.UnfollowTopicReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindFollowTopic(c *gin.Context) (*web.FollowTopicReq, mir.Error) {
	obj := new(web.FollowTopicReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindStickTopic(c *gin.Context) (*web.StickTopicReq, mir.Error) {
	obj := new(web.StickTopicReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindThumbsDownTweetReply(c *gin.Context) (*web.TweetReplyThumbsReq, mir.Error) {
	obj := new(web.TweetReplyThumbsReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindThumbsUpTweetReply(c *gin.Context) (*web.TweetReplyThumbsReq, mir.Error) {
	obj := new(web.TweetReplyThumbsReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindThumbsDownTweetComment(c *gin.Context) (*web.TweetCommentThumbsReq, mir.Error) {
	obj := new(web.TweetCommentThumbsReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindThumbsUpTweetComment(c *gin.Context) (*web.TweetCommentThumbsReq, mir.Error) {
	obj := new(web.TweetCommentThumbsReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindDeleteCommentReply(c *gin.Context) (*web.DeleteCommentReplyReq, mir.Error) {
	obj := new(web.DeleteCommentReplyReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindCreateCommentReply(c *gin.Context) (*web.CreateCommentReplyReq, mir.Error) {
	obj := new(web.CreateCommentReplyReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindDeleteComment(c *gin.Context) (*web.DeleteCommentReq, mir.Error) {
	obj := new(web.DeleteCommentReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindCreateComment(c *gin.Context) (*web.CreateCommentReq, mir.Error) {
	obj := new(web.CreateCommentReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindVisibleTweet(c *gin.Context) (*web.VisibleTweetReq, mir.Error) {
	obj := new(web.VisibleTweetReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindStickTweet(c *gin.Context) (*web.StickTweetReq, mir.Error) {
	obj := new(web.StickTweetReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindLockTweet(c *gin.Context) (*web.LockTweetReq, mir.Error) {
	obj := new(web.LockTweetReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindCollectionTweet(c *gin.Context) (*web.CollectionTweetReq, mir.Error) {
	obj := new(web.CollectionTweetReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindStarTweet(c *gin.Context) (*web.StarTweetReq, mir.Error) {
	obj := new(web.StarTweetReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindDeleteTweet(c *gin.Context) (*web.DeleteTweetReq, mir.Error) {
	obj := new(web.DeleteTweetReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindCreateTweet(c *gin.Context) (*web.CreateTweetReq, mir.Error) {
	obj := new(web.CreateTweetReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindDownloadAttachment(c *gin.Context) (*web.DownloadAttachmentReq, mir.Error) {
	obj := new(web.DownloadAttachmentReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindDownloadAttachmentPrecheck(c *gin.Context) (*web.DownloadAttachmentPrecheckReq, mir.Error) {
	obj := new(web.DownloadAttachmentPrecheckReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) BindUploadAttachment(c *gin.Context) (*web.UploadAttachmentReq, mir.Error) {
	obj := new(web.UploadAttachmentReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPrivBinding) mustEmbedUnimplementedPrivBinding() {}
