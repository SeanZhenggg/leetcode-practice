package heapAndPq

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"log"
)

type post struct {
	UserId  int
	TweetId int
	Time    int
}

type Twitter struct {
	follows       map[int]map[int]struct{}
	posts         map[int][]post
	followerPosts map[int]*priorityqueue.Queue
	timeCount     int
}

func TwitterConstructor() Twitter {
	return Twitter{
		follows:       make(map[int]map[int]struct{}),
		posts:         make(map[int][]post),
		followerPosts: make(map[int]*priorityqueue.Queue),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.timeCount++
	if this.posts[userId] == nil {
		this.posts[userId] = make([]post, 0)
	}
	this.posts[userId] = append(this.posts[userId], post{TweetId: tweetId, Time: this.timeCount})
	this.createUserFollowerPosts(userId)
	this.followerPosts[userId].Enqueue(post{UserId: userId, TweetId: tweetId, Time: this.timeCount})

	for uId, m := range this.follows {
		if _, found := m[userId]; found {
			this.followerPosts[uId].Enqueue(post{UserId: userId, TweetId: tweetId, Time: this.timeCount})
		}
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	this.createUserFollowerPosts(userId)
	j := 0
	ret := make([]int, 0, 10)
	temp := make([]post, 0, this.followerPosts[userId].Size())
	for i := 0; this.followerPosts[userId].Size() > 0 && j < 10; i++ {
		p, _ := this.followerPosts[userId].Dequeue()
		post, _ := p.(post)
		temp = append(temp, post)
		followUsersList := this.follows[userId]
		if _, found := followUsersList[post.UserId]; userId != post.UserId && !found {
			continue
		}
		ret = append(ret, post.TweetId)
		j++
	}

	for i := 0; i < len(temp); i++ {
		this.followerPosts[userId].Enqueue(temp[i])
	}
	return ret
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.follows[followerId] == nil {
		this.follows[followerId] = make(map[int]struct{})
	}
	if _, found := this.follows[followerId][followeeId]; found {
		return
	}

	this.follows[followerId][followeeId] = struct{}{}

	this.createUserFollowerPosts(followerId)
	for _, followeePost := range this.posts[followeeId] {
		this.followerPosts[followerId].Enqueue(post{UserId: followeeId, TweetId: followeePost.TweetId, Time: followeePost.Time})
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.follows[followerId], followeeId)
}

func (this *Twitter) createUserFollowerPosts(userId int) {
	if this.followerPosts[userId] == nil {
		this.followerPosts[userId] = priorityqueue.NewWith(func(a, b interface{}) int {
			pA, pB := a.(post), b.(post)
			return pB.Time - pA.Time
		})
	}
}

type post2 struct {
	Time    int
	TweetId int
}

type Twitter2 struct {
	follows   map[int]map[int]bool
	posts     map[int][]post2
	timeCount int
}

func TwitterConstructor2() Twitter2 {
	return Twitter2{
		follows: make(map[int]map[int]bool),
		posts:   make(map[int][]post2),
	}
}

func (this *Twitter2) PostTweet(userId int, tweetId int) {
	this.timeCount++
	if this.posts[userId] == nil {
		this.posts[userId] = make([]post2, 0)
	}
	this.posts[userId] = append(this.posts[userId], post2{Time: this.timeCount, TweetId: tweetId})
}

func (this *Twitter2) GetNewsFeed(userId int) []int {
	ret := make([]int, 0, 10)

	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		pA, pB := a.([]int), b.([]int)
		return pB[0] - pA[0]
	})

	if this.follows[userId] == nil {
		this.follows[userId] = make(map[int]bool)
	}
	this.follows[userId][userId] = true

	for followeeId := range this.follows[userId] {
		posts := this.posts[followeeId]
		if len(posts) > 0 {
			index := len(posts) - 1
			pq.Enqueue([]int{posts[index].Time, followeeId, posts[index].TweetId, index - 1})
		}
	}

	for pq.Size() > 0 && len(ret) < 10 {
		value, _ := pq.Dequeue()
		pInfo := value.([]int)
		followeeId, tweetId, index := pInfo[1], pInfo[2], pInfo[3]
		ret = append(ret, tweetId)
		if index >= 0 {
			posts := this.posts[followeeId]
			pq.Enqueue([]int{posts[index].Time, followeeId, posts[index].TweetId, index - 1})
		}
	}

	return ret
}

