### Working with Concurrency in Go

Go, often referred to as Golang, is well-known for making it remarkably easy to work with concurrency. In order to make a particular function run concurrently, all we have to do is prepend the word "go" to the function call, and it cheerfully runs in the background, as a GoRoutine. Go's built in scheduler takes are of making sure that a given GoRoutine runs when it should, and as efficiently as it can.

However, this does not mean that working with concurrency is simple in Go—thread safe programming takes careful planning, and most importantly it requires that developers have an absolutely solid understanding of how Go deals with concurrency.

In the standard library, Go offers us several ways of dealing with concurrently running parts of our program, right in the standard library: `sync.WaitGroup`, which lets us wait for tasks to finish; `sync.Mutex`, which allows us to lock and unlock resources, so that no two GoRoutines can access the same memory location at the same time; and finally, `Channels`, which allow GoRoutines to send and receive data to and from each other.

Go's approach to concurrency is fairly straightforward, and is more or less summed up this mantra: `Don't communicate by sharing memory; instead, share memory by communicating. Channels are the means by which we usually share memory by communicating`.

We'll cover the use of WaitGroups, Mutexes, and Channels, and we'll do so in detail. We'll also cover some of the problems inherent in concurrency, including early program termination and race conditions. Initially, we'll gain a good understanding of how these things work by solving some of the classic problems found in the field of computer science, including the `Dining Philosophers`, the `Producer/Consumer` problem, and the `Sleeping Barber`. These problems are classics for a reason: they force a developer to figure out the best approach to working with code that run concurrently, or in parallel. 

Finally, we'll finish with a more "real-world" problem, where we have to register a customer for some kind of subscription service, and take care of invoicing, registration, and all the things necessary to get a new customer up and running. We'll do so, naturally, as quickly as we can by dividing the necessary tasks up into smaller tasks, and having them run concurrently.

Overall:
  * Learn about the various ways Go makes working with concurrent programing simple.
  * Understand how concurrency works, and its advantages and pitfalls.
  * Learn how WaitGroups, Mutexes, and channels work.
  * Master concurrency by working with classic computer science problems, and by building a real-world example.
