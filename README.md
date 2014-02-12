go-search
=========

An example using [Ferret](https://github.com/argusdusty/Ferret) and [Martini](https://github.com/codegangsta/martini) to make a really fast Search Service.

You can see it in action over on [Heroku](http://search-shakespeare.herokuapp.com).

The text of Shakespeare is provided by [Project Gutenberg](http://www.gutenberg.org).

### First Go Project?

I'd suggest by reading a blog post I wrote on [setting up your environment](http://www.zhubert.com/blog/2014/02/11/setting-up-go/) and then just

```
go get -u github.com/zhubert/go-search
```

That will grab this repo, put it into the GOPATH where it belongs and install go-search into $GOPATH/bin, which should be in your PATH.

### Heroku Deployable

This repo also shows how to use the Go buildpack to deploy to Heroku, just like in this [blog post](http://blog.wercker.com/2013/07/10/deploying-golang-to-heroku.html)
