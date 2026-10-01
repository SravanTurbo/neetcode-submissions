type Tweet struct {
	tweetID   int
	timestamp int
}

type Twitter struct {
	time 	  int
	tweets    map[int][]*Tweet        //userID vs Tweets
	followMap map[int]map[int]bool    //followerID vs [followeeIDs]bool
	feedLimit int
}

func Constructor() Twitter {
	return Twitter{
		time:      0,
		tweets:    make(map[int][]*Tweet),
		followMap: make(map[int]map[int]bool),
		feedLimit: 10,
	}
}

//O(1)
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.time++

	tweet := &Tweet{
		tweetID: tweetId,
		timestamp: this.time,
	}

    this.tweets[userId] = append(this.tweets[userId], []*Tweet{tweet}...)
}

//O(nlog(n)) - where n is # of tweets
func (this *Twitter) GetNewsFeed(userId int) []int {
    tweetsForFeed := []*Tweet{}
	tweetsForFeed = append(tweetsForFeed, this.tweets[userId]...)

	for followeeId, _ := range this.followMap[userId] {
		tweetsForFeed = append(tweetsForFeed, this.tweets[followeeId]...)
	}

	sort.Slice(tweetsForFeed, func(i, j int) bool {
		return tweetsForFeed[i].timestamp > tweetsForFeed[j].timestamp
	})

	feed := []int{}
	for i := range this.feedLimit {
		if i < len(tweetsForFeed) {
			feed = append(feed, tweetsForFeed[i].tweetID)
		}
	}

	return feed
}


//O(1)
func (this *Twitter) Follow(followerId int, followeeId int)  {
	if followerId == followeeId {
		return
	}

	following := this.followMap[followerId]
	if following == nil {
		following = make(map[int]bool)
		this.followMap[followerId] = following
	}

	following[followeeId] = true
}

//O(1)
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	following := this.followMap[followerId]
	if following == nil {
		return
	}

	_, exists := following[followeeId]
	if !exists {
		return
	}

	delete(following, followeeId)
}