func (this *Twitter2) Follow(followerId int, followeeId int) {
	if this.follows[followerId] == nil {
		this.follows[followerId] = make(map[int]bool)
	}
	this.follows[followerId][followeeId] = true
}

func (this *Twitter2) Unfollow(followerId int, followeeId int) {
	if this.follows[followerId] != nil {
		delete(this.follows[followerId], followeeId)
	}
}

func Test_TwitterConstructor() {
	twitter := TwitterConstructor()

	twitter.PostTweet(1, 5)
	log.Printf("getNewsFeed: %v", twitter.GetNewsFeed(1))
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	log.Printf("getNewsFeed: %v", twitter.GetNewsFeed(1))
	twitter.Unfollow(1, 2)
	log.Printf("getNewsFeed: %v", twitter.GetNewsFeed(1))

	log.Println()

	twitter2 := TwitterConstructor()

	twitter2.PostTweet(1, 5)
	twitter2.PostTweet(1, 7)
	log.Printf("getNewsFeed: %v", twitter2.GetNewsFeed(1))
	twitter2.PostTweet(2, 9)
	twitter2.PostTweet(1, 11)
	log.Printf("getNewsFeed: %v", twitter2.GetNewsFeed(1))
	twitter2.Follow(1, 2)
	log.Printf("getNewsFeed: %v", twitter2.GetNewsFeed(1))
	twitter2.Unfollow(1, 2)
	log.Printf("getNewsFeed: %v", twitter2.GetNewsFeed(1))
	log.Printf("getNewsFeed: %v", twitter2.GetNewsFeed(2))

	log.Println()

	twitter3 := TwitterConstructor()

	twitter3.PostTweet(1, 5)
	twitter3.PostTweet(1, 3)
	twitter3.PostTweet(1, 101)
	twitter3.PostTweet(1, 13)
	twitter3.PostTweet(1, 10)
	twitter3.PostTweet(1, 2)
	twitter3.PostTweet(1, 94)
	twitter3.PostTweet(1, 505)
	twitter3.PostTweet(1, 333)
	log.Printf("getNewsFeed: %v", twitter3.GetNewsFeed(1))
}

func Test_TwitterConstructor2() {
	twitter := TwitterConstructor()

	twitter.PostTweet(1, 5)
	log.Printf("getNewsFeed: %v", twitter.GetNewsFeed(1))
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	log.Printf("getNewsFeed: %v", twitter.GetNewsFeed(1))
	twitter.Unfollow(1, 2)
	log.Printf("getNewsFeed: %v", twitter.GetNewsFeed(1))

	log.Println()

	twitter2 := TwitterConstructor()

	twitter2.PostTweet(1, 5)
	twitter2.PostTweet(1, 7)
	log.Printf("getNewsFeed: %v", twitter2.GetNewsFeed(1))
	twitter2.PostTweet(2, 9)
	twitter2.PostTweet(1, 11)
	log.Printf("getNewsFeed: %v", twitter2.GetNewsFeed(1))
	twitter2.Follow(1, 2)
	log.Printf("getNewsFeed: %v", twitter2.GetNewsFeed(1))
	twitter2.Unfollow(1, 2)
	log.Printf("getNewsFeed: %v", twitter2.GetNewsFeed(1))
	log.Printf("getNewsFeed: %v", twitter2.GetNewsFeed(2))

	log.Println()

	twitter3 := TwitterConstructor()

	twitter3.PostTweet(1, 5)
	twitter3.PostTweet(1, 3)
	twitter3.PostTweet(1, 101)
	twitter3.PostTweet(1, 13)
	twitter3.PostTweet(1, 10)
	twitter3.PostTweet(1, 2)
	twitter3.PostTweet(1, 94)
	twitter3.PostTweet(1, 505)
	twitter3.PostTweet(1, 333)
	log.Printf("getNewsFeed: %v", twitter3.GetNewsFeed(1))
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
