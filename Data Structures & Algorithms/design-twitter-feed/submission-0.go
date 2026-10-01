import "time"

type Tweet struct {
	tweetID   int
	timestamp time.Time
}

type Twitter struct {
	tweets    map[int][]*Tweet //userID vs Tweets
	following map[int][]int    //followerID vs followeeIDs
	feedLimit int
}

func Constructor() Twitter {
	return Twitter{
		tweets:    make(map[int][]*Tweet),
		following: make(map[int][]int),
		feedLimit: 10,
	}
}

//O(1)
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	tweet := &Tweet{
		tweetID: tweetId,
		timestamp: time.Now(),
	}

    this.tweets[userId] = append(this.tweets[userId], []*Tweet{tweet}...)
}

//O(tlog(t)) - where t is # of tweets
func (this *Twitter) GetNewsFeed(userId int) []int {
    tweetsForFeed := []*Tweet{}
	tweetsForFeed = append(tweetsForFeed, this.tweets[userId]...)

	for _, followeeId := range this.following[userId] {
		tweetsForFeed = append(tweetsForFeed, this.tweets[followeeId]...)
	}

	sort.Slice(tweetsForFeed, func(i, j int) bool {
		return tweetsForFeed[i].timestamp.After(tweetsForFeed[j].timestamp)
	})

	feed := []int{}
	for i := range this.feedLimit {
		if i < len(tweetsForFeed) {
			feed = append(feed, tweetsForFeed[i].tweetID)
		}
	}

	return feed
}


//O(n) - where n is len of users
func (this *Twitter) Follow(followerId int, followeeId int)  {
    for _, followedId := range this.following[followerId] {
		if followedId == followeeId {
			return
		}
	}

    this.following[followerId] = append(this.following[followerId], []int{followeeId}...)
}

//O(n) - where n is len of users
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    for i, followedId := range this.following[followerId] {
		if followedId == followeeId {
			this.following[followerId] = append(this.following[followerId][:i], 
												this.following[followerId][i+1:]...)
		}
	}
}
